package banner

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/proto"
	"github.com/needon1997/theshop-api/internal/common/validation"
	"github.com/needon1997/theshop-api/internal/goods_api/forms"
	"github.com/needon1997/theshop-api/internal/goods_api/global"
	"net/http"
)

func List(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	rsp, err := client.BannerList(spanCtx, &empty.Empty{})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func New(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	bannerForm := &forms.BannerForm{}
	err := validation.ValidateFormJSON(c, &bannerForm)
	if err != nil {
		return
	}
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	rsp, err := client.CreateBanner(spanCtx, &proto.BannerRequest{
		Url:   bannerForm.Url,
		Index: bannerForm.Index,
		Image: bannerForm.Image,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func Delete(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	id := common.Atoi32(c.Param("id"))
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	_, err := client.DeleteBanner(spanCtx, &proto.BannerRequest{Id: id})
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
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	id := common.Atoi32(c.Param("id"))
	bannerForm := forms.BannerForm{}
	err := validation.ValidateFormJSON(c, &bannerForm)
	if err != nil {
		return
	}
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	_, err = client.UpdateBanner(spanCtx, &proto.BannerRequest{
		Id:    id,
		Index: bannerForm.Index,
		Image: bannerForm.Image,
		Url:   bannerForm.Url,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
}
