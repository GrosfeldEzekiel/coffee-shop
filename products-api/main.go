package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GrosfeldEzekiel/coffee-shop/common/protos"
	"github.com/GrosfeldEzekiel/coffee-shop/products-api/data"
	"github.com/GrosfeldEzekiel/coffee-shop/products-api/handlers"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	// gRPC client
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	cc := protos.NewCurrencyClient(conn)
	// Handlers

	pdb := data.NewProductDB(cc, log)

	ph := handlers.NewProducts(log, pdb)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products", ph.GetProducts)

	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetProduct).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.GetProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.CreateProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	origins := []string{"http://localhost:3000", "http://172.24.101.211:3000"}
	ch := gohandlers.CORS(gohandlers.AllowedOrigins(origins), gohandlers.AllowedHeaders([]string{"application/json"}))

	s := http.Server{
		Addr:         ":8080",
		Handler:      ch(sm),
		ErrorLog:     log.StandardLogger(&hclog.StandardLoggerOptions{}),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Info("Listening on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			log.Error(err.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	log.Info("Gracefull shutdown: ", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)

	defer cancel()
}
