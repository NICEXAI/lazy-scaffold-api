package config

import (
	"os"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Info Options

func init() {
	currentPath, _ := os.Getwd()
	configName := "settings.yaml"

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(currentPath)
	viper.AddConfigPath(path.Join(currentPath, "./config"))

	viper.OnConfigChange(func(in fsnotify.Event) {
		loadConfig()
	})

	loadConfig()
}

func loadConfig() {
	if err := viper.ReadInConfig(); err != nil {
		panic("基础配置文件加载失败，请检查/config文件是否设置正确。")
	}

	if err := viper.Unmarshal(&Info); err != nil {
		panic("基础配置文件解析失败：" + err.Error())
	}
}
