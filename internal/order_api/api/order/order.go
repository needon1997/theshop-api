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

func List(ctx *gin.Context) {
	//check whether it is admin or normal user
	Iclaims, exist := ctx.Get("claims")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "access unauthorized"})
	}
	claim, _ := Iclaims.(common.JWTUserInfoClaim)
	request := &proto.OrderFilterRequest{}
	if claim.Role == USER_ROLE {
		request.UserId = int32(claim.Id)
	}
	request.PagePerNums = common.Atoi32(ctx.DefaultQuery("psize", "10"))
	request.Pages = common.Atoi32(ctx.DefaultQuery("pn", "1"))
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	orderList, err := orderClient.OrderList(context.Background(), request)
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, orderList)
	return
}
func New(ctx *gin.Context) {
	orderForm := &forms.CreateOrderForm{}
	err := validation.ValidateFormJSON(ctx, orderForm)
	if err != nil {
		return
	}
	IuserId, _ := ctx.Get("userID")
	Itoken, _ := ctx.Get("x-token")
	token := Itoken.(string)
	userId := IuserId.(int64)
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	order, err := orderClient.CreateOrder(context.Background(), &proto.OrderRequest{
		UserId:  int32(userId),
		Address: orderForm.Address,
		Mobile:  orderForm.Mobile,
		Name:    orderForm.Name,
		Note:    orderForm.Post,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	payClient := proto.NewPaymentClient(global.PaymentSvcConn)
	response, err := payClient.CreatePayment(context.Background(), &proto.CreatePaymentRequest{
		OrderId:  order.Id,
		Currency: "CAD",
		Total:    100,
		Token:    token,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"redirect": response.AcceptUrl,
	})
}
func Detail(ctx *gin.Context) {
	id := common.Atoi32(ctx.Param("id"))
	Iclaims, exist := ctx.Get("claims")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "access unauthorized"})
	}
	claim, _ := Iclaims.(common.JWTUserInfoClaim)
	request := &proto.OrderRequest{}
	request.Id = id
	if claim.Role == USER_ROLE {
		request.UserId = int32(claim.Id)
	}
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	orderDetail, err := orderClient.OrderDetail(context.Background(), request)
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, orderDetail)
	return
}
