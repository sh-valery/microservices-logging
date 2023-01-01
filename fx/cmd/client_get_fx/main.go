package main

import (
	"context"
	"flag"
	r "github.com/sh-valery/microservices-logging/fx/internal/rpc_gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var host = flag.String("host", "localhost:50051", "server address for connection, default: localhost:50051")

func main() {
	flag.Parse()

	// init rpc connection
	conn, err := grpc.Dial(*host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connection error: %v", err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalf("connection close error: %v", err)
		}
	}()

	client := r.NewFxServiceClient(conn)

	// build request
	request := &r.FxServiceRequest{
		SourceCurrencyCode: "USD",
		TargetCurrencyCode: "CHF",
	}
	log.Printf("request from the client: %+v", request)

	// send request
	response, err := client.GetFxRate(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response from the rpc server: %+v", response)
}
