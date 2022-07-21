package routes

import (
	"net/http"
	database "server/database/sqlc"
	"server/utils"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateClassroomDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Section     string `json:"section"`
	Room        string `json:"room"`
	Subject     string `json:"subject"`
}

type UpdateClassroomDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Section     string    `json:"section"`
	Room        string    `json:"room"`
	Subject     string    `json:"subject"`
	InviteCode  uuid.UUID `json:"invite_code"`
}

func (s *Server) getClassrooms(c echo.Context) error {
	classes, err := s.DB.ListClass(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(classes, ""))
}

func (s *Server) getAllClassrooms(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) getClassroom(c echo.Context) error {
	uid := c.Param("id")
	uuid, err := uuid.Parse(uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	class, err := s.DB.GetClass(c.Request().Context(), uuid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(http.StatusBadRequest, utils.NewResponse(class, ""))
}

func (s *Server) createNewClassroom(c echo.Context) error {
	var payload CreateClassroomDTO
	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	class_param := database.CreateClassParams{
		ID:          uuid.New(),
		AdminID:     jwt_payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
		Section:     payload.Section,
		Room:        payload.Room,
		Subject:     payload.Subject,
		InviteCode:  uuid.New(),
		Visibility:  database.VisibilityPUBLIC,
	}

	new_classroom, err := s.DB.CreateClass(c.Request().Context(), class_param)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(new_classroom, ""))
}

func (s *Server) updateClassroom(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	var payload UpdateClassroomDTO
	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	classroom, err := s.DB.GetClass(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Classroom not found."))
	}

	token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(token)

	if classroom.AdminID != jwt_payload.ID {
		return c.JSON(http.StatusUnauthorized, utils.NewResponse(nil, "You are not authorized to perform this action."))
	}

	update_class_param := database.UpdateClassParams{
		Name:        payload.Name,
		Description: payload.Description,
		Section:     payload.Section,
		Room:        payload.Room,
		Subject:     payload.Subject,
		InviteCode:  payload.InviteCode,
	}

	updated_classroom, err := s.DB.UpdateClass(c.Request().Context(), update_class_param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(updated_classroom, ""))
}

func (s *Server) deleteClassroom(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(token)

	classroom, err := s.DB.GetClass(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	if jwt_payload.ID != classroom.AdminID {
		return c.JSON(http.StatusUnauthorized, utils.NewResponse(nil, "You are not authorized to perform this action."))
	}

	deleted_classroom, err := s.DB.DeleteClass(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(deleted_classroom, ""))
}

func (s *Server) getClassroomUsers(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) getClassroomPosts(c echo.Context) error {
	return c.String(200, "TODO")
}
