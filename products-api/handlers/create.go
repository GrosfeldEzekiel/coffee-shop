package handlers

import (
	"net/http"

	"github.com/GrosfeldEzekiel/coffee-shop/products-api/data"
)

func (p *Products) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling POST")

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(product)
}
