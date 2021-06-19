package order

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/proto"
	"github.com/needon1997/theshop-api/internal/common/validation"
	"github.com/needon1997/theshop-api/internal/order_api/forms"
	"github.com/needon1997/theshop-api/internal/order_api/global"
	"net/http"
)

const (
	USER_ROLE  uint8 = 1
	ADMIN_ROLE uint8 = 2
)

func List(c *gin.Context) {
	ispanCtx, _ := c.Get("c")
	spanCtx := ispanCtx.(context.Context)
	//check whether it is admin or normal user
	Iclaims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access unauthorized"})
	}
	claim, _ := Iclaims.(common.JWTUserInfoClaim)
	request := &proto.OrderFilterRequest{}
	if claim.Role == USER_ROLE {
		request.UserId = int32(claim.Id)
	}
	request.PagePerNums = common.Atoi32(c.DefaultQuery("psize", "10"))
	request.Pages = common.Atoi32(c.DefaultQuery("pn", "1"))
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	orderList, err := orderClient.OrderList(spanCtx, request)
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, orderList)
	return
}
func New(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	orderForm := &forms.CreateOrderForm{}
	err := validation.ValidateFormJSON(c, orderForm)
	if err != nil {
		return
	}
	IuserId, _ := c.Get("userID")
	Itoken, _ := c.Get("x-token")
	token := Itoken.(string)
	userId := IuserId.(int64)
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	order, err := orderClient.CreateOrder(spanCtx, &proto.OrderRequest{
		UserId:  int32(userId),
		Address: orderForm.Address,
		Mobile:  orderForm.Mobile,
		Name:    orderForm.Name,
		Note:    orderForm.Post,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	payClient := proto.NewPaymentClient(global.PaymentSvcConn)
	response, err := payClient.CreatePayment(spanCtx, &proto.CreatePaymentRequest{
		OrderSn:  order.OrderSn,
		Currency: "CAD",
		Total:    int32(order.Total * 100),
		Token:    token,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"redirect": response.AcceptUrl,
	})
}
func Detail(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	orderSn := c.Param("ordersn")
	Iclaims, exist := c.Get("claims")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access unauthorized"})
	}
	claim, _ := Iclaims.(common.JWTUserInfoClaim)
	request := &proto.OrderRequest{}
	request.OrderSn = orderSn
	if claim.Role == USER_ROLE {
		request.UserId = int32(claim.Id)
	}
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	orderDetail, err := orderClient.OrderDetail(spanCtx, request)
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, orderDetail)
	return
}
