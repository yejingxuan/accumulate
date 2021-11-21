package persistence

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"github.com/yejingxuan/accumulate/domain/entity"
	"github.com/yejingxuan/accumulate/domain/repository"
	"github.com/yejingxuan/accumulate/infrastructure/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

// Repositories 总仓储机构提，包含多个领域仓储接口，以及一个DB实例
type Repositories struct {
	db        *gorm.DB
	StockRepo repository.StockRepo
}

// NewRepositories 初始化所有域的总仓储实例，将实例通过依赖注入方式，将DB实例注入到领域层
func NewRepositories() (*Repositories, error) {
	dbCfg := config.CoreConf.Server.DB
	//d, err := gorm.Open("postgres", dbCfg.Dsn)
	d, err := gorm.Open(sqlite.Open("data.db"),&gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "storage: PostgreSQL connection error")
	}
	db, err := d.DB()
	db.SetMaxOpenConns(dbCfg.MaxConn)
	db.SetMaxIdleConns(dbCfg.MaxIdle)
	db.SetConnMaxLifetime(time.Hour)

	d.AutoMigrate(&entity.Stock{})

	return &Repositories{
		db:        d,
		StockRepo: NewStockPersis(d),
	}, nil
}

// closes the database connection
func (s *Repositories) Close() error {
	return nil
}
