package router

import (
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common/middlewares"
	"github.com/needon1997/theshop-api/internal/user_api/api"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/list", middlewares.JWTAUTH(), middlewares.ADMIN_AUTH(), api.GetUserList)
		userRouter.POST("/pw_login", api.UserPasswordLogin)
		userRouter.POST("/register", api.Register)
	}
}
