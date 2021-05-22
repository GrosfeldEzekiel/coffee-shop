package handlers

import (
	"net/http"

	"github.com/GrosfeldEzekiel/coffee-shop/products-api/data"
)

// A list of products in the response
// swagger:response productsResponse
type productsResponse struct {
	// All of products
	// in: body
	Body []data.Product
}

// swagger:route GET /products products listProducts
// Returns a list of
// responses:
//	200: productsResponse
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	products := data.GetProducts()

	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marsh json", http.StatusInternalServerError)
	}
}
