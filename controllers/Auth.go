package controllers

import (
	"net/http"
	"poliklinik/models"

	"github.com/flosch/pongo2/v4"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	data := pongo2.Context{
		"title": "Login",
	}
	return c.Render(http.StatusOK, "views/auth/login.html", data)
}
func Register(c echo.Context) error {
	data := pongo2.Context{
		"title": "Register",
	}
	return c.Render(http.StatusOK, "views/auth/register.html", data)
}
func RegisterPost(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	resp := models.Register(username, password)

	if resp.Status != 200 {
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, resp)
}
