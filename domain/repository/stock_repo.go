package repository

import "github.com/yejingxuan/accumulate/domain/entity"

type StockRepo interface {
	CreateStock(stock *entity.Stock) error //添加信息
	GetStockInfoByCode(code string) (*entity.Stock, error) //获取详情
}
