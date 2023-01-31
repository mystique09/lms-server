package route

import (
	"server/api/controller"
	"server/bootstrap"
	"server/database/store"
	"server/repository"
	"server/usecase"

	"github.com/labstack/echo/v4"
)

func NewClassroomRouter(app *bootstrap.Application, st store.Store, group *echo.Group) {
	clr := repository.NewClassroomRepository(st)
	clc := &controller.ClassroomController{
		ClassroomUsecase: usecase.NewClassroomUsecase(clr),
	}

	group.GET("", clc.GetClassrooms)
	group.GET("/:id", clc.GetClassroom)
	group.GET("/:id/members", clc.GetMembers)
	group.POST("", clc.CreateClassroom)
	group.PUT("/:id", clc.UpdateClassroom)
	group.DELETE("/:id", clc.DeleteClassroom)
}
