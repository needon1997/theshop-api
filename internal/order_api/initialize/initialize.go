package initialize

import (
	"github.com/go-playground/validator/v10"
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/validation"
	"github.com/needon1997/theshop-api/internal/order_api/global"
	"go.uber.org/zap"
	"io"
)

var traceCloser io.Closer

func Initialize() {
	var err error
	ParseFlag()
	common.LoadConfig(*ConfigPath, *DevMode)
	common.NewLogger(*DevMode)
	traceCloser = common.InitJaeger()
	global.OrderSvcConn, err = grpc_client.GetOrderSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get order svc connection", "error", err.Error)
	}
	global.GoodsSvcConn, err = grpc_client.GetGoodsSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get goods svc connection", "error", err.Error)
	}
	global.InventoryConn, err = grpc_client.GetInventorySvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get inventory svc connection", "error", err.Error)
	}
	global.PaymentSvcConn, err = grpc_client.GetPaymentSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get payment svc connection", "error", err.Error)
	}
	RegisterValidator(map[string]validator.Func{"mobile": validation.ValidateMobile})
}

func Finalize() {
	global.OrderSvcConn.Close()
	global.GoodsSvcConn.Close()
	global.PaymentSvcConn.Close()
	global.InventoryConn.Close()
	traceCloser.Close()
	zap.S().Sync()
	zap.L().Sync()
}
