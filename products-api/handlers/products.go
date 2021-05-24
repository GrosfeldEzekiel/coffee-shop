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
	"net/http"

	"github.com/GrosfeldEzekiel/coffee-shop/products-api/data"
	"github.com/hashicorp/go-hclog"
)

type Products struct {
	l          hclog.Logger
	productsDB *data.ProductsDB
}

// A list of products in the response
// swagger:response productsResponse
type productsResponse struct {
	// All of products
	// in: body
	Body []data.Product
}

// swagger:response editedProduct
type editedProduct struct {
	// Edited Product
	// in: body
	Body data.Product
}

// swagger:parameters editProduct
type productIdParameter struct {
	// The ID of the product
	// in: path
	// required: true
	ID int `json:"id"`
}

func NewProducts(l hclog.Logger, productsDb *data.ProductsDB) *Products {
	return &Products{l, productsDb}
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := &data.Product{}

		err := product.FromJSON(r.Body)

		err = product.Validate()

		if err != nil {
			http.Error(rw, fmt.Sprintf("Bad input: %s", err), http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(rw, "Unable serialize JSON", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)

		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)

	})
}
