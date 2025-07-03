package api

import (
	"bossfi-indexer/src/app/service"
	"bossfi-indexer/src/core/result"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserBalanceApi struct {
	svc *service.UserBalanceService
}

func NewUserBalanceApi() *UserBalanceApi {
	return &UserBalanceApi{
		svc: service.NewUserBalanceService(),
	}
}

// GetBalanceByAddress godoc
// @Summary      获取用户余额接口
// @Description  根据地址获取用户的Token余额信息
// @Tags         用户余额
// @Accept       json
// @Produce      json
//
// @Param        address  path  string  true  "用户地址 (例如: 0x69b821F23bc4E537d82a65593b032B8ad13B6c0c)"
//
// @Success      200  {object}  model.UserBalance  "成功返回用户余额数据"
// @Failure      400  {object}  result.Response  "参数错误"
// @Failure      404  {object}  result.Response  "数据库中未找到该地址的余额记录"
// @Failure      500  {object}  result.Response  "数据库查询失败或其他内部错误"
//
// @Router       /api/v1/user_balance/{address} [GET]
func (s *UserBalanceApi) GetBalanceByAddress(c *gin.Context) {
	address := c.Params.ByName("address")
	if address == "" || address == "0x0000000000000000000000000000000000000000" {
		result.Error(c, result.InvalidParameter)
		return
	}

	balance, err := s.svc.GetByAddress(address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result.Error(c, result.DBNotExist)
			return
		}
		result.Error(c, result.DBQueryFailed)
		return
	}

	result.OK(c, balance)
}
