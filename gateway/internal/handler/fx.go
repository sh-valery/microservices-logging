package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	r "github.com/sh-valery/microservices-logging/gateway/internal/rpc_gen"
	"net/http"
)

// FX interface in the example ss equal the rpc call,
// but in real project rpc should be incapsulated in the service layer with some business logic
type FX interface {
	GetFxRate(ctx context.Context, request *r.FxServiceRequest) (*r.FxServiceResponse, error)
}

var FXService FX

func HandleFXRequest(c *gin.Context) {
	// Parse request body
	rpcRequest := &r.FxServiceRequest{}
	err := c.ShouldBindJSON(rpcRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Call service layer in or case just rpc
	result, err := FXService.GetFxRate(c, rpcRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Write response
	c.JSON(http.StatusOK, result)
}
