package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common"
	"net/http"
)

func JWTAUTH() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			token = c.Query("x-token")
			if token == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": "please login",
				})
				c.Abort()
				return
			}
		}
		claim, err := common.ValidateTokenAndRetrieveInfo(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("claims", claim)
		c.Set("userID", claim.Id)
		c.Next()
	}
}

func JwtTokenPassThrough() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			token = c.Query("x-token")
			if token == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": "please login",
				})
				c.Abort()
				return
			}
		}
		c.Set("x-token", token)
	}
}
