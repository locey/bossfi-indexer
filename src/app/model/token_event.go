package model

import (
	"bossfi-indexer/src/core/db"
	"time"
)

type TokenEvent struct {
	ID          int64     `json:"id" gorm:"column:id;primaryKey"`
	BlockNumber int64     `json:"block_number" gorm:"column:block_number"`
	BlockHash   string    `json:"block_hash" gorm:"column:block_hash"`
	TxHash      string    `json:"tx_hash" gorm:"column:tx_hash"`
	LogIndex    int       `json:"log_index" gorm:"column:log_index"`
	EventType   int       `json:"event_type" gorm:"column:event_type"`
	FromAddress string    `json:"from_address" gorm:"column:from_address"`
	ToAddress   string    `json:"to_address" gorm:"column:to_address"`
	Amount      string    `json:"amount" gorm:"column:amount;type:numeric(78,0)"`
	Confirmed   bool      `json:"confirmed" gorm:"column:confirmed;default:false"`
	Deleted     bool      `json:"deleted" gorm:"column:deleted;default:false"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	ModifyTime  time.Time `json:"modify_time" gorm:"column:modify_time;autoUpdateTime"`
}

const (
	EventTypeTransfer = 1
	EventTypeMint     = 2
)

func (TokenEvent) TableName() string {
	return "bii_token_event"
}

type TokenEventModel struct{}

// Create 创建记录
func (m *TokenEventModel) Create(event *TokenEvent) error {
	return db.DB.Create(event).Error
}

// GetByTxHashAndIndex 查询单条记录
func (m *TokenEventModel) GetByTxHashAndIndex(txHash string, index int) (*TokenEvent, error) {
	var event TokenEvent
	err := db.DB.Where("tx_hash = ? AND log_index = ?", txHash, index).First(&event).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// ListByBlockNumber 查询区块中的事件
func (m *TokenEventModel) ListByBlockNumber(blockNumber int64) ([]*TokenEvent, error) {
	var events []*TokenEvent
	err := db.DB.Where("block_number = ?", blockNumber).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

// ListAll 查询所有事件
func (m *TokenEventModel) ListAll() ([]*TokenEvent, error) {
	var events []*TokenEvent
	err := db.DB.Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

// Page 分页查询
func (m *TokenEventModel) Page(page, pageSize int) ([]*TokenEvent, int64, error) {
	var events []*TokenEvent
	var total int64

	res := db.DB.Model(&TokenEvent{})

	if err := res.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := res.Offset((page - 1) * pageSize).Limit(pageSize).Find(&events).Error; err != nil {
		return nil, 0, err
	}

	return events, total, nil
}

func (m *TokenEventModel) GetEarlyUnConfirmBlock(finalizedNumber uint64) ([]*TokenEvent, error) {
	var events []*TokenEvent
	err := db.DB.Model(&TokenEvent{}).Scopes(NotDeleted).Where("block_number < ?", finalizedNumber).Where("confirmed = ?", false).Order("block_number ASC").Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (m *TokenEventModel) GetLastBlockNumber() int64 {
	var blockNumber int64
	err := db.DB.Model(&TokenEvent{}).Scopes(NotDeleted).Select("block_number").Order("id DESC").Limit(1).Pluck("block_number", &blockNumber).Error
	if err != nil {
		return 0
	}
	return blockNumber
}

func (m *TokenEventModel) ConfirmedByIds(ids []int64) error {
	return db.DB.Model(&TokenEvent{}).Scopes(NotDeleted).Where("id in (?)", ids[:]).UpdateColumn("confirmed", true).UpdateColumn("modify_time", time.Now()).Error
}

func (m *TokenEventModel) DeleteByIds(ids []int64) error {
	return db.DB.Model(&TokenEvent{}).Where("id in (?)", ids[:]).UpdateColumn("deleted", true).UpdateColumn("modify_time", time.Now()).Error
}
