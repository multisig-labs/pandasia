package merkle

import (
	"context"
	"strings"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ethereum/go-ethereum/common"
	"github.com/multisig-labs/pandasia/pkg/db"
	"golang.org/x/exp/slog"
)

type ValidatorAddress struct {
	// P-avax1blah
	Addr string
	// 0x1234...
	AddrHex string
}

// const TREE_TYPE_DELEGATOR = "delegator"
const TREE_TYPE_VALIDATOR = "validator"

// addr is P-avax1blah
func GenerateProof(jsonTree string, addr string) ([][]byte, error) {
	var addrHex string

	if strings.HasPrefix(addr, "P-") {
		// Parse P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww into just the address bytes
		_, _, addrBytes, err := address.Parse(addr)
		if err != nil {
			return nil, err
		}
		addrHex = common.BytesToAddress(addrBytes).Hex()
	} else {
		addrHex = addr
	}

	smtTree, err := smt.Load([]byte(jsonTree))
	if err != nil {
		return nil, err
	}

	value := []interface{}{
		smt.SolAddress(addrHex),
	}

	// slog.Info("GenerateProof", "root", hexutil.Encode(smtTree.GetRoot()), "addr", addr, "addrHex", addrHex)
	proof, err := smtTree.GetProof(value)
	if err != nil {
		return nil, err
	}

	return proof, nil
}

func GenerateTree(vaddrs []ValidatorAddress) ([]byte, error) {
	leafEncodings := []string{smt.SOL_ADDRESS}
	values := [][]interface{}{}
	for _, vaddr := range vaddrs {
		values = append(values, []interface{}{smt.SolAddress(vaddr.AddrHex)})
	}

	tree, err := smt.Of(values, leafEncodings)
	if err != nil {
		return nil, err
	}

	jsonValue, err := tree.TreeMarshal()
	if err != nil {
		return nil, err
	}

	slog.Info("TreeMarshal", "bytes", len(jsonValue))
	return jsonValue, nil
}

func SaveTreeToDB(ctx context.Context, queries *db.Queries, treeType string, height int, tree []byte) (int64, error) {
	args := db.CreateMerkleTreeAndReturnIdParams{
		Height:   int64(height),
		TreeType: treeType,
		Tree:     string(tree),
	}
	return queries.CreateMerkleTreeAndReturnId(ctx, args)
}

func FindTreeByType(ctx context.Context, queries *db.Queries, treeType string) ([]byte, error) {
	t, err := queries.FindMerkleTreeByType(ctx, treeType)
	return []byte(t.Tree), err
}

func LoadAddrsFromDB(ctx context.Context, queries *db.Queries, txType int, height int) ([]ValidatorAddress, error) {
	vaddrs := []ValidatorAddress{}
	args := db.FindAddrsForMerkleTreeParams{
		TypeID: int64(txType),
		Height: int64(height),
	}
	addrs, err := queries.FindAddrsForMerkleTree(ctx, args)
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		// Parse P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww into just the address bytes
		_, _, addrBytes, err := address.Parse(addr.String)
		if err != nil {
			return nil, err
		}
		vaddr := ValidatorAddress{Addr: addr.String, AddrHex: common.BytesToAddress(addrBytes).Hex()}
		vaddrs = append(vaddrs, vaddr)
	}
	return vaddrs, nil
}

func VerifyTree(ctx context.Context, queries *db.Queries, root string) (bool, error) {
	dbTree, err := queries.FindMerkleTreeByRoot(ctx, root)
	if err != nil {
		return false, err
	}
	slog.Info("verifying tree", "id", dbTree.ID, "height", dbTree.Height)

	smtTree, err := smt.Load([]byte(dbTree.Tree))
	if err != nil {
		return false, err
	}

	// TODO map string tree type to tx type id
	txType := 12
	vaddrs, err := LoadAddrsFromDB(ctx, queries, txType, int(dbTree.Height))
	if err != nil {
		return false, err
	}

	slog.Info("loading addrs from db", "len(vaddrs)", len(vaddrs))
	okCnt := 0

	for _, a := range vaddrs {
		value := []interface{}{
			smt.SolAddress(a.AddrHex),
		}

		proof, err := smtTree.GetProof(value)
		if err == nil {
			ok, err := smtTree.Verify(proof, value)
			if err == nil && ok {
				okCnt++
			}
		}
	}

	slog.Info("Verification", "ok", okCnt, "err", len(vaddrs)-okCnt)
	return true, nil
}
