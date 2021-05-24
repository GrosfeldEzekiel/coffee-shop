package server

import (
	"context"
	"log"

	"github.com/GrosfeldEzekiel/coffee-shop/common/protos"
	"github.com/GrosfeldEzekiel/coffee-shop/currency/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Currency struct {
	rates *data.ExchangeRates
	l     *log.Logger
	protos.UnimplementedCurrencyServer
}

func NewCurrency(r *data.ExchangeRates, l *log.Logger) *Currency {
	return &Currency{r, l, protos.UnimplementedCurrencyServer{}}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	if rr.Base == rr.Destination {
		err := status.Newf(codes.InvalidArgument, "Base currency cannot be equal to destination")
		err, _ = err.WithDetails(rr)
		return nil, err.Err()
	}

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}
