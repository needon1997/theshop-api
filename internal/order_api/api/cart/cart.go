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

func List(c *gin.Context) {
	ispanCtx, _ := c.Get("c")
	spanCtx := ispanCtx.(context.Context)
	uidStr, _ := c.Get("userID")
	uid, _ := uidStr.(int64)
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	cartItemList, err := orderClient.CartItemList(spanCtx, &proto.UserInfo{Id: int32(uid)})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	goodsId := make([]int32, 0)
	for i := 0; i < int(cartItemList.Total); i++ {
		goodsId = append(goodsId, cartItemList.Data[i].GoodsId)
	}
	if len(goodsId) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"total": 0,
		})
		return
	}
	goodClient := proto.NewGoodsClient(global.GoodsSvcConn)
	goodsList, err := goodClient.BatchGetGoods(spanCtx, &proto.BatchGoodsIdInfo{Id: goodsId})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
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
			"goods_price": goodsList.Data[i].ShopPrice,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"total": goodsList.Total,
		"data":  rspData,
	})
	return
}
func Delete(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	uidStr, _ := c.Get("userID")
	uid, _ := uidStr.(int64)
	gid := common.Atoi32(c.Param("id"))
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	_, err := orderClient.DeleteCartItem(spanCtx, &proto.CartItemRequest{UserId: int32(uid), GoodsId: gid})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete success",
	})
	return
}
func New(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	uidStr, _ := c.Get("userID")
	uid, _ := uidStr.(int64)
	cartItemForm := &forms.CartItemCreateForm{}
	err := validation.ValidateFormJSON(c, cartItemForm)
	if err != nil {
		return
	}
	goodsClient := proto.NewGoodsClient(global.GoodsSvcConn)
	//check existence of the goods
	_, err = goodsClient.GetGoodsDetail(spanCtx, &proto.GoodInfoRequest{Id: cartItemForm.GoodsId})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	inventoryClient := proto.NewInventoryClient(global.InventoryConn)
	//check availability of the goods
	invDetail, err := inventoryClient.InvDetail(spanCtx, &proto.GoodsInvInfo{GoodsId: cartItemForm.GoodsId})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	if invDetail.Num < cartItemForm.Nums {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "insufficient inventory",
		})
		return
	}
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	response, err := orderClient.CreateCartItem(spanCtx, &proto.CartItemRequest{UserId: int32(uid), GoodsId: cartItemForm.GoodsId, Nums: cartItemForm.Nums})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, response)
	return
}

func Edit(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	uidStr, _ := c.Get("userID")
	uid, _ := uidStr.(int64)
	cartItemForm := &forms.CartItemUpdateForm{}
	err := validation.ValidateFormJSON(c, cartItemForm)
	if err != nil {
		return
	}
	orderClient := proto.NewOrderClient(global.OrderSvcConn)
	_, err = orderClient.UpdateCartItem(spanCtx, &proto.CartItemRequest{
		UserId:  int32(uid),
		GoodsId: cartItemForm.GoodsId,
		Nums:    cartItemForm.Nums,
		Checked: *cartItemForm.Checked,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
	return
}
