package cmd

import (
	"errors"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"lazy-scaffold-api/internal/config"
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
		// 初始化系统配置
		if err := config.Setup(cfgFile); err != nil {
			return errors.New(color.RedString("系统配置初始化失败: %v", err))
		}
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
