package persistence

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"op-register/domain/repository"
	"op-register/infrastructure/config"
	"time"
)

// Repositories 总仓储机构提，包含多个领域仓储接口，以及一个DB实例
type Repositories struct {
	AdjunctRepo      repository.AdjunctRepo
	CategoryRepo     repository.CategoryRepo
	EnvRepo          repository.EnvRepo
	ImageRepo        repository.ImageRepo
	LogInfoRepo      repository.LogInfoRepo
	OperatorRepo     repository.OperatorRepo
	OperatorTempRepo repository.OperatorTempRepo
	ServiceRepo      repository.ServiceRepo
	TagRepo          repository.TagsRepo
	db               *gorm.DB
}

// NewRepositories 初始化所有域的总仓储实例，将实例通过依赖注入方式，将DB实例注入到领域层
func NewRepositories() (*Repositories, error) {
	dbCfg := config.CoreConf.Server.DB
	d, err := gorm.Open("postgres", dbCfg.Dsn)
	if err != nil {
		return nil, errors.Wrap(err, "storage: PostgreSQL connection error")
	}

	d.DB().SetMaxOpenConns(dbCfg.MaxConn)
	d.DB().SetMaxIdleConns(dbCfg.MaxIdle)
	d.DB().SetConnMaxLifetime(time.Hour)
	d.LogMode(dbCfg.LogMode)

	return &Repositories{
		AdjunctRepo:      NewAdjunctPersis(d),
		CategoryRepo:     NewCategoryPersis(d),
		EnvRepo:          NewEnvPersis(d),
		ImageRepo:        NewImagePersis(d),
		LogInfoRepo:      NewLogInfoPersis(d),
		OperatorRepo:     NewOperatorPersis(d),
		OperatorTempRepo: NewOperatorTempPersis(d),
		ServiceRepo:      NewServicePersis(d),
		TagRepo:          NewTagPersis(d),
		db:               d,
	}, nil
}

// closes the database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}
