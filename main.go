package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"poliklinik/config"
	"poliklinik/middlewares"
	"poliklinik/router"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stnc/pongo4echo"
)

func init() {
	// os.Setenv("HOST", "localhost")
	// os.Setenv("PORT", "1323")
	os.Setenv("SECRET_KEY", "secret")
	os.Setenv("GO_ENV", "development")
}

func main() {
	config := config.LoadConfig(".")
	e := echo.New()
	pongo := pongo4echo.Renderer{}
	if os.Getenv("GO_ENV") != "production" {
		pongo.Debug = true
		e.Debug = true
	}

	e.Renderer = pongo

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} ${uri} - ${method} ${status} ${latency_human}\n",
		CustomTimeFormat: "2006/01/02 15:04:05.00000",
	}))
	cookie := sessions.NewFilesystemStore("", []byte(os.Getenv("SECRET_KEY")))
	cookie.MaxLength(100000)
	storesess := &session.Config{
		Store: cookie,
	}

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))))
	e.Use(session.MiddlewareWithConfig(*storesess))

	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		// XFrameOptions:      "ALLOW-FROM http://blumb.id",
	}))

	//public folder
	e.Static("/public", "public")

	// public router
	router.Home(e)
	router.Auth(e)

	//use middleware for auth
	e2 := e
	e2.Use(middlewares.AuthSessionMiddleware)

	// need authorization router
	router.Dokter(e2)
	router.Pasien(e2)
	router.Periksa(e2)



	host := fmt.Sprintf("%s:%v", config.ServerHost, config.ServerPort)
	fmt.Println("Server started on", host)
	srv := http.Server{
		Addr:    host,
		Handler: e,
	}

	log.Printf("server started on %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// e.Logger.Fatal(err)
		fmt.Println("Error main", err.Error())
	}
}
