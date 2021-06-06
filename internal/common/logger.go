package common

import (
	"github.com/needon1997/theshop-api/internal/common/config"
	"go.uber.org/zap"
)

func NewLogger(devMode bool) {
	logPath := config.ServerConfig.LogConfig.LogPath
	var logger *zap.Logger
	var err error
	if devMode {
		logger, err = NewDevLogger(logPath)
	} else {
		logger, err = NewProductLogger(logPath)
	}
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}

func NewDevLogger(logPath string) (*zap.Logger, error) {
	if logPath == "" {
		return zap.NewDevelopment()
	}
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{logPath, "stdout"}
	return config.Build()
}

func NewProductLogger(logPath string) (*zap.Logger, error) {
	if logPath == "" {
		return zap.NewProduction()
	}
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{logPath, "stderr"}
	return config.Build()

}
