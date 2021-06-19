package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common/middlewares"
	"github.com/needon1997/theshop-api/internal/userop_api/router"
	"net/http"
)

func InitializeRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(middlewares.Cors(), middlewares.Trace())
	engine.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"health": true,
		})
	})
	apiRouter := engine.Group("/up/v1")
	router.InitUserRouter(apiRouter)
	return engine
}
