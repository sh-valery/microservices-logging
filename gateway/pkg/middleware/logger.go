package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
)

const DefaultTrackHeader = "X-Request-ID"

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
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
			logger.Error("Can't parse request body", zap.Error(err))
		}

		logger.Info("Request",
			zap.String(DefaultTrackHeader, c.GetHeader(DefaultTrackHeader)),
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

func InitLogger() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	Logger = logger
	Sugar = Logger.Sugar()
	Logger.Info("Initialized logger done")
	Logger.Info("Test info output 1/3")
	Logger.Error("Test error output 2/3")
	Logger.Warn("Test warn output 3/3")
}
