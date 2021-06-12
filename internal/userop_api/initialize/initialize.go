package initialize

import (
	"github.com/go-playground/validator/v10"
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/validation"
	"github.com/needon1997/theshop-api/internal/user_api/global"
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
	RegisterValidator(map[string]validator.Func{
		"mobile": validation.ValidateMobile,
	})
	global.UserSvcConn, err = grpc_client.GetUserSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get user svc connection", "error", err.Error)
	}
	global.EmailSvcConn, err = grpc_client.GetEmailSvcConn()
	if err != nil {
		zap.S().Errorw("Fail to get email svc connection", "error", err.Error)
	}
}

func Finalize() {
	global.EmailSvcConn.Close()
	global.UserSvcConn.Close()
	err := common.DeRegisterFromConsul()
	if err != nil {
		zap.S().Errorw("Fail to deregister from consul", "error", err.Error)
	}
	zap.S().Sync()
	zap.L().Sync()
}
