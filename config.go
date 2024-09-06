package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		logger.DPanic("Failed to read config file", zap.Error(err))
		os.Exit(1)
	}
}
