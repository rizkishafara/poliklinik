package router

import (
	"poliklinik/controllers"

	"github.com/labstack/echo/v4"
)

var PublicRoutes = []string{
	"/",
	"/login",
	"/register",
	"/logout",
}

func Home(e *echo.Echo) {
	// e.GET("/", func(c echo.Context) error {
	// 	return c.Redirect(302, "/login")
	// })
	e.GET("/", controllers.Home)
}
func Auth(e *echo.Echo) {
	e.GET("/login", controllers.Login)
	e.POST("/login", controllers.LoginPost)
	e.GET("/register", controllers.Register)
	e.POST("/register", controllers.RegisterPost)
	e.GET("/logout", controllers.Logout)
}
func Dokter(e *echo.Echo) {
	e.GET("/dokter", controllers.PageDokter)
	e.GET("/dokter/get", controllers.GetDokter)
	e.POST("/dokter/create", controllers.CreateDokter)
	e.POST("/dokter/update", controllers.UpdateDokter)
	e.POST("/dokter/delete", controllers.DeleteDokter)
}
func Pasien(e *echo.Echo) {
	e.GET("/pasien", controllers.PagePasien)
	e.GET("/pasien/get", controllers.GetPasien)
	e.POST("/pasien/create", controllers.CreatePasien)
	e.POST("/pasien/update", controllers.UpdatePasien)
	e.POST("/pasien/delete", controllers.DeletePasien)
}
func Periksa(e *echo.Echo) {
	e.GET("/periksa", controllers.PagePeriksa)
	e.GET("/periksa/get", controllers.GetPeriksa)
	e.POST("/periksa/create", controllers.CreatePeriksa)
	e.POST("/periksa/update", controllers.UpdatePeriksa)
	e.POST("/periksa/delete", controllers.DeletePeriksa)
}
