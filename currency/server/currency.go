package server

import (
	"context"
	"log"

	"github.com/GrosfeldEzekiel/coffee-shop/common/protos"
	"github.com/GrosfeldEzekiel/coffee-shop/currency/data"
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
	c.l.Println("Handling get rate, base: ", rr.GetBase())
	c.l.Println("Handling get rate, destination:  ", rr.GetDestination())

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}
