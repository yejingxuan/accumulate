package main

import (
	"github.com/spf13/cobra"
	"github.com/yejingxuan/accumulate/cmd"
	_ "github.com/yejingxuan/accumulate/docs"
	"log"
)

// @title accumulate服务API接口
// @version 1.0.0
// @description accumulate服务API接口
func main() {
	log.Println("accumulate====version: v1.0.0")
	log.Println("accumulate api doc===http://127.0.0.1:16666/accumulate/v1/swagger/index.html")
	rootCmd := &cobra.Command{Use: "accumulate"}
	rootCmd.AddCommand(cmd.Server())
	if err := rootCmd.Execute(); err != nil {
		log.Println("rootCmd.Execute failed", err.Error())
	}
}
