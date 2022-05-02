package api

import (
	"github.com/go-playground/validator/v10"
)

func CustomValidation(v *validator.Validate) {

	v.RegisterValidation("isBool", func(fl validator.FieldLevel) bool {
		return fl.Field().String() == "<bool Value>"
	})
	v.RegisterValidation("is-dna", func(fl validator.FieldLevel) bool {
		// var dna []string
		data := fl.Field().Interface().([]string)
		return ValidateDna(data)
	})

}

func ValidateDna(data []string) bool {
	countDna := len(data)
	for _, cadena := range data {
		if len(string(cadena)) != countDna {
			return false
		}
	}
	return true
}
