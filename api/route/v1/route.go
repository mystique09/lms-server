package route

import (
	"server/api/middleware"
	"server/bootstrap"
	"server/database/store"

	"github.com/labstack/echo/v4"
)

func Setup(app *bootstrap.Application, st store.Store, routeV1 *echo.Group) {
	publicRouterV1 := routeV1.Group("")
	publicRouterV1.GET("", indexRoute)
	publicRouterV1.GET("/health", healthRoute)
	NewLoginRouter(app, st, publicRouterV1)
	NewSignupRouter(app, st, publicRouterV1)
	NewRefreshTokenRouter(app, st, publicRouterV1)
	NewAccessTokenRouter(app, st, publicRouterV1)

	userGroup := routeV1.Group("/profile", middleware.AuthMiddleware(app.TokenMaker))
	NewProfileRouter(app, st, userGroup)

	classroomsGroup := routeV1.Group("/classrooms", middleware.AuthMiddleware(app.TokenMaker))
	NewClassroomRouter(app, st, classroomsGroup)
}

func indexRoute(c echo.Context) error {
	return c.HTML(200, `
		Welcome, you are in the backend of lms-server. <br> 
		<a href="/api/v1/health">Server status</a> <br>
		<a href="/api/v1/docs">Visit APIv1 Documentation</a>		
	`)
}

func healthRoute(c echo.Context) error {
	return c.String(200, `{health: 100, status: "good"}`)
}
