package middlewares

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"poliklinik/router"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthSessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		if isPublicRoute(ctx.Request().RequestURI) {
			return next(ctx)
		}
		fmt.Println("path: ", ctx.Request().RequestURI)
		gob.Register(map[string]interface{}{})
		sess, errS := session.Get("session", ctx)
		if errS != nil {
			ctx.Logger().Error(errS)
			return ctx.Redirect(http.StatusSeeOther, "/")
		}

		auth := sess.Values["Auth"]
		if auth != true {
			fmt.Println(auth)
			fmt.Println("auth false")
			return ctx.Redirect(http.StatusSeeOther, "/")
		}

		return next(ctx)
	}
}

func isPublicRoute(path string) bool {

	for _, route := range router.PublicRoutes {
		fmt.Println("route: ", route)
		if route == path {
			return true
		}
	}
	return false
}
