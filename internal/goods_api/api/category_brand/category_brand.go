package category_brand

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/proto"
	"github.com/needon1997/theshop-api/internal/common/validation"
	"github.com/needon1997/theshop-api/internal/goods_api/forms"
	"github.com/needon1997/theshop-api/internal/goods_api/global"
	"net/http"
)

func List(c *gin.Context) {
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	pn := common.Atoi32(c.DefaultQuery("pn", "0"))
	pnum := common.Atoi32(c.DefaultQuery("pnum", "10"))
	rsp, err := client.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{
		Pages:       pn,
		PagePerNums: pnum,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func Select(c *gin.Context) {
	id := common.Atoi32(c.Param("id"))
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	rsp, err := client.GetCategoryBrandList(context.Background(), &proto.CategoryInfoRequest{
		Id: id,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func New(c *gin.Context) {
	catBrandForm := forms.CatBrandForm{}
	err := validation.ValidateFormJSON(c, &catBrandForm)
	if err != nil {
		return
	}
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	rsp, err := client.CreateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		BrandId:    catBrandForm.BrandID,
		CategoryId: catBrandForm.CategoryID,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func Delete(c *gin.Context) {
	id := common.Atoi32(c.Param("id"))
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	_, err := client.DeleteCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		Id: id,
	})
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
	catBrandForm := &forms.CatBrandForm{}
	err := validation.ValidateFormJSON(c, &catBrandForm)
	if err != nil {
		return
	}
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	_, err = client.UpdateCategoryBrand(context.Background(), &proto.CategoryBrandRequest{
		Id:         id,
		BrandId:    catBrandForm.BrandID,
		CategoryId: catBrandForm.CategoryID,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
}
