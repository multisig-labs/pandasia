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

// type envConfig struct {
// 	RPCHostURL              string `env:"ETH_RPC_URL,default="`
// }

// var EnvConfig *envConfig

// func init() {
// 	EnvConfig = &envConfig{}
// 	if err := env.Bind(EnvConfig); err != nil {
// 		log.Fatalf("error binding config to env: %v", err)
// 	}
// 	if _, err := time.ParseDuration(EnvConfig.ValidationTxDelay); err != nil {
// 		log.Fatalf("invalid duration VALIDATION_TX_DELAY: %s %s", EnvConfig.ValidationTxDelay, err)
// 	}
// 	log.Infof("EnvConfig: %+v", EnvConfig)
// }

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
