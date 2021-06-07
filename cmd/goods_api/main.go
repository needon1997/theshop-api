package main

import (
	"fmt"
	"github.com/needon1997/theshop-api/internal/common/config"
	"github.com/needon1997/theshop-api/internal/goods_api/initialize"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	initialize.Initialize()
	defer initialize.Finalize()
	if config.ServerConfig.Host == "" {
		zap.S().Panic("host should not be empty")
	}
	if config.ServerConfig.Port == 0 {
		zap.S().Panic("port should not be empty")
	}
	engine := initialize.InitializeRouter()
	zap.S().Debugf("server start, host: %s:%v", config.ServerConfig.Host, config.ServerConfig.Port)
	go func() {
		err := engine.Run(fmt.Sprintf("%s:%v", config.ServerConfig.Host, config.ServerConfig.Port))
		if err != nil {
			zap.S().Panic("server start error:", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Infow("shut down sever")
}
