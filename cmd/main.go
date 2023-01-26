package main

import (
	"os"
	"server/api/middleware"
	"server/api/route/v1"
	"server/bootstrap"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()
	logger := zerolog.New(os.Stdout)

	e := echo.New()
	e.Use(middleware.LoggerMiddleware(&logger))
	e.Use(middleware.CorsMiddleware(&app.Env))
	e.Use(middleware.RateLimitMiddleware(20))
	e.Validator = bootstrap.NewValidator()

	routeV1 := e.Group("/api/v1")

	route.Setup(&app, app.Store, routeV1)
	app.Launch(e)
}
