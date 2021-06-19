package router

import (
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common/middlewares"
	"github.com/needon1997/theshop-api/internal/userop_api/api/address"
	"github.com/needon1997/theshop-api/internal/userop_api/api/message"
	"github.com/needon1997/theshop-api/internal/userop_api/api/user_fav"
)

func InitUserRouter(router *gin.RouterGroup) {
	addressRouter := router.Group("/address").Use(middlewares.JWTAUTH())
	{
		addressRouter.GET("", address.List)
		addressRouter.DELETE("/:id", address.Delete)
		addressRouter.POST("", address.New)
		addressRouter.PUT("/:id", address.Edit)
	}
	messageRouter := router.Group("/message").Use(middlewares.JWTAUTH())
	{
		messageRouter.GET("", message.List)
		messageRouter.POST("", message.New)
	}
	UserFavRouter := router.Group("/userfavs").Use(middlewares.JWTAUTH())
	{
		UserFavRouter.DELETE("/:id", user_fav.Delete)
		UserFavRouter.GET("/:id", user_fav.Detail)
		UserFavRouter.POST("", user_fav.New)
		UserFavRouter.GET("", user_fav.List)
	}
}
