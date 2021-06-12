package cart

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

func List(ctx *gin.Context) {
	uidStr, _ := ctx.Get("userID")
	uid, _ := uidStr.(int64)
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	cartItemList, err := orderClient.CartItemList(context.Background(), &proto.UserInfo{Id: int32(uid)})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	goodsId := make([]int32, 0)
	for i := 0; i < int(cartItemList.Total); i++ {
		goodsId = append(goodsId, cartItemList.Data[i].GoodsId)
	}
	goodClient := proto.NewGoodsClient(global.GoodsSvcConn)
	goodsList, err := goodClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{Id: goodsId})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	rspData := make([]map[string]interface{}, 0)
	for i := 0; i < int(goodsList.Total); i++ {
		rspData = append(rspData, map[string]interface{}{
			"id":          cartItemList.Data[i].Id,
			"goods_id":    cartItemList.Data[i].GoodsId,
			"goods_name":  goodsList.Data[i].Name,
			"goods_image": goodsList.Data[i].GoodsFrontImage,
			"num":         cartItemList.Data[i].Nums,
			"checked":     cartItemList.Data[i].Checked,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"total": goodsList.Total,
		"data":  rspData,
	})
	return
}
func Delete(ctx *gin.Context) {
	uidStr, _ := ctx.Get("userID")
	uid, _ := uidStr.(int64)
	gid := common.Atoi32(ctx.Param("gid"))
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	_, err := orderClient.DeleteCartItem(context.Background(), &proto.CartItemRequest{UserId: int32(uid), GoodsId: gid})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "delete success",
	})
	return
}
func New(ctx *gin.Context) {
	uidStr, _ := ctx.Get("userID")
	uid, _ := uidStr.(int64)
	cartItemForm := &forms.CartItemCreateForm{}
	err := validation.ValidateFormJSON(ctx, cartItemForm)
	if err != nil {
		return
	}
	goodsClient := proto.NewGoodsClient(global.GoodsSvcConn)
	//check existence of the goods
	_, err = goodsClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{Id: cartItemForm.GoodsId})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	inventoryClient := proto.NewInventoryClient(global.InventoryConn)
	//check availability of the goods
	invDetail, err := inventoryClient.InvDetail(context.Background(), &proto.GoodsInvInfo{GoodsId: cartItemForm.GoodsId})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	if invDetail.Num < cartItemForm.Nums {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "insufficient inventory",
		})
		return
	}
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	response, err := orderClient.CreateCartItem(context.Background(), &proto.CartItemRequest{UserId: int32(uid), GoodsId: cartItemForm.GoodsId, Nums: cartItemForm.Nums})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, response)
	return
}

func Edit(ctx *gin.Context) {
	uidStr, _ := ctx.Get("userID")
	uid, _ := uidStr.(int64)
	cartItemForm := &forms.CartItemUpdateForm{}
	err := validation.ValidateFormJSON(ctx, cartItemForm)
	if err != nil {
		return
	}
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	_, err = orderClient.UpdateCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  int32(uid),
		GoodsId: cartItemForm.GoodsId,
		Nums:    cartItemForm.Nums,
		Checked: *cartItemForm.Checked,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
	return
}
