package router

import (
	"poliklinik/controllers"

	"github.com/labstack/echo/v4"
)

var PublicRoutes = []string{
	"/",
	"/login",
	"/register",
}

func Home(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(302, "/login")
	})
}
func Auth(e *echo.Echo) {
	e.GET("/login", controllers.Login)
	e.GET("/register", controllers.Register)
	e.POST("/register", controllers.RegisterPost)
}
