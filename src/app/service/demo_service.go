package service

import (
	"bossfi-indexer/src/app/model"
)

type DemoService struct {
	dao *model.DemoModel
}

func NewDemoService() *DemoService {
	return &DemoService{dao: &model.DemoModel{}}
}

// Create 创建记录
func (s *DemoService) Create(demo *model.Demo) error {
	return s.dao.Create(demo)
}

// GetById 查询单条记录
func (s *DemoService) GetById(id int64) (*model.Demo, error) {
	return s.dao.GetById(id)
}

// Update 更新记录
func (s *DemoService) Update(demo *model.Demo) error {
	return s.dao.UpdateById(demo)
}

// Delete 软删除记录
func (s *DemoService) Delete(id int64) error {
	return s.dao.DeleteById(id)
}

// List 查询所有未删除记录
func (s *DemoService) List() ([]*model.Demo, error) {
	return s.dao.List()
}

// Page 查询分页数据
func (s *DemoService) Page(page, pageSize int) ([]*model.Demo, int64, error) {
	return s.dao.Page(page, pageSize)
}
