package merkle

import (
	"context"
	"os"
	"testing"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ava-labs/avalanchego/utils/cb58"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/multisig-labs/pandasia/pkg/db"
	"github.com/multisig-labs/pandasia/pkg/pchain"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/sha3"
)

func TestAddr(t *testing.T) {
	addr := "P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww"
	_, _, addrBytes, err := address.Parse(addr)
	require.NoError(t, err)
	addrHex := common.BytesToAddress(addrBytes).Hex()
	t.Logf("Addr: %s", addrHex)
	t.Fatal()
}

func TestDecodeSig(t *testing.T) {
	sig := "24eWufzWvm38teEhNQmtE9N5BD12CWUawv1YtbYkuxeS5gGCN6CoZBgU4V4WDrLa5anYyTLGZT8nqiEsqX7hm1k3jofswfx"
	b, err := cb58.Decode(sig)
	require.NoError(t, err)
	// 6ac1cc3277dffe75d9cc8264acacc9f464762bab7ef73921a67dee1a398bd337 39cf19e2ff4c36ba64ed3684af9a72b59b7ccd16833666c81e84fb001bbb315a 00
	t.Logf("Addr: %x", b)
	t.Logf("R: %x  S: %x  V: %x", b[0:31], b[32:64], b[65:])
	t.Fatal()

}

func TestGenerate(t *testing.T) {
	ctx := context.Background()
	_, queries := db.OpenDB("../../data/pandasia.db")
	height, err := queries.MaxHeight(ctx)
	require.NoError(t, err)
	vaddrs, err := LoadAddrsFromDB(ctx, queries, pchain.AddDelegatorTxId, int(height))
	require.NoError(t, err)
	j, err := GenerateTree(vaddrs)
	require.NoError(t, err)
	err = os.WriteFile("tree.json", j, 0666)
	require.NoError(t, err)
	t.Fatal()
}

func TestProof(t *testing.T) {
	b, err := os.ReadFile("tree.json")
	require.NoError(t, err)
	p, err := GenerateProof(string(b), "0x7bDF8B86561d98d77e5BFc4B0eD20b7beB8fCdb6")
	require.NoError(t, err)
	t.Logf("Proof: %x", p)
	t.Fatal()
}

func Test0(t *testing.T) {
	leafEncodings := []string{smt.SOL_ADDRESS}
	values := [][]interface{}{
		{smt.SolAddress("0x1111111111111111111111111111111111111111")},
		{smt.SolAddress("0x424328BF10CDaEEDa6bb05A78cfF90a0BEA12c02")},
	}

	t1, err := smt.Of(values, leafEncodings)
	require.NoError(t, err)
	t.Logf("Merke Root: %x", t1.GetRoot())
	jsonValue, err := t1.TreeMarshal()
	require.NoError(t, err)
	t.Logf("%s", jsonValue)
	p, err := GenerateProof(string(jsonValue), "0x424328BF10CDaEEDa6bb05A78cfF90a0BEA12c02")
	require.NoError(t, err)
	t.Logf("Proof: %x", p)
	t.Fatal()
}

func TestLoadAddrFromTree(t *testing.T) {
	j := `
		{
			"format": "standard-v1",
			"tree": [
				"0x1261849b132545d29a7685fd7046d6200577cd912dda7f56b5a0d6dc16cb220d",
				"0x76b6437b39eb2e64dca0afc556654206c04fe45f2f779295a04526d4facaf34f",
				"0x3548906981a17dc8fedb3df19439a12bd993dbd555e0fb9e9a6eb967ee4401d0",
				"0xa52e4f15acf7ff84255187ffb785366ccc84195823bcbd8760bf1b857cfe0b28",
				"0x62ae3d5468d7a5dc6fcbc9597e8d38755e36ca08d189e51a4d9dfaab4c39746a",
				"0x4ebd6c0d9a8b9bb3495d47825b1171aeb966d8593a0a995994f2b4e0167bba8c",
				"0x20b2f891eaf390d96349554ee7297a8a8972c13215d1aa9dd752f9c6822c1888"
			],
			"values": [
				{
					"value": [
						"0x2892441ba9a160bcdc596dcd2cc3ad83c3493589"
					],
					"treeIndex": 3
				},
				{
					"value": [
						"0xeb520c1b17806339ca51ebe528860814351fb508"
					],
					"treeIndex": 4
				},
				{
					"value": [
						"0x424328bf10cdaeeda6bb05a78cff90a0bea12c02"
					],
					"treeIndex": 6
				},
				{
					"value": [
						"0x5cf998275803a7277926912defdf177b2e97b0b4"
					],
					"treeIndex": 5
				}
			],
			"leafEncoding": [
				"address"
			]
		}
	`
	addrs, err := LoadAddrsFromTree(j)
	require.NoError(t, err)
	spew.Dump(addrs)
	t.Fatal()
}

// func Test1(t *testing.T) {
// 	addrs := []string{
// 		"P-avax19zfygxaf59stehzedhxjesads0p5jdvfeedal0",
// 		"P-avax1adfqcxchsp3nnjj3a0jj3psgzs63ldggzew7c9",
// 		// "P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww",
// 		// "P-avax1tnuesf6cqwnjw7fxjyk7lhch0vhf0v95wj5jvy",
// 	}
// 	vaddrs := []mt.DataBlock{}
// 	for _, a := range addrs {
// 		// Parse P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww into just the address bytes
// 		_, _, addrBytes, err := address.Parse(a)
// 		require.NoError(t, err)
// 		t.Logf("Addr: %x", addrBytes)
// 		vaddr := ValidatorAddress{Addr: a, Data: addrBytes}
// 		vaddrs = append(vaddrs, &vaddr)
// 	}

// 	cfg := mt.Config{HashFunc: Hash, SortSiblingPairs: true, DisableLeafHashing: false, Mode: mt.ModeProofGenAndTreeBuild}
// 	tree, err := mt.New(&cfg, vaddrs)
// 	require.NoError(t, err)

// 	rootHash := tree.Root
// 	t.Logf("Merkle Root: %x\n", rootHash)
// 	proofs := tree.Proofs
// 	for i := 0; i < len(proofs); i++ {
// 		_, err := tree.Verify(vaddrs[i], proofs[i])
// 		require.NoError(t, err)
// 		t.Logf("Value: %v  Proof: %x", vaddrs[i], proofs[i])
// 	}

// 	t.Fatal()
// }

func Hash(data []byte) ([]byte, error) {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil), nil
}
