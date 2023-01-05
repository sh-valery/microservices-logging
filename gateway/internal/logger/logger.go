package logger

import (
	"context"
	"go.uber.org/zap"
)

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

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

// WithContext tries to insert a track header from the context to the log message if context exists.
func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return Logger
	}

	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		requestID = "unknown"
	}

	return Logger.With(zap.String("requestID", requestID))
}
