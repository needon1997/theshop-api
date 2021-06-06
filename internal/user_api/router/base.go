package router

import (
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/user_api/api"
)

func InitBaseRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/base")
	{
		userRouter.POST("/email", api.SendEmail)
	}
}
