package route

import (
	"server/bootstrap"
	"server/database/store"

	"github.com/labstack/echo/v4"
)

func Setup(app *bootstrap.Application, st store.Store, routeV1 *echo.Group) {
	publicRouterV1 := routeV1.Group("")
	publicRouterV1.GET("/health", healthRoute)
	NewLoginRouter(app, st, publicRouterV1)

	protectedRouterV1 := routeV1.Group("")
	// TODO!
	// remove this one
	protectedRouterV1.GET("/protected", func(c echo.Context) error {
		return c.JSON(200, `{"message": "protected"}`)
	})
}

func healthRoute(c echo.Context) error {
	return c.JSON(200, `{health: 100, status: "good"}`)
}
