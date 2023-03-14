package main

import (
	"os"
	"server/api/middleware"
	"server/api/route/v1"
	"server/bootstrap"
	"server/ui"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()
	logger := zerolog.New(os.Stdout)
	renderer := ui.NewTemplateRenderer("ui/public/**.*.html")

	e := echo.New()
	e.Renderer = renderer
	e.Static("/static", "ui/static")
	e.Use(middleware.LoggerMiddleware(&logger))
	e.Use(middleware.CorsMiddleware(&app.Env))
	e.Use(middleware.RateLimitMiddleware(20))
	e.Validator = bootstrap.NewValidator()

	router := e.Group("")

	route.Setup(&app, app.Store, router)
	app.Launch(e)
}
