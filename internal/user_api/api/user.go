package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common/grpc_client"
	"github.com/needon1997/theshop-api/internal/common/proto"
	"github.com/needon1997/theshop-api/internal/common/validation"
	"github.com/needon1997/theshop-api/internal/user_api/forms"
	"github.com/needon1997/theshop-api/internal/user_api/global"
	"github.com/needon1997/theshop-api/internal/user_api/utils"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

const ERROR string = "error"
const MSG string = "msg"
const TOKEN string = "token"

func GetUserList(c *gin.Context) {
	zap.S().Debug("Get users list")
	client := proto.NewUserSVCClient(global.UserSvcConn)
	pn := c.DefaultQuery("pn", "0")
	pageNumber, err := strconv.Atoi(pn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}
	psize := c.DefaultQuery("psize", "10")
	pageSize, err := strconv.Atoi(psize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err.Error(),
		})
		return
	}
	zap.S().Debug("invoke gRPC GetUserList service")
	rsp, err := client.GetUserList(context.Background(), &proto.PageInfoRequest{
		PageSize:   uint32(pageSize),
		PageNumber: uint32(pageNumber),
	})
	if err != nil {
		zap.S().Errorf("[GetUserList]   [fail to get user list]  ERROR: %s", err.Error())
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func UserPasswordLogin(c *gin.Context) {
	zap.S().Debug("User Login")
	client := proto.NewUserSVCClient(global.UserSvcConn)
	emailPasswordForm := forms.EmailPasswordLoginForm{}
	err := validation.ValidateFormJSON(c, &emailPasswordForm)
	if err != nil {
		return
	}
	zap.S().Debug("invoke gRPC getUserByEmail service")
	userInfo, err := client.GetUserByEmail(context.Background(), &proto.EmailRequest{Email: emailPasswordForm.Email})
	if err != nil {
		zap.S().Errorf("[UserPasswordLogin]   [fail to get user by mobile]  ERROR: %s", err.Error())
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	zap.S().Debug("invoke gRPC getUserByEmail service success")
	zap.S().Debug("invoke gRPC compare password service")
	rsp, err := client.ComparePassword(context.Background(), &proto.ComparePasswordRequest{EncryptPwd: userInfo.Password, Password: emailPasswordForm.Password})
	if err != nil {
		zap.S().Errorf("[UserPasswordLogin]   [fail to validate password]  ERROR: %s", err.Error())
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	if rsp.Result {
		token := utils.NewJwtToken(utils.JWTUserInfoClaim{Id: int64(userInfo.Id), Nickname: userInfo.NickName, Role: uint8(userInfo.Role), StandardClaims: utils.GetStandardClaim()})
		c.JSON(http.StatusOK, gin.H{
			MSG:   "login success",
			TOKEN: token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			MSG: "password incorrect",
		})
	}
}

func Register(c *gin.Context) {
	zap.S().Debug("User Register")
	registerForm := forms.RegisterForm{}
	err := validation.ValidateFormJSON(c, &registerForm)
	if err != nil {
		return
	}
	emailClient := proto.NewEmailSvcClient(global.EmailSvcConn)
	response, err := emailClient.VerifyVerificationCode(context.Background(), &proto.VerifyCodeRequest{Email: registerForm.Email, Code: registerForm.Code})
	if err != nil {
		zap.S().Errorf("[UserRegister]   [fail to validate]  ERROR: %s", err.Error())
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	if !response.Match {
		zap.S().Debug("[UserRegister]   [verification code not match]")
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "verification code error",
		})
		return
	}
	usrClient := proto.NewUserSVCClient(global.UserSvcConn)
	zap.S().Debug("invoke gRPC registerUser service")
	result, err := usrClient.CreateUser(context.Background(), &proto.CreateUserInfoRequest{Email: registerForm.Email, Password: registerForm.Password, NickName: registerForm.Nickname})
	if err != nil {
		zap.S().Errorf("[UserRegister]   [fail to register new user]  ERROR: %s", err.Error())
		grpc_client.ParseGrpcErrorToHttp(err, c)
		return
	}
	result.Password = ""
	c.JSON(http.StatusOK, result)
	return
}
