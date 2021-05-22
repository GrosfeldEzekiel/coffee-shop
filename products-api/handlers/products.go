// Package classification Product API.
//
// documentation for product API
//
//     Schemes: http, https
//     BasePath: /
//     Version: 1.0.0
//
//	   Consumes:
//	   - application/json
//
//	   Produces:
//	   - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/GrosfeldEzekiel/coffee-shop/products-api/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling POST")

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(product)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling PUT")

	vars := mux.Vars(r)

	idString := vars["id"]

	id, _ := strconv.Atoi(idString)

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	uError := data.UpdateProduct(id, product)

	if uError == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if uError != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := &data.Product{}

		err := product.FromJSON(r.Body)

		err = product.Validate()

		if err != nil {
			http.Error(rw, fmt.Sprintf("Bad input: %s", err), http.StatusBadRequest)
		}

		if err != nil {
			http.Error(rw, "Unable serialize JSON", http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)

		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)

	})
}
