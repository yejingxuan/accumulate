package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yejingxuan/accumulate/application"
	"github.com/yejingxuan/accumulate/infrastructure/config"
	"github.com/yejingxuan/accumulate/infrastructure/crawler"
	"github.com/yejingxuan/accumulate/infrastructure/logger"
	"github.com/yejingxuan/accumulate/infrastructure/persistence"
	"github.com/yejingxuan/accumulate/interface/api"
	"go.uber.org/zap"
	"os"
)

func Server() *cobra.Command {
	cfg := "./config.toml"
	cmdServer := &cobra.Command{
		Use:   "server",
		Short: "Start Run",
		Run: func(cmd *cobra.Command, args []string) {
			//初始化配置文件
			config.Init(cfg)
			//初始化日志
			logCfg := config.CoreConf.Log
			err := logger.InitLog(
				logger.Path(logCfg.LogPath),
				logger.Level(logCfg.LogLevel),
				logger.Compress(logCfg.Compress),
				logger.MaxSize(logCfg.MaxSize),
				logger.MaxBackups(logCfg.MaxBackups),
				logger.MaxAge(logCfg.MaxAge),
				logger.Format(logCfg.Format),
			)
			if err != nil {
				fmt.Printf("initLog failed: %v\n", err)
				os.Exit(1)
			}
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			// 初始化基础层实例 - DB实例
			persisDB, err := persistence.NewRepositories()
			if err != nil || persisDB == nil {
				logger.Fatal("DB init failed", zap.Error(err))
				os.Exit(1)
			}

			// 初始化应用层实例
			stockApp := application.NewStockApp(persisDB.StockRepo)

			crawler.ExecXueQiuJob(persisDB.StockRepo)

			//初始化http接口
			err = api.SetupHttpServer(stockApp)
			if err != nil {
				logger.Fatal("router.Run error", zap.Error(err))
				os.Exit(1)
			}
			return nil
		},
	}
	return cmdServer
}
