package handlers

import (
	"github.com/a-h/templ"
	"github.com/andersonribeir0/clever/internal/components"
	"github.com/andersonribeir0/clever/internal/components/sidebar"
	"github.com/andersonribeir0/clever/internal/components/sidebaritem"
	"github.com/labstack/echo/v4"
)

func HandleHome(a *API, c echo.Context) error {
	profile := sidebaritem.Profile{
		Name:        "Anderson",
		Description: "Software Engineer",
	}
	profileItemCmp := sidebaritem.ProfileItem(profile)
	sidebar := sidebar.Sidebar(profileItemCmp)
	t := templ.Handler(components.Home(sidebar))

	return t.Component.Render(c.Request().Context(), c.Response().Writer)
}
