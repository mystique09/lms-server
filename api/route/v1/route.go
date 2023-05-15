package route

import (
	"server/api/middleware"
	"server/bootstrap"
	"server/database/store"

	"github.com/labstack/echo/v4"
)

func Setup(app *bootstrap.Application, st store.Store, router *echo.Group) {
	publicRouterV1 := router.Group("/api/v1")
	privateRouter := router.Group("/api/v1", middleware.AuthMiddleware(app.TokenMaker))
	uiRouter := router.Group("")

	publicRouterV1.GET("/health", healthRoute)
	NewLoginRouter(app, st, publicRouterV1)
	NewSignupRouter(app, st, publicRouterV1)
	NewRefreshTokenRouter(app, st, publicRouterV1)
	NewAccessTokenRouter(app, st, publicRouterV1)

	userGroup := privateRouter.Group("/users")
	NewProfileRouter(app, st, userGroup)

	classroomsGroup := privateRouter.Group("/classrooms")
	NewClassroomRouter(app, st, classroomsGroup)

	uiRouter.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.page.html", nil)
	})
}

func healthRoute(c echo.Context) error {
	return c.String(200, `{health: 100, status: "good"}`)
}
