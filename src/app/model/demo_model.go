package model

import (
	"bossfi-indexer/src/core/db"
	"time"
)

type DemoModel struct {
}

type Demo struct {
	ID         int64                  `json:"id" gorm:"column:id;primaryKey"`
	Address    string                 `json:"address" gorm:"column:address"`
	Logs       map[string]interface{} `json:"logs" gorm:"column:logs;type:jsonb;serializer:json"`
	Deleted    bool                   `json:"deleted" gorm:"column:deleted;default:false"`
	CreateTime time.Time              `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	ModifyTime time.Time              `json:"modify_time" gorm:"column:modify_time;autoUpdateTime"`
}

func (Demo) TableName() string {
	return "bossfi_demo"
}

// Create 创建记录
func (m *DemoModel) Create(demo *Demo) error {
	return db.DB.Create(demo).Error
}

// GetById 查询单条记录
func (m *DemoModel) GetById(id int64) (*Demo, error) {
	var demo Demo
	err := db.DB.Scopes(NotDeleted).Where("id = ?", id).First(&demo).Error
	if err != nil {
		return nil, err
	}
	return &demo, nil
}

// UpdateById 更新记录
func (m *DemoModel) UpdateById(demo *Demo) error {
	return db.DB.Save(demo).Error
}

// DeleteById 逻辑删除记录
func (m *DemoModel) DeleteById(id int64) error {
	return db.DB.Scopes(NotDeleted).Model(&Demo{}).
		Where("id = ?", id).
		Update("deleted", true).Error
}

// List 查询所有未删除记录
func (m *DemoModel) List() ([]*Demo, error) {
	var list []*Demo
	err := db.DB.Scopes(NotDeleted).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// Page 查询分页数据
func (m *DemoModel) Page(page, size int) ([]*Demo, int64, error) {
	var list []*Demo
	var total int64

	res := db.DB.Scopes(NotDeleted).Model(&Demo{})

	// 获取总数
	if err := res.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := res.Offset((page - 1) * size).Limit(size).Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
