package evm

import (
	"bossfi-indexer/src/common/abi/tokenabi"
	"bossfi-indexer/src/core/chainclient/domain"
	"bossfi-indexer/src/core/config"
	"bossfi-indexer/src/core/log"
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	zazap "go.uber.org/zap"
	"math/big"
	"os"
)

type Evm struct {
	client *ethclient.Client
	chain  *config.ChainConfig
}

func New(chainConfig config.ChainConfig) (*Evm, error) {
	c, err := ethclient.Dial(chainConfig.Endpoint)
	if err != nil {
		log.Logger.Error("")
	}

	return &Evm{
		client: c,
		chain:  &chainConfig,
	}, err
}

func (c *Evm) Client() interface{} {
	return c.client
}

func (c *Evm) ChainConfig() *config.ChainConfig {
	return c.chain
}

func (c *Evm) GetBossToken() (*tokenabi.AbiFilterer, error) {
	contractAddress := common.HexToAddress(c.chain.BossTokenAddress) // BossToken 实际地址
	bossToken, err := tokenabi.NewAbiFilterer(contractAddress, c.client)
	if err != nil {
		panic(err)
	}

	return bossToken, err
}

func (c *Evm) GetSafeBlock() (*domain.Block, error) {
	block, err := c.client.BlockByNumber(context.Background(), big.NewInt(rpc.SafeBlockNumber.Int64()))
	if err != nil {
		log.Logger.Error("GetSafeBlock failed!", zazap.Error(err))
		return nil, err
	}

	return domain.ToBlock(block), nil
}

func (c *Evm) GetFinalizedBlock() (*domain.Block, error) {
	block, err := c.client.BlockByNumber(context.Background(), big.NewInt(rpc.FinalizedBlockNumber.Int64()))
	if err != nil {
		log.Logger.Error("GetFinalizedBlock failed!", zazap.Error(err))
		return nil, err
	}
	return domain.ToBlock(block), nil
}

func (c *Evm) GetLogs(fromBlock *big.Int, toBlock *big.Int, contractAddress string) ([]types.Log, error) {
	q := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics:    nil,
	}
	logs, err := c.client.FilterLogs(context.Background(), q)
	if err != nil {
		log.Logger.Error("GetLogs failed!", zazap.Error(err))
		return nil, err
	}

	return logs, nil
}

// LoadABI 从文件加载 ABI
func LoadABI(path string) (abi.ABI, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return abi.ABI{}, err
	}

	var contractABI abi.ABI
	if err := json.Unmarshal(raw, &contractABI); err != nil {
		return abi.ABI{}, err
	}

	return contractABI, nil
}
