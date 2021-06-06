package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/proto"
	"github.com/needon1997/theshop-api/internal/common/validation"
	"github.com/needon1997/theshop-api/internal/user_api/forms"
	"github.com/needon1997/theshop-api/internal/user_api/global"
	"go.uber.org/zap"
	"net/http"
)

func SendEmail(c *gin.Context) {
	zap.S().Debug("send Email")
	client := proto.NewEmailSvcClient(global.EmailSvcConn)
	emailForm := forms.EmailForm{}
	err := validation.ValidateFormJSON(c, &emailForm)
	if err != nil {
		return
	}
	_, err = client.SendVerificationCode(context.Background(), &proto.ReceiverInfoRequest{Email: emailForm.Email})
	if err != nil {
		zap.S().Errorf("[SendEmail]   [fail to send email]  ERROR: %s", err.Error())
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "email sent",
	})
	return
}
