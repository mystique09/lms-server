package main

import (
	"server/api/route/v1"
	"server/bootstrap"

	"github.com/labstack/echo/v4"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()

	e := echo.New()
	e.Validator = bootstrap.NewValidator()
	routeV1 := e.Group("/api/v1")

	route.Setup(&app, app.Store, routeV1)
	app.Launch(e)
}
