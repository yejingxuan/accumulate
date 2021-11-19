package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yejingxuan/accumulate/application"
	"github.com/yejingxuan/accumulate/infrastructure/logger"
	"github.com/yejingxuan/accumulate/interface/dto"
	"go.uber.org/zap"
)

type stockHandler struct {
	stockApp application.StockAppInterface
}

//stockHandler 构造函数
func NewStockHandler(stockApp application.StockAppInterface) *stockHandler {
	return &stockHandler{stockApp}
}

// @Summary 获取stock详细信息
// @Tags 信息获取
// @Router /accumulate/v1/info/stock/{code} [get]
// @Param code path string true "stock-code"
// @Success 200 {string} json "{"code":0,"data":{},"msg":"ok"}"
func (h stockHandler) GetStockInfo(c *gin.Context) {
	code := c.Param("code")
	res, err := h.stockApp.GetStockInfoByCode(code)
	if err != nil {
		logger.Error("GetStockInfoById err", zap.Any("err", err))
		dto.JSON(c, dto.ErrBadRequest, err)
		return
	}
	dto.JSON(c, dto.Success, res)
}

// @Summary 更新全部数据
// @Tags 任务执行
// @Router /accumulate/v1/exec/update/all [post]
// @Success 200 {string} json "{"code":0,"data":{},"msg":"ok"}"
func (h stockHandler) UpdateAll(c *gin.Context) {
	err := h.stockApp.UpdateAll()
	if err != nil {
		logger.Error("UpdateAll err", zap.Any("err", err))
		dto.JSON(c, dto.ErrBadRequest, err)
		return
	}
	dto.JSON(c, dto.Success, nil)
}
