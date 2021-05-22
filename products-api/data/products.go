package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required,desc"`
	Price       float32 `json:"price" validate:"required,gt=0"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
}

type Products []*Product

var ErrorProductNotFound = fmt.Errorf("Product not found")

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation(`desc`, validateDescription)
	return validate.Struct(p)
}

func validateDescription(fl validator.FieldLevel) bool {
	//Should have at least a dot
	regex := regexp.MustCompile(`[.]`)
	validate := regex.FindAllString(fl.Field().String(), -1)

	if len(validate) > 0 {
		return true
	}

	return false
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	p.CreatedAt = time.Now().UTC().String()
	p.UpdatedAt = time.Now().UTC().String()
	productList = append(productList, p)
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func GetProducts() Products {
	return productList
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProductById(id)

	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

func findProductById(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrorProductNotFound
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Delicious Coffee",
		Price:       10,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Delicious Coffee, but strong",
		Price:       12,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
