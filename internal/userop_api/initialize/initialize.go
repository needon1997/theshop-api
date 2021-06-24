package initialize

import (
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/userop_api/global"
	"go.uber.org/zap"
	"io"
)

var TraceCloser io.Closer

func Initialize() {
	var err error
	ParseFlag()
	common.LoadConfig(*ConfigPath, *DevMode)
	common.NewLogger(*DevMode)
	TraceCloser = common.InitJaeger()
	global.UserOpSvcConn, err = grpc_client.GetUserOpSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get userop svc connection", "error", err.Error)
	}
	global.GoodsSvcConn, err = grpc_client.GetGoodsSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get goods svc connection", "error", err.Error)
	}
}

func Finalize() {
	global.UserOpSvcConn.Close()
	global.GoodsSvcConn.Close()
	TraceCloser.Close()
	zap.S().Sync()
	zap.L().Sync()
}
