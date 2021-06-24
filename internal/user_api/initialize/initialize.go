package initialize

import (
	"github.com/go-playground/validator/v10"
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/validation"
	"github.com/needon1997/theshop-api/internal/user_api/global"
	"go.uber.org/zap"
	"io"
)

var TraceCloser io.Closer

func Initialize() {
	var err error
	ParseFlag()
	common.LoadConfig(*ConfigPath, *DevMode)
	common.NewLogger(*DevMode)
	RegisterValidator(map[string]validator.Func{
		"mobile": validation.ValidateMobile,
	})
	TraceCloser = common.InitJaeger()
	global.UserSvcConn, err = grpc_client.GetUserSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get user svc connection", "error", err)
	}
	global.EmailSvcConn, err = grpc_client.GetEmailSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get email svc connection", "error", err)
	}
}

func Finalize() {
	global.EmailSvcConn.Close()
	global.UserSvcConn.Close()
	TraceCloser.Close()
	zap.S().Sync()
	zap.L().Sync()
}
