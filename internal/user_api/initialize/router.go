package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common/middlewares"
	"github.com/needon1997/theshop-api/internal/user_api/router"
	"net/http"
)

func InitializeRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(middlewares.Cors())
	engine.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"health": true,
		})
	})
	apiRouter := engine.Group("/u/v1")
	router.InitUserRouter(apiRouter)
	router.InitBaseRouter(apiRouter)
	return engine
}
