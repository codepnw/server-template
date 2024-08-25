package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		accessToken := strings.TrimPrefix(token, "Bearer ")

		if accessToken == "" || accessToken == token {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("accessToken", accessToken)
		c.Next()
	}
}
