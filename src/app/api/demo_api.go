package api

import (
	"bossfi-indexer/src/app/model"
	"bossfi-indexer/src/app/service"
	"bossfi-indexer/src/core/result"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DemoApi struct {
	svc *service.DemoService
}

func NewDemoApi() *DemoApi {
	return &DemoApi{
		svc: service.NewDemoService(),
	}
}

// Create godoc
// @Summary      创建数据
// @Description  用于测试服务器连通性
// @Tags         示例接口
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string]string
// @Router       /demo/create [POST]
func (s *DemoApi) Create(c *gin.Context) {
	var req model.Demo
	if err := c.ShouldBindJSON(&req); err != nil {
		result.Error(c, result.InvalidParameter)
		return
	}
	if err := s.svc.Create(&req); err != nil {
		result.Error(c, result.DBCreateFailed)
		return
	}
	result.OK(c, req)
}

// GetById 查询数据
func (s *DemoApi) GetById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	demo, err := s.svc.GetById(id)
	if err != nil {
		result.Error(c, result.DBNotExist)
		return
	}
	result.OK(c, demo)
}

// Update 更新数据
func (s *DemoApi) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req model.Demo
	if err := c.ShouldBindJSON(&req); err != nil {
		result.Error(c, result.InvalidParameter)
		return
	}
	req.ID = id
	if err := s.svc.Update(&req); err != nil {
		result.Error(c, result.DBUpdateFailed)
		return
	}
	result.OK(c, req)
}

// Delete 删除数据
func (s *DemoApi) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := s.svc.Delete(id); err != nil {
		result.Error(c, result.DBDeleteFailed)
		return
	}
	result.OK(c, nil)
}

// List 查询列表
func (s *DemoApi) List(c *gin.Context) {
	list, err := s.svc.List()
	if err != nil {
		result.Error(c, result.DBQueryFailed)
		return
	}
	result.OK(c, list)
}

// Page 分页查询数据
func (s *DemoApi) Page(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	list, total, err := s.svc.Page(page, pageSize)
	if err != nil {
		result.Error(c, result.DBQueryFailed)
		return
	}

	// 返回分页结果
	result.OK(c, gin.H{
		"list":  list,
		"total": total,
	})
}
