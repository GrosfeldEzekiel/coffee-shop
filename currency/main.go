package main

import (
	"log"
	"net"
	"os"

	"github.com/GrosfeldEzekiel/coffee-shop/common/protos"
	"github.com/GrosfeldEzekiel/coffee-shop/currency/data"
	"github.com/GrosfeldEzekiel/coffee-shop/currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	l := log.New(os.Stdout, "user-service", log.LstdFlags)
	log := hclog.Default()
	gs := grpc.NewServer()

	rates, err := data.NewRates(log)
	if err != nil {
		l.Printf("Error getting rates")
		os.Exit(1)
	}

	cs := server.NewCurrency(rates, l)

	protos.RegisterCurrencyServer(gs, cs)

	listener, err := net.Listen("tcp", ":9092")

	if err != nil {
		l.Printf("Error serving")
		os.Exit(1)
	}

	gs.Serve(listener)
}
