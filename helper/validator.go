package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/triyasmkom/rest-api-echo/models/response"
	"strings"
)

var Validator = validator.New()
func Params(s interface{}) response.Body {
	var  errMessage []string

	if err := Validator.Struct(s); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errMessage = append(errMessage, err.Field() + " tidak valid")
		}
		return response.Body{
			Status: false,
			Error: strings.Join(errMessage, ", "),
		}
	}

	return response.Body{
		Status: true,
		Data: s,
	}
}
