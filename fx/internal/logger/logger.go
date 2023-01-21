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
	logger.With()
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

func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return Logger
	}

	requestID, ok := ctx.Value("requestid").(string) // ctx.set make the key lowercase
	if !ok {
		requestID = "unknown"
	}

	return Logger.With(zap.String("requestid", requestID))
}
