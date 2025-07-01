package service

import (
	"bossfi-indexer/src/app/model"
)

type UserBalanceService struct {
	dao *model.UserBalanceModel
}

func NewUserBalanceService() *UserBalanceService {
	return &UserBalanceService{
		dao: &model.UserBalanceModel{},
	}
}

// Create 创建记录
func (s *UserBalanceService) Create(balance *model.UserBalance) error {
	return s.dao.Create(balance)
}

// GetByAddress 查询单条记录
func (s *UserBalanceService) GetByAddress(address string) (*model.UserBalance, error) {
	return s.dao.GetByAddress(address)
}

// Update 更新记录
func (s *UserBalanceService) Update(balance *model.UserBalance) error {
	return s.dao.Update(balance)
}

// Delete 软删除记录
func (s *UserBalanceService) Delete(id int64) error {
	return s.dao.DeleteByID(id)
}

// List 查询所有未删除记录
func (s *UserBalanceService) List() ([]*model.UserBalance, error) {
	return s.dao.List()
}

// Page 分页查询
func (s *UserBalanceService) Page(page, pageSize int) ([]*model.UserBalance, int64, error) {
	return s.dao.Page(page, pageSize)
}
