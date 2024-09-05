package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logger.DPanic("Failed to read config file", zap.Error(err))
	}
}
