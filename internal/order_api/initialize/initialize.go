package initialize

import (
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/order_api/global"
	"go.uber.org/zap"
)

func Initialize() {
	ParseFlag()
	common.LoadConfig(*ConfigPath, *DevMode)
	common.NewLogger(*DevMode)
	zap.S().Infof("register to consul")
	err := common.RegisterSelfToConsul()
	if err != nil {
		zap.S().Errorw("Fail to register to consul", "error", err.Error)
	}
	global.OrderSvcConn, err = grpc_client.GetEmailSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get order svc connection", "error", err.Error)
	}
}

func Finalize() {
	global.OrderSvcConn.Close()
	err := common.DeRegisterFromConsul()
	if err != nil {
		zap.S().Errorw("Fail to deregister from consul", "error", err.Error)
	}
	zap.S().Sync()
	zap.L().Sync()
}
