package utils

import (
	"mercado-level-1/pkg/translation"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Key     string
	Message string
}

func GenerateMessage(v *validator.Validate, errs validator.ValidationErrors) map[string]interface{} {

	trans := *translation.TranslationValidation(v)

	var errors []Error

	for _, e := range errs {
		error := Error{
			Key:     strings.Split(e.Translate(trans), " ")[0],
			Message: e.Translate((trans)),
		}
		errors = append(errors, error)
	}

	ctx := make(map[string]interface{})
	ctx["errors"] = errors
	return ctx
}
