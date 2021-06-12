package payment

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/proto"
	"github.com/needon1997/theshop-api/internal/order_api/global"
	"net/http"
)

func Execute(ctx *gin.Context) {
	id := common.Atoi32(ctx.Param("id"))
	IuserId, _ := ctx.Get("userID")
	userId := IuserId.(int64)
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	_, err := orderClient.OrderDetail(context.Background(), &proto.OrderRequest{Id: id, UserId: int32(userId)})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	paymentId := ctx.Query("paymentId")
	payerId := ctx.Query("PayerID")
	paymentClient := proto.NewPaymentClient(global.PaymentSvcConn)
	response, err := paymentClient.ExecutePayment(context.Background(), &proto.ExecutePaymentRequest{
		PaymentId: paymentId,
		PayerId:   payerId,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	_, err = orderClient.UpdateOrderStatus(context.Background(), &proto.OrderStatus{OrderId: id, Status: response.Status})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "payment success",
	})
}
