package message

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/proto"
	"github.com/needon1997/theshop-api/internal/common/validation"
	"github.com/needon1997/theshop-api/internal/userop_api/forms"
	"github.com/needon1997/theshop-api/internal/userop_api/global"
	"net/http"
)

func List(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	request := &proto.MessageRequest{}
	Iclaim, _ := c.Get("claims")
	claim := Iclaim.(*common.JWTUserInfoClaim)
	if claim.Role != 2 {
		IuserId, _ := c.Get("userID")
		userId := IuserId.(int64)
		request.UserId = int32(userId)
	}
	messageClient := proto.NewMessageClient(global.UserOpSvcConn)
	messageList, err := messageClient.MessageList(spanCtx, request)
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, messageList)
	return
}

func New(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	IuserId, _ := c.Get("userID")
	userId := IuserId.(int64)
	messageForm := forms.MessageForm{}
	err := validation.ValidateFormJSON(c, &messageForm)
	if err != nil {
		return
	}
	messageClient := proto.NewMessageClient(global.UserOpSvcConn)
	response, err := messageClient.CreateMessage(spanCtx, &proto.MessageRequest{
		UserId:      int32(userId),
		MessageType: messageForm.MessageType,
		Subject:     messageForm.Subject,
		Message:     messageForm.Message,
		File:        messageForm.File,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, response)
	return

}
