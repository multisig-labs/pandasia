package pchain

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/blocks"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
	"github.com/ava-labs/coreth/plugin/evm"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"github.com/tidwall/gjson"
	"golang.org/x/exp/slog"
)

var mainnetCtx *snow.Context
var factory secp256k1.Factory

const RewardValidatorTxId = 20
const AddValidatorTxId = 12
const AddDelegatorTxId = 14

func init() {
	mainnetCtx = genMainnetCtx()
}

// The JSON for a block doesnt include tx type ids, which is annoying, so we fix it here
// Also make everything into basic types
type Block struct {
	Height uint64
	Id     string
	TypeId int64
	Txs    []*Tx
}

type Tx struct {
	Id             string
	TypeId         int64
	UnsignedTx     string
	RecoveredAddrP string
	RecoveredAddrC string
	// if TypeId=20, tx id of the validator or delegator tx signifying the rewards were earned
	EarnedRewardForTxId string
}

func NewBlock(blkBytes []byte) (*Block, error) {
	blkStruct, _, err := decodeBlock(mainnetCtx, blkBytes)
	if err != nil {
		return nil, err
	}

	b := &Block{
		Height: blkStruct.Height(),
		Id:     blkStruct.ID().String(),
		TypeId: parseTypeID(blkStruct.Bytes()),
		Txs:    []*Tx{},
	}

	for _, txStruct := range blkStruct.Txs() {
		j, err := json.Marshal(txStruct.Unsigned)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal error: %w", err)
		}

		t := &Tx{
			Id:         txStruct.TxID.String(),
			TypeId:     parseTypeID(txStruct.Bytes()),
			UnsignedTx: string(j),
		}

		if t.TypeId == RewardValidatorTxId {
			t.EarnedRewardForTxId = gjson.GetBytes(j, "txID").String()
		}

		if len(txStruct.Creds) > 0 {
			unsignedBytes := txStruct.Unsigned.Bytes()
			sigBytes := txStruct.Creds[0].(*secp256k1fx.Credential).Sigs[0]
			recoveredAddrP, recoveredAddrC, err := recoverAddrs(unsignedBytes, sigBytes)
			if err != nil {
				slog.Warn("unable to recoverAddrs txid: %s err: %v\n", txStruct.TxID, err)
			}
			t.RecoveredAddrP = recoveredAddrP
			t.RecoveredAddrC = recoveredAddrC
		}

		b.Txs = append(b.Txs, t)
	}
	return b, nil
}

func SuperFetchBlocks(uri string, startHeight int, num int) ([]*Block, error) {
	blks := make([]*Block, 0, num)
	ch := make(chan []byte, 100)
	end := make(chan struct{})
	go func() {
		for b := range ch {
			blk, err := NewBlock(b)
			if err != nil {
				slog.Warn("unable to create NewBlock", "bytes", b)
			} else {
				blks = append(blks, blk)
			}
		}
		end <- struct{}{}
	}()

	err := superFetch(uri, startHeight, num, ch)
	if err != nil {
		return nil, err
	}
	<-end
	return blks, nil
}

func FetchBlocks(uri string, startHeight int, num int) ([]*Block, error) {
	ctx := context.Background()
	c := platformvm.NewClient(uri)

	blks := make([]*Block, 0, num)
	for i := 0; i < startHeight+num; i++ {
		height := startHeight + i
		b, err := fetchBlock(ctx, c, height)
		if err != nil {
			slog.Warn("error fetching block", "height", height, "err", err)
		} else {
			blks = append(blks, b)
		}
	}
	return blks, nil
}

func MaxHeight(uri string) int64 {
	ctx := context.Background()
	c := platformvm.NewClient(uri)
	height, err := c.GetHeight(ctx)
	if err != nil {
		slog.Warn("error fetching height", "err", err, "ur")
		return 0
	}
	return int64(height)
}

func fetchBlock(ctx context.Context, c platformvm.Client, height int) (*Block, error) {
	mainnetCtx := genMainnetCtx()

	blkBytes, err := c.GetBlockByHeight(ctx, uint64(height))
	if err != nil {
		return nil, err
	}
	// Get both the block and the marshalled json
	blkStruct, _, err := decodeBlock(mainnetCtx, blkBytes)
	if err != nil {
		return nil, err
	}

	return NewBlock(blkStruct.Bytes())
}

// CodecID is first 4 bytes, typeID is next 2 bytes
func parseTypeID(b []byte) int64 {
	var result int64 = int64(b[4])<<8 | int64(b[5])
	return result
}

