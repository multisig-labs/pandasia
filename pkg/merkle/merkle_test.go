package merkle

import (
	"context"
	"os"
	"testing"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ethereum/go-ethereum/common"
	"github.com/multisig-labs/pandasia/pkg/db"
	"github.com/multisig-labs/pandasia/pkg/pchain"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/sha3"
)

func TestAddr(t *testing.T) {
	addr := "P-avax1u6xzj47hp5g0atfx4haadhzp6jah52sngfgdvg"
	_, _, addrBytes, err := address.Parse(addr)
	require.NoError(t, err)
	addrHex := common.BytesToAddress(addrBytes).Hex()
	t.Logf("Addr: %s", addrHex)
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
