package handlers

import (
	"net/http"
	"strconv"

	coffee_errors "github.com/GrosfeldEzekiel/coffee-shop/common/errors"
	"github.com/gorilla/mux"
)

// swagger:route GET /products products listProducts
// Returns a list of all the products
// responses:
//	200: productsResponse

// Get all the products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	cur := r.URL.Query().Get("currency")
	products, err := p.productsDB.GetProducts(cur)
	if err != nil {
		e := coffee_errors.Errors{coffee_errors.Error{Mesagge: "There was a problem with the currency"}}
		coffee_errors.HandleError(e, rw, http.StatusBadRequest)
	}

	err = products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marsh json", http.StatusInternalServerError)
	}
}

func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idString := vars["id"]

	id, _ := strconv.Atoi(idString)

	cur := r.URL.Query().Get("currency")

	product, err := p.productsDB.GetProduct(id, cur)
	if err != nil {
		e := coffee_errors.Errors{coffee_errors.Error{Mesagge: err.Error()}}
		coffee_errors.HandleError(e, rw, http.StatusBadRequest)
		return
	}

	err = product.ToJSONSingle(rw)

	if err != nil {
		http.Error(rw, "Unable to marsh json", http.StatusInternalServerError)
	}
}
