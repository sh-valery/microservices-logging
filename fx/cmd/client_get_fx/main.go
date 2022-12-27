package client_get_fx

import (
	"context"
	r "github.com/sh-valery/microservices-logging/internal/rpc_gen"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// todo: 1. complete client connection
	client := r.NewFxServiceClient(grpc.ClientConnInterface())

	})
	request := &r.FxServiceRequest{
		SourceCurrencyCode: "USD",
		TargetCurrencyCode: "CHF",
	}
	response, err := client.GetFxRate(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response from the rpc server: %v", response)
}
