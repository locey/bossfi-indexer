package model

import (
	"gorm.io/gorm"
	"math/big"
	"time"
)

type UserBalance struct {
	ID         int64     `json:"id" gorm:"column:id;primaryKey"`
	Address    string    `json:"address" gorm:"column:address"`
	Balance    string    `json:"balance" gorm:"column:balance;type:numeric(78,0)"`
	Points     string    `json:"points" gorm:"column:points;type:numeric(78,0);default:0"`
	Deleted    bool      `json:"deleted" gorm:"column:deleted;default:false"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	ModifyTime time.Time `json:"modify_time" gorm:"column:modify_time;autoUpdateTime"`
}

func (UserBalance) TableName() string {
	return "bii_user_balance"
}

type UserBalanceModel struct {
	DB *gorm.DB
}

// Create 创建记录
func (m *UserBalanceModel) Create(balance *UserBalance) error {
	return m.DB.Create(balance).Error
}

// GetByAddress 查询单条记录
func (m *UserBalanceModel) GetByAddress(address string) (*UserBalance, error) {
	var balance UserBalance
	err := m.DB.Scopes(NotDeleted).Model(&UserBalance{}).Where("address = ?", address).First(&balance).Error
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

// Update 更新记录
func (m *UserBalanceModel) Update(balance *UserBalance) error {
	return m.DB.Save(balance).Error
}

// DeleteByID 逻辑删除记录
func (m *UserBalanceModel) DeleteByID(id int64) error {
	return m.DB.Scopes(NotDeleted).Model(&UserBalance{}).
		Where("id = ?", id).
		Update("deleted", true).Error
}

// List 查询所有未删除记录
func (m *UserBalanceModel) List() ([]*UserBalance, error) {
	var list []*UserBalance
	err := m.DB.Scopes(NotDeleted).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// Page 分页查询
func (m *UserBalanceModel) Page(page, pageSize int) ([]*UserBalance, int64, error) {
	var list []*UserBalance
	var total int64

	res := m.DB.Scopes(NotDeleted).Model(&UserBalance{})

	if err := res.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := res.Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *UserBalanceModel) AddPoints(address string, points *big.Int) error {
	// 原子性
	if err := m.DB.Model(&UserBalance{}).Where("address = ?", address).UpdateColumn("points", gorm.Expr("points + ?", points)).Error; err != nil {
		return err
	}

	return nil
}
