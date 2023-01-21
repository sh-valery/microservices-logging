package main

import (
	"context"
	"flag"
	"fmt"
	l "github.com/sh-valery/microservices-logging/fx/internal/logger"
	"github.com/sh-valery/microservices-logging/fx/internal/middleware"
	r "github.com/sh-valery/microservices-logging/fx/internal/rpc_gen"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type FX struct {
	// you can inject any service here, dal to connect to db, cache and any others
	// you can mock dependencies in tests
	r.UnimplementedFxServiceServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func (f *FX) GetFxRate(ctx context.Context, request *r.FxServiceRequest) (*r.FxServiceResponse, error) {
	l.WithContext(ctx).Sugar().Infof("GetFxRate request: %+v", request)

	// way 1: use grpc metadata to get requestID form context, middleware passes it from header to context
	l.WithContext(ctx).Info("Header(context) RequestID: ")

	// way 2: use Base request struct in all requests/response,  it works for all request types
	l.Sugar.Info("Base struct RequestID: ",
		zap.String("requestID", request.GetBase().GetRequestID()))

	// error flow
	rate, err := getFxRate(request.GetSourceCurrencyCode(), request.GetTargetCurrencyCode())
	if err != nil {
		l.Sugar.Errorw("GetFxRate error",
			"requestID", request.Base.GetRequestID(),
			"error", err,
			"sourceCurrencyCode", request.GetSourceCurrencyCode(),
			"targetCurrencyCode", request.GetTargetCurrencyCode())
		return nil, err
	}

	// succeed flow
	return &r.FxServiceResponse{
		SourceCurrencyCode: request.GetSourceCurrencyCode(),
		TargetCurrencyCode: request.GetTargetCurrencyCode(),
		Rate:               rate,
	}, nil
}

// it should be an external service
func getFxRate(sourceCurrencyCode, targetCurrencyCode string) (float64, error) {
	// currency code validation
	if sourceCurrencyCode == "" || targetCurrencyCode == "" {
		return 0, fmt.Errorf("empty currency code")
	}

	if sourceCurrencyCode != "USD" {
		return 0, fmt.Errorf("source currency code is not supported")
	}

	rates := map[string]float64{
		"USD": 1,
		"EUR": 0.85,
		"CHF": 0.9,
		"GBP": 0.75,
		"JPY": 110,
		"CNY": 6.5,
	}

	rate, ok := rates[targetCurrencyCode]
	if !ok {
		return 0, fmt.Errorf("target currency code is not supported")
	}

	return rate, nil
}

func main() {
	flag.Parse()
	l.InitLogger()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		l.Sugar.Fatalf("failed to listen: %v", err)
	}
	// run grpc server without middleware, if you need grpc streaming
	s := grpc.NewServer(grpc.UnaryInterceptor(middleware.HeaderMiddleware))

	r.RegisterFxServiceServer(s, &FX{})
	l.Sugar.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		l.Sugar.Fatalf("failed to serve: %v", err)
	}
}
