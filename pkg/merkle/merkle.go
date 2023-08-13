package merkle

import (
	"context"

	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ethereum/go-ethereum/common"
	"github.com/multisig-labs/pandasia/pkg/db"
)

type ValidatorAddress struct {
	// P-avax1blah
	Addr string
	// 0x1234...
	AddrHex string
}

func GenerateProof(jsonTree string, addr string) (any, error) {
	tree, err := smt.Load([]byte(jsonTree))
	if err != nil {
		return "", err
	}
	value := []interface{}{
		smt.SolAddress(addr),
	}
	proof, err := tree.GetProof(value)
	if err != nil {
		return "", err
	}
	return proof, nil
}

func GenerateTree(vaddrs []ValidatorAddress) (string, error) {
	leafEncodings := []string{smt.SOL_ADDRESS}
	values := [][]interface{}{}
	for _, vaddr := range vaddrs {
		values = append(values, []interface{}{smt.SolAddress(vaddr.AddrHex)})
	}
	tree, err := smt.Of(values, leafEncodings)
	if err != nil {
		return "", err
	}
	jsonValue, err := tree.TreeMarshal()
	if err != nil {
		return "", err
	}
	return string(jsonValue), nil
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
