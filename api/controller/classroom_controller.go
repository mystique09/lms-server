package controller

import (
	"database/sql"
	"log"
	"net/http"
	"server/database/postgresql"
	"server/domain"
	"server/internal/tokenutil"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ClassroomController struct {
	GetUsecase    domain.GetClassroomUsecase
	CreateUsecase domain.CreateClassroomUsecase
	UpdateUsecase domain.UpdateClassroomUsecase
	DeleteUsecase domain.DeleteClassroomUsecase
}

func (clc *ClassroomController) GetClassrooms(c echo.Context) error {
	query := c.QueryParam("offset")

	if query == "" || query == "1" {
		query = "0"
	}

	offset, err := strconv.Atoi(query)

	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewError(err.Error()))
	}

	payload, ok := c.Get("user").(*tokenutil.Payload)
	if !ok {
		return c.JSON(http.StatusBadRequest, domain.NewError(domain.NO_PAYLOAD))
	}

	classrooms, err := clc.GetUsecase.GetClasroomByUser(c.Request().Context(), postgresql.GetAllClassFromUserParams{
		AdminID: payload.UserID,
		Offset:  int32(offset) * 10,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusBadRequest, domain.NewError(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, domain.NewError(err.Error()))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_ALL, classrooms))
}

func (clc *ClassroomController) GetClassroom(c echo.Context) error {
	id := c.Param("id")
	class_id, err := uuid.Parse(id)

	log.Println(class_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewError(err.Error()))
	}

	classroom, err := clc.GetUsecase.GetByID(c.Request().Context(), class_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusBadRequest, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(http.StatusInternalServerError, domain.NewError(domain.INTERNAL_ERROR))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_ONE, classroom))
}

func (clc *ClassroomController) CreateClassroom(c echo.Context) error {
	var request domain.CreateClassroomRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewError(err.Error()))
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewError(err.Error()))
	}

	// TODO: checked whether the current in the header is the admin in the request
	payload, ok := c.Get("user").(*tokenutil.Payload)
	if !ok {
		return c.JSON(http.StatusBadRequest, domain.NewError(domain.NO_PAYLOAD))
	}

	if request.AdminID != payload.UserID {
		return c.JSON(http.StatusBadRequest, domain.NewError(domain.UNAUTHORIZED))
	}

	classroom_arg := postgresql.CreateClassParams{
		ID:          uuid.New(),
		AdminID:     request.AdminID,
		Name:        request.Name,
		Description: request.Description,
		Section:     request.Section,
		Subject:     request.Subject,
		Room:        request.Room,
		InviteCode:  uuid.New(),
		Visibility:  postgresql.VisibilityPUBLIC,
	}

	if err := clc.CreateUsecase.Create(c.Request().Context(), &classroom_arg); err != nil {
		return c.JSON(400, domain.NewError(domain.ERROR_DEFAULT))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_CREATE, &classroom_arg))
}

func (clc *ClassroomController) UpdateClassroom(c echo.Context) error {
	return c.String(200, "[PUT] TODO!")
}

func (clc *ClassroomController) DeleteClassroom(c echo.Context) error {
	return c.String(200, "[DELETE] TODO!")
}
