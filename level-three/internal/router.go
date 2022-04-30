package api

import (
	"mercado-level-1/internal/xmen"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Setup(router *echo.Echo, v *validator.Validate) {
	xmen.SetXmenRoutes(router, v)
}
