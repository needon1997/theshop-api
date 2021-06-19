package payment

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/proto"
	"github.com/needon1997/theshop-api/internal/order_api/global"
	"net/http"
	"time"
)

const ORDER_TIMEOUT int = 890

func Execute(c *gin.Context) {
	ispanCtx, _ := c.Get("c")
	spanCtx := ispanCtx.(context.Context)
	orderSn := c.Param("ordersn")
	IuserId, _ := c.Get("userID")
	userId := IuserId.(int64)
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	orderInfo, err := orderClient.OrderDetail(spanCtx, &proto.OrderRequest{OrderSn: orderSn, UserId: int32(userId)})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	addTime, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", orderInfo.OrderInfo.AddTime)
	if orderInfo.OrderInfo.Status != "paying" || time.Now().Second()-addTime.Second() > ORDER_TIMEOUT {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "order timeout",
		})
		return
	}
	paymentId := c.Query("paymentId")
	payerId := c.Query("PayerID")
	paymentClient := proto.NewPaymentClient(global.PaymentSvcConn)
	_, err = paymentClient.ExecutePayment(spanCtx, &proto.ExecutePaymentRequest{
		PaymentId: paymentId,
		PayerId:   payerId,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	_, err = orderClient.UpdateOrderStatus(spanCtx, &proto.OrderStatus{OrderSn: orderSn, Status: "approved"})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "payment success",
	})
}
