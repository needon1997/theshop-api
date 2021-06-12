package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func RegisterValidator(registerMap map[string]validator.Func) {
	zap.S().Infow("Initializing validators...")
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		for k, f := range registerMap {
			zap.S().Infof("Register validator for tag:%s", k)
			v.RegisterValidation(k, f)
		}
	} else {
		zap.S().Error("Validator Engine Request Error")
	}
}
