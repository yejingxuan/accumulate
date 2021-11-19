package application

import (
	"github.com/yejingxuan/accumulate/domain/entity"
	"github.com/yejingxuan/accumulate/domain/repository"
	"github.com/yejingxuan/accumulate/infrastructure/crawler"
)

type StockAppInterface interface {
	GetStockInfoByCode(code string) (*entity.Stock, error) //获取stock详细信息

	UpdateAll() error //更新全部数据
}

//StockAppInterface实现
type stockApp struct {
	stockRepo repository.StockRepo
}

//构造函数
func NewStockApp(repo repository.StockRepo) *stockApp {
	return &stockApp{stockRepo: repo}
}

//获取stock详细信息
func (s stockApp) GetStockInfoByCode(code string) (*entity.Stock, error) {
	info, err := s.stockRepo.GetStockInfoByCode(code)
	return info, err
}

//更新全部数据
func (s stockApp) UpdateAll() error {
	crawler.ExecXueQiuJob(s.stockRepo)
	return nil
}
