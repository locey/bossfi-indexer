package model

import (
	"bossfi-indexer/src/core/db"
	"gorm.io/gorm"
	"time"
)

type UserPointsLog struct {
	ID           int64     `json:"id" gorm:"column:id;primaryKey"`
	Address      string    `json:"address" gorm:"column:address;type:varchar(42)"`
	LogType      int       `json:"log_type" gorm:"column:log_type"`
	BeforePoints string    `json:"before_points" gorm:"column:before_points;type:numeric(78,0)"`
	ChangePoints string    `json:"change_points" gorm:"column:change_points;type:numeric(78,0)"`
	AfterPoints  string    `json:"after_points" gorm:"column:after_points;type:numeric(78,0)"`
	TxHash       string    `json:"tx_hash" gorm:"column:tx_hash;type:varchar(66)"`
	Deleted      bool      `json:"deleted" gorm:"column:deleted;default:false"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	ModifyTime   time.Time `json:"modify_time" gorm:"column:modify_time;autoUpdateTime"`
}

const (
	PointsTypeSystem  = 0 // 系统
	PointsTypeIncome  = 1 // 收入
	PointsTypeExpense = 2 // 支出
)

func (UserPointsLog) TableName() string {
	return "bii_user_points_log"
}

type UserPointsLogModel struct {
	DB *gorm.DB
}

func NewUserPointsLogModel() *UserPointsLogModel {
	return &UserPointsLogModel{DB: db.DB}
}

// Create 创建记录
func (m *UserPointsLogModel) Create(log *UserPointsLog) error {
	return m.DB.Create(log).Error
}

// CreateBatch 批量创建记录
func (m *UserPointsLogModel) CreateBatch(logs []*UserPointsLog) error {
	return m.DB.Create(logs).Error
}

// GetByID 查询单条记录
func (m *UserPointsLogModel) GetByID(id int64) (*UserPointsLog, error) {
	var log UserPointsLog
	err := m.DB.Model(&UserPointsLog{}).Where("id = ?", id).First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// ListByAddress 查询某个地址的所有日志
func (m *UserPointsLogModel) ListByAddress(address string) ([]*UserPointsLog, error) {
	var logs []*UserPointsLog
	err := m.DB.Model(&UserPointsLog{}).Where("address = ?", address).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// Page 分页查询
func (m *UserPointsLogModel) Page(page, pageSize int) ([]*UserPointsLog, int64, error) {
	var logs []*UserPointsLog
	var total int64

	res := m.DB.Model(&UserPointsLog{})

	if err := res.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := res.Offset((page - 1) * pageSize).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}
