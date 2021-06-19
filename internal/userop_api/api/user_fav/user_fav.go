package user_fav

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
	IuserId, _ := c.Get("userID")
	userId := IuserId.(int64)
	userFavClient := proto.NewUserFavClient(global.UserOpSvcConn)
	list, err := userFavClient.GetFavList(spanCtx, &proto.UserFavRequest{
		UserId: int32(userId),
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	goodsId := make([]int32, 0)
	for i := 0; i < int(list.Total); i++ {
		goodsId = append(goodsId, list.Data[i].GoodsId)
	}
	goodsClient := proto.NewGoodsClient(global.GoodsSvcConn)
	goods, err := goodsClient.BatchGetGoods(spanCtx, &proto.BatchGoodsIdInfo{Id: goodsId})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	userFavGoods := make([]interface{}, 0)
	for i := 0; i < int(list.Total); i++ {
		for j := 0; j < int(goods.Total); j++ {
			if list.Data[i].GoodsId == goods.Data[j].Id {
				userFavGoods = append(userFavGoods, map[string]interface{}{
					"id":    list.Data[i].GoodsId,
					"name":  goods.Data[j].Name,
					"price": goods.Data[j].ShopPrice,
					"image": goods.Data[j].GoodsFrontImage,
				})
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"total": goods.Total,
		"data":  userFavGoods,
	})
	return

}

func Delete(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	IuserId, _ := c.Get("userID")
	userId := IuserId.(int64)
	goodsId := common.Atoi32(c.Param("id"))
	userFavClient := proto.NewUserFavClient(global.UserOpSvcConn)
	_, err := userFavClient.DeleteUserFav(spanCtx, &proto.UserFavRequest{UserId: int32(userId), GoodsId: goodsId})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "Delete favorite success",
	})
}

func New(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	IuserId, _ := c.Get("userID")
	userId := IuserId.(int64)
	userFavForm := forms.UserFavForm{}
	err := validation.ValidateFormJSON(c, &userFavForm)
	if err != nil {
		return
	}
	userFavClient := proto.NewUserFavClient(global.UserOpSvcConn)
	_, err = userFavClient.AddUserFav(spanCtx, &proto.UserFavRequest{UserId: int32(userId), GoodsId: userFavForm.GoodsId})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "add favorite success",
	})
	return
}

func Detail(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	IuserId, _ := c.Get("userID")
	userId := IuserId.(int64)
	goodsId := common.Atoi32(c.Param("id"))
	userFavClient := proto.NewUserFavClient(global.UserOpSvcConn)
	_, err := userFavClient.GetUserFavDetail(spanCtx, &proto.UserFavRequest{UserId: int32(userId), GoodsId: goodsId})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.Status(http.StatusOK)
	return
}
