package grpc_client

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/needon1997/theshop-api/internal/common/config"
	_ "github.com/needon1997/theshop-api/internal/common/grpc_consul_resolver"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

const INTERNAL_ERROR = "server internal error"
const UNAVAILABLE = "server unavailable"
const UNKNOWN_ERROR = "unknown error"
const ERROR = "error"

var opts = []grpc_retry.CallOption{
	grpc_retry.WithMax(3),
	grpc_retry.WithPerRetryTimeout(2 * time.Second),
	grpc_retry.WithCodes(codes.Unavailable, codes.DeadlineExceeded, codes.Unknown),
}

func ParseGrpcErrorToHttp(err error, c *gin.Context) {
	ispan, _ := c.Get("span")
	span := ispan.(opentracing.Span)
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			switch s.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					ERROR: s.Message(),
				})
			case codes.Internal:
				span.SetTag("error", true)
				c.JSON(http.StatusInternalServerError, gin.H{
					ERROR: INTERNAL_ERROR,
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					ERROR: s.Message(),
				})
			case codes.Unavailable:
				span.SetTag("error", true)
				c.JSON(http.StatusInternalServerError, gin.H{
					ERROR: UNAVAILABLE,
				})
			case codes.AlreadyExists:
				c.JSON(http.StatusInternalServerError, gin.H{
					ERROR: s.Message(),
				})
			default:
				span.SetTag("error", true)
				c.JSON(http.StatusInternalServerError, gin.H{
					ERROR: UNKNOWN_ERROR,
				})
			}
			return
		}
	}
}

const LB_TEMPLATE = "kubernetes://%s:%v/"

func GetUserSvcConn() (*grpc.ClientConn, error) {
	zap.S().Debug("Get connect gRPC user service server")
	url := fmt.Sprintf(LB_TEMPLATE, config.ServerConfig.ServiceConfig.UserServiceName, 80)
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)), grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	if err != nil {
		zap.S().Errorf("[GetUserSvcClient]  [fail to connect with service provider]   ERROR: %s", err.Error())
		return nil, err
	}
	return conn, nil
}
func GetEmailSvcConn() (*grpc.ClientConn, error) {
	zap.S().Debug("Get connect gRPC email service server")
	url := fmt.Sprintf(LB_TEMPLATE, config.ServerConfig.ServiceConfig.EmailServiceName, 80)
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)), grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	if err != nil {
		zap.S().Errorf("[GetEmailSvcClient]  [fail to connect with service provider]   ERROR: %s", err.Error())
		return nil, err
	}
	return conn, nil
}
func GetGoodsSvcConn() (*grpc.ClientConn, error) {
	zap.S().Debug("Get connect gRPC goods service server")
	url := fmt.Sprintf(LB_TEMPLATE, config.ServerConfig.ServiceConfig.GoodsServiceName, 80)
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)), grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	if err != nil {
		zap.S().Errorf("[GetGoodsSvcClient]  [fail to connect with service provider]   ERROR: %s", err.Error())
		return nil, err
	}
	return conn, nil
}
func GetInventorySvcConn() (*grpc.ClientConn, error) {
	zap.S().Debug("Get connect gRPC inventory service server")
	url := fmt.Sprintf(LB_TEMPLATE, config.ServerConfig.ServiceConfig.InventoryServiceName, 80)
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)), grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	if err != nil {
		zap.S().Errorf("[GetInventorySvcClient]  [fail to connect with service provider]   ERROR: %s", err.Error())
		return nil, err
	}
	return conn, nil
}
func GetOrderSvcConn() (*grpc.ClientConn, error) {
	zap.S().Debug("Get connect gRPC order service server")
	url := fmt.Sprintf(LB_TEMPLATE, config.ServerConfig.ServiceConfig.OrderServiceName, 80)
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)), grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	if err != nil {
		zap.S().Errorf("[GetOrderSvcClient]  [fail to connect with service provider]   ERROR: %s", err.Error())
		return nil, err
	}
	return conn, nil
}
func GetPaymentSvcConn() (*grpc.ClientConn, error) {
	zap.S().Debug("Get connect gRPC payment service server")
	url := fmt.Sprintf(LB_TEMPLATE, config.ServerConfig.ServiceConfig.PaymentServiceName, "")
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)), grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	if err != nil {
		zap.S().Errorf("[GetPaymentSvcClient]  [fail to connect with service provider]   ERROR: %s", err.Error())
		return nil, err
	}
	return conn, nil
}
func GetUserOpSvcConn() (*grpc.ClientConn, error) {
	zap.S().Debug("Get connect gRPC userop service server")
	url := fmt.Sprintf(LB_TEMPLATE, config.ServerConfig.ServiceConfig.UserOpServiceName, 80)
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)), grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
	if err != nil {
		zap.S().Errorf("[GetUserOpSvcConn]  [fail to connect with service provider]   ERROR: %s", err.Error())
		return nil, err
	}
	return conn, nil
}
