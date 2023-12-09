package routes

import (
	"github.com/andersonribeir0/clever/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Static("/static", "./static")
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	api := handlers.NewAPI()

	e.GET("/", api.CreateHandler(handlers.HandleHome))

	return e
}
