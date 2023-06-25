package config

import (
	"github.com/spf13/viper"
	"os"
)

// Initialize 触发加载 config 包的所有 init 函数
func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
