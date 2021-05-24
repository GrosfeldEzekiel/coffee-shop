package handlers

import (
	"net/http"

	coffee_helper "github.com/GrosfeldEzekiel/coffee-shop/common/helpers"
	"github.com/GrosfeldEzekiel/coffee-shop/products-api/data"
)

func (p *Products) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Info("Handling POST")

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	prod := data.AddProduct(product)

	coffee_helper.ToJSON(prod, rw)
}
