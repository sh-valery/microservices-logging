package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func headerToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		ctx = context.WithValue(ctx, k, v[0])
	}

	if md.Get("requestID") == nil {
		ctx = context.WithValue(ctx, "requestID", "unknown")
	}

	return ctx
}

func HeaderMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ctx = headerToContext(ctx, md)
	}
	return handler(ctx, req)
}
