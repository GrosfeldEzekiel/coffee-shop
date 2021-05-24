package handlers

import (
	"net/http"
	"strconv"

	"github.com/GrosfeldEzekiel/coffee-shop/products-api/data"
	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} products editProduct
// Edits the product
// responses:
//	200: editedProduct

// Update product
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Info("Handling PUT")

	vars := mux.Vars(r)

	idString := vars["id"]

	id, _ := strconv.Atoi(idString)

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	err := data.UpdateProduct(id, product)

	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}

	err = product.ToJSONSingle(rw)

	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
	}
}
