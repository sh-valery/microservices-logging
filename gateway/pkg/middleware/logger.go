package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		c.Next()
		cost := time.Since(start)

		// log body for POST, PUT, PATCH. Our gateway expects only json format
		var reqJSONBody []byte
		var err error
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			reqJSONBody, err = json.Marshal(c.Request)
			if err != nil {
				logger.Error("Can't parse request body", zap.Error(err))
				reqJSONBody = []byte("Unable to parse request body")
			}
		}

		// log basic info
		logger.Info(path,
			zap.String(DefaultTrackHeader, c.GetHeader(DefaultTrackHeader)),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("requestBody", string(reqJSONBody)),
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
