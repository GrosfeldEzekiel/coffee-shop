package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/GrosfeldEzekiel/coffee-shop/common/protos"
	"github.com/go-playground/validator"
	"github.com/hashicorp/go-hclog"
)

// swagger:model
type Product struct {
	// the id of the product
	//
	// required: false
	ID int `json:"id"`

	// the name of the product
	//
	// required: true
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required,desc"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
}

type Products []*Product

type ProductsDB struct {
	currency protos.CurrencyClient
	log      hclog.Logger
}

func NewProductDB(c protos.CurrencyClient, l hclog.Logger) *ProductsDB {
	return &ProductsDB{c, l}
}

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

	return len(validate) > 0
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Product) ToJSONSingle(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
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

func (p *ProductsDB) GetProducts(currency string) (Products, error) {
	if currency == "" {
		return productList, nil
	}

	pr := Products{}

	rate, _ := p.getRate(currency)

	for _, p := range productList {
		np := *p
		np.Price = np.Price * rate
		pr = append(pr, &np)
	}

	return pr, nil
}

func (p *ProductsDB) GetProduct(id int, currency string) (Product, error) {
	pr, _, _ := findProductById(id)

	product := *pr

	if currency == "" {
		return product, nil
	}

	rate, _ := p.getRate(currency)

	product.Price = product.Price * rate
	return product, nil
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
	{
		ID:          1,
		Name:        "Latte",
		Description: "Delicious Coffee",
		Price:       10,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Delicious Coffee, but strong",
		Price:       12,
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}

func (p *ProductsDB) getRate(destination string) (float64, error) {
	rr := &protos.RateRequest{
		Base:        protos.Currencies_EUR,
		Destination: protos.Currencies(protos.Currencies_value[destination]),
	}

	resp, err := p.currency.GetRate(context.Background(), rr)
	return resp.Rate, err
}
