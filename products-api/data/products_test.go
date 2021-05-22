package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		ID:          1,
		Name:        "Name",
		Description: "ASDAdF",
		Price:       2.0,
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
