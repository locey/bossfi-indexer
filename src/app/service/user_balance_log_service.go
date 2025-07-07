package service

import (
	"bossfi-indexer/src/app/model"
	"bossfi-indexer/src/core/db"
)

type UserBalanceLogService struct {
	dao *model.UserBalanceLogModel
}

func NewUserBalanceLogService() *UserBalanceLogService {
	return &UserBalanceLogService{
		dao: &model.UserBalanceLogModel{DB: db.DB},
	}
}

// Create 创建记录
func (s *UserBalanceLogService) Create(log *model.UserBalanceLog) error {
	return s.dao.Create(log)
}

// GetByID 查询单条记录
func (s *UserBalanceLogService) GetByID(id int64) (*model.UserBalanceLog, error) {
	return s.dao.GetByID(id)
}

// ListByAddress 查询某个地址的所有日志
func (s *UserBalanceLogService) ListByAddress(address string) ([]*model.UserBalanceLog, error) {
	return s.dao.ListByAddress(address)
}

// Page 分页查询
func (s *UserBalanceLogService) Page(page, pageSize int) ([]*model.UserBalanceLog, int64, error) {
	return s.dao.Page(page, pageSize)
}
