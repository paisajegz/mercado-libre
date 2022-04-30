package api

import (
	"mercado-level-1/pkg/config"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var router *echo.Echo
var validate *validator.Validate

func setConfigurations(configPath string) {
	config.Setup(configPath)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	validate = validator.New()
	CustomValidation(validate)
	loc, _ := time.LoadLocation("America/Bogota")
	time.Local = loc
	router = echo.New()
	Setup(router, validate)
	setConfigurations(configPath)
	conf := config.GetConfig()
	router.Use(middleware.Recover())
	// router.Use(middleware.BodyLimit("100M"))
	router.Use(middleware.CORS())
	port := os.Getenv("PORT")
	if port != "" {
		conf.Server.Port = port
	}
	router.Start(":" + conf.Server.Port)
}
