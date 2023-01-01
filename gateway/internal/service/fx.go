package service

import (
	"context"
	m "github.com/sh-valery/microservices-logging/gateway/internal/model"
	r "github.com/sh-valery/microservices-logging/gateway/internal/rpc_gen"
	"github.com/shopspring/decimal"
)

//go:generate mockgen -source=fx.go -destination=../mock/fx.go -package=mocks
type FXRate interface {
	GetFxRate(ctx context.Context, request *r.FxServiceRequest) (*r.FxServiceResponse, error)
}

type UUIDGenerator interface {
	GenerateUUID() string
}

type Now interface {
	Now() string
}

type FXService struct {
	FX   FXRate
	UUID UUIDGenerator
	Date Now
}

func (f *FXService) GetQuote(ctx context.Context, request *m.FXRequest) (m.FXResponse, error) {
	rpcRequest := &r.FxServiceRequest{
		SourceCurrencyCode: request.SourceCurrency,
		TargetCurrencyCode: request.TargetCurrency,
	}
	rpcResponse, err := f.FX.GetFxRate(ctx, rpcRequest)
	if err != nil {
		return m.FXResponse{}, err
	}

	distAmount := decimal.NewFromFloat(request.SourceAmount).Mul(decimal.NewFromFloat(rpcResponse.Rate))

	return m.FXResponse{
		QuotationID:    f.UUID.GenerateUUID(),
		SourceCurrency: rpcResponse.SourceCurrencyCode,
		TargetCurrency: rpcResponse.TargetCurrencyCode,
		SourceAmount:   request.SourceAmount,
		DistAmount:     distAmount.InexactFloat64(),
		Date:           f.Date.Now(),
	}, nil
}
