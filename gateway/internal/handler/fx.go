package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	m "github.com/sh-valery/microservices-logging/gateway/internal/model"
	"net/http"
)

type fxHandler struct {
	quotationService Quotation
}

func NewFxHandler(s Quotation) *fxHandler {
	return &fxHandler{
		quotationService: s,
	}
}

type Quotation interface {
	GetQuote(ctx context.Context, request *m.FXRequest) (m.FXResponse, error)
}

func (f *fxHandler) HandleFXRequest(c *gin.Context) {
	// Parse request body
	serviceRequest := &m.FXRequest{}
	err := c.ShouldBindJSON(serviceRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Call service layer in or case just rpc
	result, err := f.quotationService.GetQuote(c, serviceRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Write response
	c.JSON(http.StatusOK, result)
}
