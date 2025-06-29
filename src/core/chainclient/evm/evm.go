package evm

import (
	"bossfi-indexer/src/core/log"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Evm struct {
	client *ethclient.Client
}

func New(nodeUrl string) (*Evm, error) {
	c, err := ethclient.Dial(nodeUrl)
	if err != nil {
		log.Logger.Error("")
	}

	return &Evm{
		client: c,
	}, err
}

func (c *Evm) Client() interface{} {
	return c.client
}
