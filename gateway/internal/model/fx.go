package model

import (
	"time"
)

type FXRequest struct {
	SourceCurrency string
	TargetCurrency string
	SourceAmount   float64
}

type FXResponse struct {
	QuotationID    string
	SourceCurrency string
	TargetCurrency string
	SourceAmount   float64
	DistAmount     float64
	DateTime       time.Time
}