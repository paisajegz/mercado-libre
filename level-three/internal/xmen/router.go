package xmen

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func SetXmenRoutes(router *echo.Echo, v *validator.Validate) {
	router.POST("/mutant", VerifyMutant, ValidateOrders(v))
	router.GET("/stats", Stats)
}
