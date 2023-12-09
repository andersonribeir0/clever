package handlers

import (
	"github.com/labstack/echo/v4"
)

type API struct{}

func NewAPI() *API {
	return &API{}
}

func (a *API) CreateHandler(handlerFunc func(*API, echo.Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handlerFunc(a, c)
	}
}
