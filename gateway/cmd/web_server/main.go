package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sh-valery/microservices-logging/gateway/internal/handler"
	"github.com/sh-valery/microservices-logging/gateway/internal/rpc_gen"
	"github.com/sh-valery/microservices-logging/gateway/internal/service"
	"github.com/sh-valery/microservices-logging/gateway/internal/util"
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

	// init service layer
	client := rpc_gen.NewFxServiceClient(conn)
	serviceLayer := &service.FXService{
		FX:   client,
		UUID: util.NewUUIDGenerator(),
		Date: util.NewRealClock(),
	}

	// init handler, inject service layer into handler
	fxHandler := handler.NewFxHandler(serviceLayer)

	// init router
	r := gin.Default()
	v1 := r.Group("api/v1")
	v1.POST("/fx", func(c *gin.Context) {
		fxHandler.HandleFXRequest(c)
	})

	// run server
	err = r.Run() // listen and serve on default 0.0.0.0:8080
	if err != nil {
		log.Fatal(err)
	}
}
