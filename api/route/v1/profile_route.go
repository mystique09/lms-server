package route

import (
	"server/api/controller"
	"server/bootstrap"
	"server/database/store"
	"server/repository"
	"server/usecase"

	"github.com/labstack/echo/v4"
)

func NewProfileRouter(app *bootstrap.Application, st store.Store, group *echo.Group) {
	ur := repository.NewUserRepository(st)
	usrc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur),
	}

	group.GET("/:id", usrc.GetProfile)
	group.GET("/:id/classrooms", usrc.GetClassrooms)
}
