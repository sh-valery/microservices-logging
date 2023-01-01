package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sh-valery/microservices-logging/gateway/internal/handler"
	"log"
)

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	v1.POST("/fx", handler.HandleFXRequest)

	err := r.Run() // listen and serve on default 0.0.0.0:8080
	if err != nil {
		log.Fatal(err)
	}
}
