package service

import (
	"context"
	m "github.com/sh-valery/microservices-logging/gateway/internal/model"
	r "github.com/sh-valery/microservices-logging/gateway/internal/rpc_gen"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"time"
)

//go:generate mockgen -source=fx.go -destination=../mock/fx.go -package=mocks
type FXRate interface {
	GetFxRate(ctx context.Context, in *r.FxServiceRequest, opts ...grpc.CallOption) (*r.FxServiceResponse, error)
}

type UUIDGenerator interface {
	GenerateUUID() string
}

type Now interface {
	Now() time.Time
}

type FXService struct {
	FX   FXRate
	UUID UUIDGenerator
	Date Now
}

func (f *FXService) GetQuote(ctx context.Context, request *m.FXRequest) (m.FXResponse, error) {
	requestID := ctx.Value("requestID").(string)
	rpcRequest := &r.FxServiceRequest{
		SourceCurrencyCode: request.SourceCurrency,
		TargetCurrencyCode: request.TargetCurrency,

		Base: &r.Base{
			RequestID: &requestID,
		},
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
		DateTime:       f.Date.Now(),
	}, nil
}
