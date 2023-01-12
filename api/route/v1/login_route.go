package route

import (
	"server/api/controller"
	"server/bootstrap"
	"server/database/store"
	"server/repository"
	"server/usecase"

	"github.com/labstack/echo/v4"
)

func NewLoginRouter(app *bootstrap.Application, st store.Store, group *echo.Group) {
	ur := repository.NewUserRepository(st)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, app.TokenMaker),
		Env:          &app.Env,
	}

	group.POST("/login", lc.Login)
}
