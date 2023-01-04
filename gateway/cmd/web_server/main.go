package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sh-valery/microservices-logging/gateway/internal/handler"
	l "github.com/sh-valery/microservices-logging/gateway/internal/logger"
	"github.com/sh-valery/microservices-logging/gateway/internal/rpc_gen"
	"github.com/sh-valery/microservices-logging/gateway/internal/service"
	"github.com/sh-valery/microservices-logging/gateway/internal/util"
	"github.com/sh-valery/microservices-logging/gateway/pkg/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var host = flag.String("host", "localhost:50051", "server address for connection, default: localhost:50051")

func main() {
	flag.Parse()
	l.InitLogger()

	// init rpc connection
	l.Sugar.Infof("Init rpc connection to %s ", *host)
	conn, err := grpc.Dial(*host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.Sugar.Fatalf("rpc connection error: %v", err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			l.Sugar.Fatalf("connection close error: %v", err)
		}
	}()
	client := rpc_gen.NewFxServiceClient(conn)
	l.Sugar.Info("Init rpc client success")

	// init service layer
	l.Sugar.Info("Init service layer")
	serviceLayer := &service.FXService{
		FX:   client,
		UUID: util.NewUUIDGenerator(),
		Date: util.NewRealClock(),
	}

	// init handler, inject service layer into handler
	l.Sugar.Info("Init handler layer")
	fxHandler := handler.NewFxHandler(serviceLayer)

	// init middleware
	l.Sugar.Info("Init router and middleware")
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.TrackHeader())

	r.Use(middleware.LoggerMiddleware())

	// init router
	v1 := r.Group("api/v1")
	v1.POST("/fx", func(c *gin.Context) {
		fxHandler.HandleFXRequest(c)
	})

	// run server
	l.Sugar.Info("Run server")
	err = r.Run() // listen and serve on default 0.0.0.0:8080
	if err != nil {
		log.Fatal(err)
	}
}
