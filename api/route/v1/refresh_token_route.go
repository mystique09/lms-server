package route

import (
	"server/api/controller"
	"server/bootstrap"
	"server/database/store"
	"server/usecase"

	"github.com/labstack/echo/v4"
)

func NewRefreshTokenRouter(app *bootstrap.Application, st store.Store, group *echo.Group) {
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(app.TokenMaker),
		Env:                 &app.Env,
	}

	group.POST("/refresh_token", rtc.RefreshToken)
}
