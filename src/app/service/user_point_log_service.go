package service

import (
	"bossfi-indexer/src/app/model"
)

type UserPointLogService struct {
	dao *model.UserPointsLogModel
}

func NewUserPointLogService() *UserPointLogService {
	return &UserPointLogService{
		dao: model.NewUserPointsLogModel(),
	}
}

// Create 创建记录
func (s *UserPointLogService) Create(log *model.UserPointsLog) error {
	return s.dao.Create(log)
}

// GetByID 查询单条记录
func (s *UserPointLogService) GetByID(id int64) (*model.UserPointsLog, error) {
	return s.dao.GetByID(id)
}

// ListByAddress 查询某个地址的所有日志
func (s *UserPointLogService) ListByAddress(address string) ([]*model.UserPointsLog, error) {
	return s.dao.ListByAddress(address)
}

// Page 分页查询
func (s *UserPointLogService) Page(page, pageSize int) ([]*model.UserPointsLog, int64, error) {
	return s.dao.Page(page, pageSize)
}
