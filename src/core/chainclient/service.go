package chainclient

import (
	"bossfi-indexer/src/common/chain"
	"bossfi-indexer/src/core/chainclient/evm"
	"errors"
)

type ChainClient interface {
	Client() interface{}
}

func New(chainId int, endpoint string) (ChainClient, error) {
	switch chainId {
	case chain.EthChainID, chain.OptimismChainID, chain.SepoliaChainID:
		return evm.New(endpoint)
	default:
		return nil, errors.New("unsupported chain id")
	}
}
