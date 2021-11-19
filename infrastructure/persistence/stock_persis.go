package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/yejingxuan/accumulate/domain/entity"
)

type StockPersis struct {
	db *gorm.DB
}

func NewStockPersis(db *gorm.DB) *StockPersis {
	return &StockPersis{db}
}

func (s StockPersis) CreateStock(stock *entity.Stock) error {
	err := s.db.Table(entity.TableStock).Create(stock).Error
	return err
}

func (s StockPersis) GetStockInfoByCode(code string) (*entity.Stock, error) {
	res := entity.Stock{}
	/*err2 := s.db.Table(entity.TableStock).CreateTable(&res).Error
	logger.Error("err", zap.Any("err", err2))*/
	err := s.db.Table(entity.TableStock).Where("symbol = ?", code).Scan(&res).Error
	return &res, err
}
