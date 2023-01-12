package main

import (
	"server/api/route/v1"
	"server/bootstrap"

	"github.com/labstack/echo/v4"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	defer app.CloseDBConnection()

	e := echo.New()
	routeV1 := e.Group("/api/v1")

	route.Setup(&env, app.Store, routeV1)
	route.Launch(&env, e)
}
