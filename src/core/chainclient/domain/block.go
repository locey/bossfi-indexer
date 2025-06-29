package domain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"time"
)

// Block 一个区块所需要的属性
type Block struct {

	// 区块高度
	Number uint64 `json:"Number"`

	// 区块时间
	Time uint64 `json:"Time"`

	// 区块nonce
	Nonce uint64 `json:"Nonce"`

	// 区块哈希
	Hash string `json:"Hash"`

	// 父区块哈希
	ParentHash string `json:"ParentHash"`

	// 交易树根节点的哈希值，用于验证区块中的交易数据
	TxHash string `json:"TxHash"`

	// 状态树根节点的哈希值，用于验证区块中的状态数据
	StateRoot string `json:"StateRoot"`

	// 收据树根节点的哈希值，用于验证区块中的收据数据
	ReceiptHash string `json:"ReceiptHash"`

	// 区块大小
	Size uint64 `json:"Size"`

	// 区块gas使用量
	GasUsed string `json:"GasUsed"`

	// 区块gasLimit
	GasLimit string `json:"GasLimit"`

	// 区块交易数
	TransactionCount int `json:"TransactionCount"`

	// 区块交易列表
	Transactions []*Transaction `json:"Transactions"`
}

type Transaction struct {
	// 交易哈希值
	Hash string `json:"hash"`

	// 接收方地址（可能是合约创建时为nil）
	To string `json:"to,omitempty"`

	// 交易金额（单位：wei）
	Value string `json:"value"`

	// Gas价格
	GasPrice string `json:"gasPrice"`

	// Gas上限
	GasLimit string `json:"gas"`

	// 随机数，用于防止重放攻击
	Nonce string `json:"nonce"`

	// 输入数据（如调用智能合约的方法和参数）
	Data []byte `json:"input"`

	// 时间戳（区块中交易的时间）
	Timestamp time.Time `json:"timestamp"`
}

func ToBlock(block *types.Block) *Block {
	var transactions []*Transaction

	for _, tx := range block.Transactions() {
		transactions = append(transactions, ToTransaction(tx))
	}

	blockDomain := &Block{
		Number:           block.Number().Uint64(),
		Time:             block.Time(),
		Nonce:            block.Nonce(),
		Hash:             block.Hash().Hex(),
		ParentHash:       block.ParentHash().Hex(),
		TxHash:           block.TxHash().Hex(),
		StateRoot:        block.Root().Hex(),
		ReceiptHash:      block.ReceiptHash().Hex(),
		Size:             block.Size(),
		GasUsed:          fmt.Sprintf("0x%x", block.GasUsed()),
		GasLimit:         fmt.Sprintf("0x%x", block.GasLimit()),
		TransactionCount: len(block.Transactions()),
		Transactions:     transactions,
	}

	return blockDomain
}

func ToTransaction(tx *types.Transaction) *Transaction {
	var toHex string
	if tx.To() != nil {
		toHex = tx.To().Hex()
	}

	return &Transaction{
		Hash:      tx.Hash().Hex(),
		To:        toHex,
		Value:     tx.Value().Text(16),
		GasPrice:  "0x" + tx.GasPrice().Text(16),
		GasLimit:  fmt.Sprintf("0x%x", tx.Gas()),
		Nonce:     fmt.Sprintf("0x%x", tx.Nonce()),
		Data:      tx.Data(),
		Timestamp: tx.Time(),
	}
}
