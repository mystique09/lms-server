package controller

import (
	domain "server/domain/classroom"

	"github.com/labstack/echo/v4"
)

type ClassroomController struct {
	GetUsecase    domain.GetClassroomUsecase
	CreateUsecase domain.CreateClassroomUsecase
	UpdateUsecase domain.UpdateClassroomUsecase
	DeleteUsecase domain.DeleteClassroomUsecase
}

func (clc *ClassroomController) GetClassrooms(c echo.Context) error {
	return c.String(200, "[GET] TODO!")
}

func (clc *ClassroomController) GetClassroom(c echo.Context) error {
	return c.String(200, "[GET:all] TODO!")
}

func (clc *ClassroomController) CreateClassroom(c echo.Context) error {
	return c.String(200, "[POST] TODO!")
}

func (clc *ClassroomController) UpdateClassroom(c echo.Context) error {
	return c.String(200, "[PUT] TODO!")
}

func (clc *ClassroomController) DeleteClassroom(c echo.Context) error {
	return c.String(200, "[DELETE] TODO!")
}
