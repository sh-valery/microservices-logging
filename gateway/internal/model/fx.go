package model

import (
	"time"
)

type FXRequest struct {
	SourceCurrency string  `json:"sourceCurrency"`
	TargetCurrency string  `json:"targetCurrency"`
	SourceAmount   float64 `json:"SourceAmount"`
}

type FXResponse struct {
	QuotationID    string    `json:"QuotationID"`
	SourceCurrency string    `json:"SourceCurrency"`
	TargetCurrency string    `json:"TargetCurrency"`
	SourceAmount   float64   `json:"SourceAmount"`
	DistAmount     float64   `json:"DistAmount"`
	DateTime       time.Time `json:"DateTime"`
}
