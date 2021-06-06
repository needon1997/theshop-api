package goods

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/proto"
	"github.com/needon1997/theshop-api/internal/common/validation"
	"github.com/needon1997/theshop-api/internal/goods_api/forms"
	"github.com/needon1997/theshop-api/internal/goods_api/global"
	"go.uber.org/zap"
	"net/http"
)

func List(c *gin.Context) {
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	req := &proto.GoodsFilterRequest{}
	req.PriceMin = common.Atoi32(c.DefaultQuery("pmin", "0"))
	req.PriceMax = common.Atoi32(c.DefaultQuery("pmax", "0"))
	if c.DefaultQuery("ih", "0") == "1" {
		req.IsHot = true
	}
	if c.DefaultQuery("in", "0") == "1" {
		req.IsNew = true
	}
	if c.DefaultQuery("it", "0") == "1" {
		req.IsTab = true
	}
	req.TopCategory = common.Atoi32(c.DefaultQuery("c", "0"))
	req.Brand = common.Atoi32(c.DefaultQuery("b", "0"))
	req.PagePerNums = common.Atoi32(c.DefaultQuery("pnum", "10"))
	req.Pages = common.Atoi32(c.DefaultQuery("pn", "1"))
	zap.S().Debug("invoke gRPC GoodsList service")
	rsp, err := client.GoodsList(context.Background(), req)
	if err != nil {
		zap.S().Errorw("[List] Get GoodsList failure", "error", err.Error())
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func New(c *gin.Context) {
	goodsForm := forms.GoodsForm{}
	err := validation.ValidateFormJSON(c, &goodsForm)
	if err != nil {
		return
	}
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	rsp, err := client.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		GoodsDesc:       goodsForm.GoodsDesc,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.Brand,
		IsNew:           *goodsForm.IsNew,
		IsHot:           *goodsForm.IsHot,
		OnSale:          *goodsForm.OnSale,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
	return
}

func Detail(c *gin.Context) {
	id := common.Atoi32(c.Param("id"))
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	rsp, err := client.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{Id: id})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	//TODO Iventory info
	c.JSON(http.StatusOK, rsp)
	return
}

func Delete(c *gin.Context) {
	id := common.Atoi32(c.Param("id"))
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	_, err := client.DeleteGoods(context.Background(), &proto.DeleteGoodsInfo{Id: id})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete success",
	})
	return
}

func Update(c *gin.Context) {
	id := common.Atoi32(c.Param("id"))
	goodsForm := forms.GoodsForm{}
	err := validation.ValidateFormJSON(c, &goodsForm)
	if err != nil {
		return
	}
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	_, err = client.UpdateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:              id,
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		GoodsDesc:       goodsForm.GoodsDesc,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.Brand,
		IsNew:           *goodsForm.IsNew,
		IsHot:           *goodsForm.IsHot,
		OnSale:          *goodsForm.OnSale,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
}

func Stocks(c *gin.Context) {
	_ = common.Atoi32(c.Param("id"))
	_ = proto.NewGoodsClient(global.GoodsSvcConn)
	//TODO get inventory
}
