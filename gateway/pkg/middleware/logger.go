package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	l "github.com/sh-valery/microservices-logging/gateway/internal/logger"
	"go.uber.org/zap"
	"io"
	"time"
)

const DefaultTrackHeader = "requestID"

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		// body can be read once, so copy it for logging and create a buffer back
		body, err := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		c.Next()
		cost := time.Since(start)

		if err != nil {
			l.Logger.Error("Can't parse request body", zap.Error(err))
		}

		l.Logger.Info("Request",
			zap.String("requestID", c.GetHeader(DefaultTrackHeader)),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("requestBody", string(body)),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("userAgent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
