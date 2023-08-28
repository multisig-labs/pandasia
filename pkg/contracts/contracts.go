package contracts

import (
	"fmt"

	"github.com/ava-labs/coreth/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/multisig-labs/pandasia/pkg/contracts/pandasia"
)

type Contracts struct {
	Url      string
	Pandasia *pandasia.PandasiaCaller
}

func NewContracts(url string, address string) (*Contracts, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("NewContracts ethclient.Dial error: %w", err)
	}

	p, err := pandasia.NewPandasiaCaller(common.HexToAddress(address), client)
	if err != nil {
		return nil, fmt.Errorf("NewContracts storage.NewStorage error: %w", err)
	}
	return &Contracts{
		Url:      url,
		Pandasia: p,
	}, nil
}
