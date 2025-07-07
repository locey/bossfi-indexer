package api

import (
	"bossfi-indexer/src/app/service"
	"bossfi-indexer/src/core/result"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserPointLogApi struct {
	svc *service.UserPointLogService
}

func NewUserPointLogApi() *UserPointLogApi {
	return &UserPointLogApi{
		svc: service.NewUserPointLogService(),
	}
}

// ListByAddress godoc
// @Summary      获取用户积分日志接口
// @Description  根据地址获取用户的积分变更记录
// @Tags         用户积分
// @Accept       json
// @Produce      json
// @Param        address  path  string  true  "用户地址"
// @Success      200  {object}  []model.UserPointLog  "成功返回用户积分日志数据"
// @Failure      400  {object}  result.Response  "参数错误"
// @Failure      404  {object}  result.Response  "数据库中未找到该地址的日志记录"
// @Failure      500  {object}  result.Response  "数据库查询失败或其他内部错误"
// @Router       /api/v1/user_point_log/{address} [GET]
func (s *UserPointLogApi) ListByAddress(c *gin.Context) {
	address := c.Params.ByName("address")
	if address == "" {
		result.Error(c, result.InvalidParameter)
		return
	}

	logs, err := s.svc.ListByAddress(address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result.Error(c, result.DBNotExist)
			return
		}
		result.Error(c, result.DBQueryFailed)
		return
	}

	result.OK(c, logs)
}
