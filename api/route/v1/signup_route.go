package route

import (
	"server/api/controller"
	"server/bootstrap"
	"server/database/store"
	"server/repository"
	"server/usecase"

	"github.com/labstack/echo/v4"
)

func NewSignupRouter(app *bootstrap.Application, st store.Store, group *echo.Group) {
	ur := repository.NewUserRepository(st)
	sc := &controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur),
		Env:           &app.Env,
	}

	group.POST("/signup", sc.Signup)
}
