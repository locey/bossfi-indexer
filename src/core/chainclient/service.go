package chainclient

import (
	"bossfi-indexer/src/common/chain"
	"bossfi-indexer/src/core/chainclient/evm"
	"bossfi-indexer/src/core/config"
	"errors"
)

type ChainClient interface {
	Client() interface{}
}

func New(chainConfig config.ChainConfig) (ChainClient, error) {
	switch chainConfig.ChainId {
	case chain.EthChainID, chain.OptimismChainID, chain.SepoliaChainID:
		return evm.New(chainConfig)
	default:
		return nil, errors.New("unsupported chain id")
	}
}
