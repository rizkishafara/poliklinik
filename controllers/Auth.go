package controllers

import (
	"net/http"
	"poliklinik/models"
	"poliklinik/utils"

	"github.com/flosch/pongo2/v4"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	data := pongo2.Context{
		"title": "Login",
	}
	return c.Render(http.StatusOK, "views/auth/login.html", data)
}
func LoginPost(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	resp := models.Login(username, password)

	if resp.Status != 200 {
		return c.JSON(http.StatusBadRequest, resp)
	}

	dataMap, ok := resp.Data.(map[string]interface{})
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid response data"})
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,                 // 24 hours
		HttpOnly: false,                 // Allow JavaScript access
		SameSite: http.SameSiteNoneMode, // Enable cross-site access
		Secure:   true,                  // Required for SameSite=None
	}

	sess.Values["Auth"] = true
	sess.Values["id"] = dataMap["id"]
	sess.Values["username"] = dataMap["username"]
	sess.Save(c.Request(), c.Response())

	resp = utils.Respon{
		Status:  resp.Status,
		Message: resp.Message,
	}

	return c.JSON(http.StatusOK, resp)
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
	confirmpassword := c.FormValue("confirmpassword")

	if password != confirmpassword {
		respon := utils.Respon{
			Status:  http.StatusBadRequest,
			Message: "Password not match",
		}
		return c.JSON(http.StatusBadRequest, respon)
	}

	resp := models.Register(username, password)

	if resp.Status != 200 {
		return c.JSON(resp.Status, resp)
	}

	resp = utils.Respon{
		Status:  resp.Status,
		Message: resp.Message,
	}
	return c.JSON(http.StatusOK, resp)
}
func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)

	sess.Values["Auth"] = false
	sess.Values["id"] = ""
	sess.Values["username"] = ""
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusFound, "/")
}