package route

import (
	"server/api/controller"
	"server/bootstrap"
	"server/database/store"
	"server/usecase"

	"github.com/labstack/echo/v4"
)

func NewAccessTokenRouter(app *bootstrap.Application, st store.Store, group *echo.Group) {
	atc := &controller.AccessTokenController{
		AccessTokenUsecase: usecase.NewAccessTokenUsecase(app.TokenMaker),
		Env:                &app.Env,
	}

	group.POST("/validate_token", atc.ValidateAccessToken)
}
