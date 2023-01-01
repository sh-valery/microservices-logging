package main

import (
	"context"
	"flag"
	"fmt"
	r "github.com/sh-valery/microservices-logging/fx/internal/rpc_gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

type FX struct {
	// you can inject any service here, dal to connect to db or cache or any others
	// you can mock dependencies in tests
	r.UnimplementedFxServiceServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func (f *FX) GetFxRate(ctx context.Context, request *r.FxServiceRequest) (*r.FxServiceResponse, error) {
	return &r.FxServiceResponse{
		SourceCurrencyCode: request.GetSourceCurrencyCode(),
		TargetCurrencyCode: request.GetTargetCurrencyCode(),
		Rate:               1.0,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	r.RegisterFxServiceServer(s, &FX{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
