package controllers

import (
	"net/http"

	"github.com/flosch/pongo2/v4"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	data := pongo2.Context{
		"title":   "Home",
		"message": "Welcome to Echo Pongo2",
	}
	return c.Render(http.StatusOK, "views/home/index.html", data)
}
