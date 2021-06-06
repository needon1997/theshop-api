package brand

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
	rsp, err := client.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       pn,
		PagePerNums: pnum,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func New(c *gin.Context) {
	brandForm := &forms.BrandForm{}
	err := validation.ValidateFormJSON(c, &brandForm)
	if err != nil {
		return
	}
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	rsp, err := client.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: brandForm.Name,
		Logo: brandForm.Logo,
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
	_, err := client.DeleteBrand(context.Background(), &proto.BrandRequest{Id: id})
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
	brandForm := &forms.BrandForm{}
	err := validation.ValidateFormJSON(c, &brandForm)
	if err != nil {
		return
	}
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	_, err = client.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id:   id,
		Name: brandForm.Name,
		Logo: brandForm.Logo,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
}
