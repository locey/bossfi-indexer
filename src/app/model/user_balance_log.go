package model

import (
	"bossfi-indexer/src/core/db"
	"time"
)

type UserBalanceLog struct {
	ID            int64     `json:"id" gorm:"column:id;primaryKey"`
	Address       string    `json:"address" gorm:"column:address"`
	LogType       int       `json:"log_type" gorm:"column:log_type"`
	BeforeBalance string    `json:"before_balance" gorm:"column:before_balance;type:numeric(78,0)"`
	ChangeBalance string    `json:"change_balance" gorm:"column:change_balance;type:numeric(78,0)"`
	AfterBalance  string    `json:"after_balance" gorm:"column:after_balance;type:numeric(78,0)"`
	TxHash        string    `json:"tx_hash" gorm:"column:tx_hash"`
	Deleted       bool      `json:"deleted" gorm:"column:deleted;default:false"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	ModifyTime    time.Time `json:"modify_time" gorm:"column:modify_time;autoUpdateTime"`
}

func (UserBalanceLog) TableName() string {
	return "bii_user_balance_log"
}

type UserBalanceLogModel struct{}

// Create 创建记录
func (m *UserBalanceLogModel) Create(log *UserBalanceLog) error {
	return db.DB.Create(log).Error
}

// GetByID 查询单条记录
func (m *UserBalanceLogModel) GetByID(id int64) (*UserBalanceLog, error) {
	var log UserBalanceLog
	err := db.DB.Scopes(NotDeleted).Where("id = ?", id).First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// ListByAddress 查询某个地址的所有日志
func (m *UserBalanceLogModel) ListByAddress(address string) ([]*UserBalanceLog, error) {
	var logs []*UserBalanceLog
	err := db.DB.Scopes(NotDeleted).Where("address = ?", address).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// Page 分页查询
func (m *UserBalanceLogModel) Page(page, pageSize int) ([]*UserBalanceLog, int64, error) {
	var logs []*UserBalanceLog
	var total int64

	res := db.DB.Scopes(NotDeleted).Model(&UserBalanceLog{})

	if err := res.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := res.Offset((page - 1) * pageSize).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}