func decodeBlock(ctx *snow.Context, b []byte) (blocks.Block, string, error) {
	decoded := decodeProposerBlock(b)

	blk, js, err := decodeInnerBlock(ctx, decoded)
	if err != nil {
		return blk, "", err
	}
	return blk, string(js), nil
}

// Tries to decode as proposal block (post-Banff) if it fails just return the original bytes
func decodeProposerBlock(b []byte) []byte {
	innerBlk, err := block.Parse(b)
	if err != nil {
		return b
	}
	return innerBlk.Block()
}

func decodeInnerBlock(ctx *snow.Context, b []byte) (blocks.Block, string, error) {
	res, err := blocks.Parse(blocks.GenesisCodec, b)
	if err != nil {
		return res, "", fmt.Errorf("blocks.Parse error: %w %v", err, b)
	}

	res.InitCtx(ctx)
	j, err := json.Marshal(res)
	if err != nil {
		return res, "", fmt.Errorf("json.Marshal error: %w", err)
	}
	return res, string(j), nil
}

// Recover the P and C chain addrs from a tx sig
func recoverAddrs(unsignedBytes []byte, sigBytes [65]byte) (string, string, error) {
	txHash := hashing.ComputeHash256(unsignedBytes)
	pk, err := factory.RecoverHashPublicKey(txHash, sigBytes[:])
	if err != nil {
		return "", "", err
	}
	recoveredAddrP, err := address.FormatBech32("avax", pk.Address().Bytes())
	if err != nil {
		return "", "", err
	}
	recoveredAddrC := evm.PublicKeyToEthAddress(pk).String()
	return "P-" + recoveredAddrP, recoveredAddrC, nil
}

// Simple context so that Marshal works
func genMainnetCtx() *snow.Context {
	pChainID, _ := ids.FromString("11111111111111111111111111111111LpoYY")
	xChainID, _ := ids.FromString("2oYMBNV4eNHyqk2fjjV5nVQLDbtmNJzq5s3qs3Lo6ftnC6FByM")
	cChainID, _ := ids.FromString("2q9e4r6Mu3U68nU1fYjgbR6JvwrRx36CohpAX5UQxse55x1Q5")
	avaxAssetID, _ := ids.FromString("FvwEAhmxKfeiG8SnEvq42hc6whRyY3EFYAvebMqDNDGCgxN5Z")
	lookup := ids.NewAliaser()
	lookup.Alias(xChainID, "X")
	lookup.Alias(cChainID, "C")
	lookup.Alias(pChainID, "P")
	c := &snow.Context{
		NetworkID:   1,
		SubnetID:    [32]byte{},
		ChainID:     [32]byte{},
		NodeID:      [20]byte{},
		XChainID:    xChainID,
		CChainID:    cChainID,
		AVAXAssetID: avaxAssetID,
		BCLookup:    lookup,
	}
	return c
}

// Use Colly to fetch num blocks concurrently
func superFetch(uri string, startHeight int, num int, ch chan []byte) error {
	threads := runtime.NumCPU() * 3
	c := colly.NewCollector(colly.AllowURLRevisit())
	c.OnResponse(func(r *colly.Response) {
		blkHex := gjson.GetBytes(r.Body, "result.block").String()
		blkBytes, err := formatting.Decode(formatting.Hex, blkHex)
		if err != nil || len(blkBytes) == 0 {
			msg := gjson.GetBytes(r.Body, "error.message").String()
			// Ignore block not found at height msgs, they are pruned unnecessary blocks
			if !strings.HasPrefix(msg, "couldn't get block at height") {
				slog.Warn("unable to decode block", "msg", msg)
			}
		} else {
			ch <- blkBytes
		}
	})
	q, _ := queue.New(
		threads, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: num},
	)
	u, _ := url.Parse(uri + "/ext/bc/P")
	h := &http.Header{"Content-Type": {"application/json"}}

	for i := startHeight; i < startHeight+num; i++ {
		body := strings.NewReader(fmt.Sprintf(`{"id":0,"jsonrpc":"2.0","method":"platform.getBlockByHeight","params":{"height":"%d","encoding":"hex"}}`, i))
		r := &colly.Request{
			URL:     u,
			Headers: h,
			Method:  "POST",
			Body:    body,
		}
		q.AddRequest(r)
	}
	err := q.Run(c)
	close(ch)
	return err
}
