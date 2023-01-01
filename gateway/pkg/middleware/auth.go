package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, password, hasAuth := c.Request.BasicAuth()

		// auth flow
		if hasAuth && user == "user" && password == "pass" {
			c.Next()
			return
		}

		// no auth flow
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
