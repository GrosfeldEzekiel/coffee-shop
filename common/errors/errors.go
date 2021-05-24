package coffee_errors

import (
	"net/http"

	coffee_helper "github.com/GrosfeldEzekiel/coffee-shop/common/helpers"
)

type Error struct {
	Mesagge string `json:"message"`
}

type Errors []Error

func NewError(message string) Error {
	return Error{Mesagge: message}
}

func HandleError(e Errors, rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)

	coffee_helper.ToJSON(e, rw)
}
