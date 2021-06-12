package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/needon1997/theshop-api/internal/common"
	"go.uber.org/zap"
	"net/http"
)

const AUTH = 2

func ADMIN_AUTH() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "unknown error",
			})
			c.Abort()
			return
		}
		userInfoClaim, ok := claims.(*common.JWTUserInfoClaim)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "unknown error",
			})
			c.Abort()
			return
		}
		if userInfoClaim.Role != AUTH {
			zap.S().Infow("attempt access", "URL", c.Request.URL, "userID", userInfoClaim.Id, "result", "rejected")
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "unauthorized acess",
			})
			c.Abort()
			return
		}
		zap.S().Infow("attempt access", "URL", c.Request.URL, "userID", userInfoClaim.Id, "result", "pass")
		c.Next()
		return
	}
}
