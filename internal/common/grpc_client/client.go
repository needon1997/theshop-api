package grpc_client

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common/config"
	_ "github.com/needon1997/theshop-api/internal/common/grpc_consul_resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

const INTERNAL_ERROR = "server internal error"
const UNAVAILABLE = "server unavailable"
const UNKNOWN_ERROR = "unknown error"
const ERROR = "error"

func ParseGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			switch s.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					ERROR: s.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					ERROR: INTERNAL_ERROR,
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					ERROR: s.Message(),
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					ERROR: UNAVAILABLE,
				})
			case codes.AlreadyExists:
				c.JSON(http.StatusInternalServerError, gin.H{
					ERROR: s.Message(),
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					ERROR: UNKNOWN_ERROR,
				})
			}
			return
		}
	}
}

const CONSUL_LB_TEMPLATE = "consul://%s/%s"

func GetUserSvcConn() (*grpc.ClientConn, error) {
	zap.S().Debug("Get connect gRPC user service server")
	url := fmt.Sprintf(CONSUL_LB_TEMPLATE, config.ServerConfig.ServiceConfig.UserServiceName, "")
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		zap.S().Errorf("[GetUserSvcClient]  [fail to connect with service provider]   ERROR: %s", err.Error())
		return nil, errors.New(INTERNAL_ERROR)
	}
	return conn, nil
}
func GetEmailSvcConn() (*grpc.ClientConn, error) {
	zap.S().Debug("Get connect gRPC email service server")
	url := fmt.Sprintf(CONSUL_LB_TEMPLATE, config.ServerConfig.ServiceConfig.EmailServiceName, "")
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		zap.S().Errorf("[GetEmailSvcClient]  [fail to connect with service provider]   ERROR: %s", err.Error())
		return nil, errors.New(INTERNAL_ERROR)
	}
	return conn, nil
}
func GetGoodsSvcConn() (*grpc.ClientConn, error) {
	zap.S().Debug("Get connect gRPC goods service server")
	url := fmt.Sprintf(CONSUL_LB_TEMPLATE, config.ServerConfig.ServiceConfig.GoodsServiceName, "")
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		zap.S().Errorf("[GetUserSvcClient]  [fail to connect with service provider]   ERROR: %s", err.Error())
		return nil, errors.New(INTERNAL_ERROR)
	}
	return conn, nil
}
