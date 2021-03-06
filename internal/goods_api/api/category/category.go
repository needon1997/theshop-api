package category

import (
	"context"
	"encoding/json"
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
	rsp, err := client.GetAllCategorysList(spanCtx, &empty.Empty{})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	cat := make([]interface{}, 0)
	json.Unmarshal([]byte(rsp.JsonData), &cat)
	c.JSON(http.StatusOK, cat)
}

func New(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	categoryForm := &forms.CategoryForm{}
	err := validation.ValidateFormJSON(c, &categoryForm)
	if err != nil {
		return
	}
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	rsp, err := client.CreateCategory(spanCtx, &proto.CategoryInfoRequest{
		Name:           categoryForm.Name,
		ParentCategory: categoryForm.Parent,
		Level:          categoryForm.Level,
		IsTab:          *categoryForm.IsTab,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func Detail(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	id := common.Atoi32(c.Param("id"))
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	rsp, err := client.GetSubCategory(spanCtx, &proto.CategoryListRequest{
		Id: id,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
	}
	c.JSON(http.StatusOK, rsp)
	return
}

func Delete(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	id := common.Atoi32(c.Param("id"))
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	_, err := client.DeleteCategory(spanCtx, &proto.DeleteCategoryRequest{Id: id})
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
	categoryUpdateForm := forms.CategoryUpdateForm{}
	err := validation.ValidateFormJSON(c, &categoryUpdateForm)
	if err != nil {
		return
	}
	client := proto.NewGoodsClient(global.GoodsSvcConn)
	_, err = client.UpdateCategory(spanCtx, &proto.CategoryInfoRequest{
		Id:    id,
		Name:  categoryUpdateForm.Name,
		IsTab: *categoryUpdateForm.IsTab,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
}
