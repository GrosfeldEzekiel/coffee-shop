package handlers

import (
	"net/http"

	"github.com/GrosfeldEzekiel/coffee-shop/products-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of all the products
// responses:
//	200: productsResponse

// Get all the products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	products := data.GetProducts()

	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marsh json", http.StatusInternalServerError)
	}
}
