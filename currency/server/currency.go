package server

import (
	"context"
	protos "gRPC/currency/protos/currency"
	log2 "log"
)

type Currency struct {
	log *log2.Logger
}

func NewCurrency(log *log2.Logger) *Currency {
	return &Currency{log: log}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Println("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
