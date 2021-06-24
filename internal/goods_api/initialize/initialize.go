package initialize

import (
	"github.com/go-playground/validator/v10"
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/goods_api/global"
	"github.com/needon1997/theshop-api/internal/goods_api/validators"
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
	RegisterValidator(map[string]validator.Func{
		"mobile": validators.ValidateMobile,
	})
	global.GoodsSvcConn, err = grpc_client.GetGoodsSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get goods svc connection", "error", err)
	}
}

func Finalize() {
	zap.S().Info("deregister from consul")
	traceCloser.Close()
	global.GoodsSvcConn.Close()
	zap.S().Sync()
	zap.L().Sync()
}
