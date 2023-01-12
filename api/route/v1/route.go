package route

import (
	"log"
	"server/bootstrap"
	"server/database/store"

	"github.com/labstack/echo/v4"
)

// Serves API from routes
func Setup(env *bootstrap.Env, st store.Store, routeV1 *echo.Group) {
	publicRouterV1 := routeV1.Group("")
	// TODO!
	// remove this one
	publicRouterV1.GET("/health", func(c echo.Context) error {
		return c.JSON(200, `{"health":"100", "status": "good"}`)
	})

	protectedRouterV1 := routeV1.Group("")
	// TODO!
	// remove this one
	protectedRouterV1.GET("/protected", func(c echo.Context) error {
		return c.JSON(200, `{"message": "protected"}`)
	})
}

func Launch(env *bootstrap.Env, router *echo.Echo) {
	log.Fatal(router.Start(env.Host))
}
