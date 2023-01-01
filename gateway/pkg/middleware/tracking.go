package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TrackHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set(DefaultTrackHeader, requestID)
		c.Request.Header.Set(DefaultTrackHeader, requestID)
		c.Header(DefaultTrackHeader, requestID)
		c.Next()
	}
}
