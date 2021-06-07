package router

import (
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/order_api/api/cart"
	"github.com/needon1997/theshop-api/internal/order_api/api/order"
)

func InitOrderRouter(router *gin.RouterGroup) {
	cartRouter := router.Group("/cart")
	{
		cartRouter.GET("/", cart.List)
		cartRouter.DELETE("/:id", cart.Delete)
		cartRouter.POST("", cart.New)
		cartRouter.PATCH("/:id", cart.Edit)
	}
	orderRouter := router.Group("/order")
	{
		orderRouter.GET("/", order.List)
		orderRouter.POST("", order.New)
		orderRouter.DELETE("/:id", order.Delete)
		orderRouter.GET("/:id", order.Detail)
	}
}
