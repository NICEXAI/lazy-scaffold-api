package cmd

import (
	"github.com/spf13/cobra"
	"lazy-scaffold-api/internal/server"
	"os"
)

var (
	cfgFile string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./settings.default.yml", "specify config file path")
}

var rootCmd = &cobra.Command{
	// lazy replace:name>lazy-scaffold-api range:1
	Use: "lazy-scaffold-api",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// 主要用于初始化服务相关的基础模块
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
