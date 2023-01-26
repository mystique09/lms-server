package route

import (
	"server/api/controller"
	"server/bootstrap"
	"server/database/store"
	"server/repository"
	"server/usecase/classroom"

	"github.com/labstack/echo/v4"
)

func NewClassroomRouter(app *bootstrap.Application, st store.Store, group *echo.Group) {
	clr := repository.NewClassroomRepository(st)
	cls := &controller.ClassroomController{
		GetUsecase:    usecase.NewGetClassroomUsecase(clr),
		CreateUsecase: usecase.NewCreateClassroomUsecase(clr),
		UpdateUsecase: usecase.NewUpdateClassroomUsecase(clr),
		DeleteUsecase: usecase.NewDeleteClassroomUsecase(clr),
	}

	group.GET("", cls.GetClassrooms)
	group.GET("/:id", cls.GetClassroom)
	group.POST("", cls.CreateClassroom)
	group.PUT("/:id", cls.UpdateClassroom)
	group.DELETE("/:id", cls.DeleteClassroom)
}
