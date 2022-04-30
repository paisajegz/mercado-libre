package xmen

import (
	"mercado-level-1/pkg/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ValidateOrders(v *validator.Validate) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			dna := new(DNA)
			_ = c.Bind(&dna)
			if err := v.Struct(dna); err != nil {
				errs := err.(validator.ValidationErrors)
				return c.JSON(http.StatusForbidden, utils.GenerateMessage(v, errs))
			}
			c.Set("dna", dna)
			return next(c)
		}
	}
}
