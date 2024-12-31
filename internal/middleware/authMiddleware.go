package middleware

import (
	"blog-server/pkg/authToken"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 authToken
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization authToken required"})
			c.Abort()
			return
		}
		// 解析token
		tokenClaims, err := authToken.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		// 续期
		if tokenClaims.ExpiresAt.After(time.Now()) {
			// 续期
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}
		c.Set("username", tokenClaims.Username)
		c.Next()
	}
}
