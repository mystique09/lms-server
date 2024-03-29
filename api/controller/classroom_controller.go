package controller

import (
	"database/sql"
	"net/http"
	"server/database/postgresql"
	"server/domain"
	"server/internal/tokenutil"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ClassroomController struct {
	ClassroomUsecase domain.ClassroomUsecase
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

	classrooms, err := clc.ClassroomUsecase.GetClasroomsByUser(c.Request().Context(), postgresql.GetAllClassFromUserParams{
		AdminID: payload.UserID,
		Offset:  int32(offset * 10),
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

	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewError(err.Error()))
	}

	classroom, err := clc.ClassroomUsecase.GetByID(c.Request().Context(), class_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusBadRequest, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(http.StatusInternalServerError, domain.NewError(domain.INTERNAL_ERROR))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_ONE, classroom))
}

func (clc *ClassroomController) JoinClassroom(c echo.Context) error {
	var code = c.Param("code")
	inviteCode, err := uuid.Parse(code)
	if err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	classroom, err := clc.ClassroomUsecase.GetByInviteCode(c.Request().Context(), inviteCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	members, err := clc.ClassroomUsecase.GetClassroomMembers(c.Request().Context(), classroom)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	payload, ok := c.Get("user").(*tokenutil.Payload)
	if !ok {
		return c.JSON(401, domain.NewError(domain.NO_PAYLOAD))
	}

	var hasJoined = false
	for i := range members {
		member := members[i]

		if member.UserID == payload.UserID {
			hasJoined = true
		}
	}

	if hasJoined {
		return c.JSON(200, domain.OkResponse("you already joined the classroom", inviteCode))
	}

	member, err := clc.ClassroomUsecase.JoinClassroom(c.Request().Context(), postgresql.AddNewClassroomMemberParams{
		ID:      uuid.New(),
		ClassID: classroom,
		UserID:  payload.UserID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(err.Error()))
	}

	return c.JSON(200, domain.OkResponse("you joined the classroom", member))
}

func (clc *ClassroomController) LeaveClassroom(c echo.Context) error {
	var id = c.Param("id")
	classId, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	payload, ok := c.Get("user").(*tokenutil.Payload)
	if !ok {
		return c.JSON(400, domain.NewError(domain.NO_PAYLOAD))
	}

	member, err := clc.ClassroomUsecase.LeaveClassroom(c.Request().Context(), postgresql.LeaveClassroomParams{
		UserID:  payload.UserID,
		ClassID: classId,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(err.Error()))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_DELETE, member))
}

func (clc *ClassroomController) GetMembers(c echo.Context) error {
	var id = c.Param("id")
	class_id, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewError(err.Error()))
	}

	members, err := clc.ClassroomUsecase.GetClassroomMembers(c.Request().Context(), class_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(404, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_ALL, members))
}

func (clc *ClassroomController) CreateClassroom(c echo.Context) error {
	var request domain.CreateClassroomRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewError(err.Error()))
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewError(err.Error()))
	}

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

	if err := clc.ClassroomUsecase.Create(c.Request().Context(), &classroom_arg); err != nil {
		return c.JSON(400, domain.NewError(domain.ERROR_DEFAULT))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_CREATE, &classroom_arg))
}

func (clc *ClassroomController) UpdateClassroom(c echo.Context) error {
	var request domain.UpdateClassroomRequest
	var id string = c.Param("id")

	class_id, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	var query = c.QueryParam("fields")
	fields := strings.Split(query, ",")

	if len(fields) == 0 || len(fields) > 5 {
		return c.JSON(400, domain.NewError("fields must be one of: [name, description, room, section, and subject]"))
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	for i := range fields {
		field := fields[i]
		if !(field == "name" || field == "description" || field == "subject" || field == "room" || field == "section") {
			return c.JSON(400, domain.NewError("field must be one of: [name, description, room, section, and subject]"))
		}
	}

	payload, ok := c.Get("user").(*tokenutil.Payload)
	if !ok {
		return c.JSON(400, domain.NewError(domain.NO_PAYLOAD))
	}

	classroom, err := clc.ClassroomUsecase.GetByID(c.Request().Context(), class_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(404, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	if classroom.AdminID != payload.UserID {
		return c.JSON(401, domain.NewError(domain.UNAUTHORIZED))
	}

	for i := range fields {
		field := fields[i]
		if field == "name" {
			if err := clc.ClassroomUsecase.UpdateClassroomName(c.Request().Context(), &postgresql.UpdateClassroomNameParams{
				Name: request.Name,
				ID:   class_id,
			}); err != nil {
				if err == sql.ErrNoRows {
					return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
				}
				return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
			}
		}

		if field == "description" {
			if err := clc.ClassroomUsecase.UpdateClassroomDescription(c.Request().Context(), &postgresql.UpdateClassroomDescriptionParams{
				Description: request.Description,
				ID:          class_id,
			}); err != nil {
				if err == sql.ErrNoRows {
					return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
				}
				return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
			}
		}

		if field == "subject" {
			if err := clc.ClassroomUsecase.UpdateClassroomSubject(c.Request().Context(), &postgresql.UpdateClassroomSubjectParams{
				Subject: request.Subject,
				ID:      class_id,
			}); err != nil {
				if err == sql.ErrNoRows {
					return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
				}
				return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
			}
		}

		if field == "room" {
			if err := clc.ClassroomUsecase.UpdateClassroomRoom(c.Request().Context(), &postgresql.UpdateClassroomRoomParams{
				Room: request.Room,
				ID:   class_id,
			}); err != nil {
				if err == sql.ErrNoRows {
					return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
				}
				return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
			}
		}

		if field == "section" {
			if err := clc.ClassroomUsecase.UpdateClassroomSection(c.Request().Context(), &postgresql.UpdateClassroomSectionParams{
				Section: request.Section,
				ID:      class_id,
			}); err != nil {
				if err == sql.ErrNoRows {
					return c.JSON(400, domain.NewError(domain.RESOURCE_NOT_FOUND))
				}
				return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
			}
		}
	}

	return c.JSON(200, domain.OkResponse(domain.OK_UPDATE, class_id))
}

func (clc *ClassroomController) DeleteClassroom(c echo.Context) error {
	var id = c.Param("id")
	class_id, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, domain.NewError(err.Error()))
	}

	payload, ok := c.Get("user").(*tokenutil.Payload)
	if !ok {
		return c.JSON(400, domain.NewError(domain.NO_PAYLOAD))
	}

	classroom, err := clc.ClassroomUsecase.GetByID(c.Request().Context(), class_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(404, domain.NewError(domain.RESOURCE_NOT_FOUND))
		}
		return c.JSON(500, domain.NewError(domain.INTERNAL_ERROR))
	}

	if classroom.AdminID != payload.UserID {
		return c.JSON(401, domain.NewError(domain.UNAUTHORIZED))
	}

	if err := clc.ClassroomUsecase.Delete(c.Request().Context(), class_id); err != nil {
		return c.JSON(400, domain.NewError(domain.ERROR_DEFAULT))
	}

	return c.JSON(200, domain.OkResponse(domain.OK_DELETE, class_id))
}
