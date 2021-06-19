package address

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
	request := &proto.AddressRequest{}
	Iclaim, _ := c.Get("claims")
	claim := Iclaim.(*common.JWTUserInfoClaim)
	if claim.Role != 2 {
		IuserId, _ := c.Get("userID")
		userId := IuserId.(int64)
		request.UserId = int32(userId)
	}
	addressClient := proto.NewAddressClient(global.UserOpSvcConn)
	addressList, err := addressClient.GetAddressList(spanCtx, request)
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, addressList)
	return
}

func Delete(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	id := common.Atoi32(c.Param("id"))
	addressClient := proto.NewAddressClient(global.UserOpSvcConn)
	_, err := addressClient.DeleteAddress(spanCtx, &proto.AddressRequest{Id: id})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete success",
	})
}

func New(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	addressForm := forms.AddressForm{}
	err := validation.ValidateFormJSON(c, &addressForm)
	if err != nil {
		return
	}
	IuserId, _ := c.Get("userID")
	userId := IuserId.(int64)
	addressClient := proto.NewAddressClient(global.UserOpSvcConn)
	address, err := addressClient.CreateAddress(spanCtx, &proto.AddressRequest{
		UserId:       int32(userId),
		Province:     addressForm.Province,
		City:         addressForm.City,
		Address:      addressForm.Address,
		SignerName:   addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, address)
	return
}

func Edit(c *gin.Context) {
	ispanCtx, _ := c.Get("ctx")
	spanCtx := ispanCtx.(context.Context)
	addressForm := forms.AddressForm{}
	err := validation.ValidateFormJSON(c, addressForm)
	if err != nil {
		return
	}
	id := common.Atoi32(c.Param("id"))
	addressClient := proto.NewAddressClient(global.UserOpSvcConn)
	_, err = addressClient.UpdateAddress(spanCtx, &proto.AddressRequest{
		Id:           id,
		Province:     addressForm.Province,
		City:         addressForm.City,
		Address:      addressForm.Address,
		SignerName:   addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})
	if err != nil {
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "update success",
	})
}
