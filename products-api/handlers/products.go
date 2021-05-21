package handlers

import (
	"log"
	"net/http"

	"github.com/GrosfeldEzekiel/coffee-shop/products-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.createProduct(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()

	err := products.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marsh json", http.StatusInternalServerError)
	}
}

func (p *Products) createProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handling POST")

	product := &data.Product{}

	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to parse JSON", http.StatusBadRequest)
	}

	data.AddProduct(product)
}
