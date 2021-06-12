package router

import (
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common/middlewares"
	"github.com/needon1997/theshop-api/internal/order_api/api/cart"
	"github.com/needon1997/theshop-api/internal/order_api/api/order"
	"github.com/needon1997/theshop-api/internal/order_api/api/payment"
)

func InitOrderRouter(router *gin.RouterGroup) {
	cartRouter := router.Group("/cart").Use(middlewares.JWTAUTH())
	{
		cartRouter.GET("/", cart.List)
		cartRouter.DELETE("/:id", cart.Delete)
		cartRouter.POST("/", cart.New)
		cartRouter.PATCH("/", cart.Edit)
	}
	orderRouter := router.Group("/order").Use(middlewares.JWTAUTH())
	{
		orderRouter.GET("/", order.List)
		orderRouter.POST("", middlewares.JwtTokenPassThrough(), order.New)
		orderRouter.GET("/:id", order.Detail)
	}
	paymentRouter := router.Group("/payment").Use(middlewares.JWTAUTH())
	{
		paymentRouter.GET("/execute/:id", payment.Execute)
	}
}
