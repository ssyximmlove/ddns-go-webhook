package main

import "go.uber.org/zap"

var logger *zap.Logger

func InitLogger() {
	zapCfg := zap.NewProductionConfig()
	zapCfg.OutputPaths = []string{"stdout", "./logs.log"}
	zapCfg.DisableStacktrace = true
	logger, _ = zapCfg.Build()
}
