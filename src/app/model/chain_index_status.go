package model

import (
	"time"
)

func init() {
	RegisterModel(&ChainIndexStatus{})
}

type ChainIndexStatus struct {
	ID               int64     `json:"id" gorm:"column:id;primaryKey"`
	ChainID          int       `json:"chain_id" gorm:"column:chain_id"`
	LastSyncBlockNum int64     `json:"last_block_num" gorm:"column:last_block_num"`
	LastSyncTime     time.Time `json:"last_sync_time" gorm:"column:last_sync_time"`
	CreateTime       time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	ModifyTime       time.Time `json:"modify_time" gorm:"column:modify_time;autoUpdateTime"`
}

func (ChainIndexStatus) TableName() string {
	return "bii_chain_index_status"
}
