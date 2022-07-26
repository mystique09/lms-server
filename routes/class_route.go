package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	database "server/database/sqlc"
	"server/utils"
	"strconv"

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

type ClassroomResponse struct {
	*database.Classroom
	Members []database.ClassroomMember `json:"members"`
}

type ClassroomJoinRequest struct {
	InviteCode uuid.UUID `json:"invite_code"`
}

func (s *Server) getClassrooms(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}
	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)
	if offset < 0 {
		return c.JSON(400, utils.NewResponse(nil, "offset must not be negative"))
	}

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	param := database.GetAllClassFromUserParams{
		AdminID: uid,
		Offset:  int32(offset) * 10,
	}

	classes, err := s.DB.GetAllClassFromUser(c.Request().Context(), param)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(classes, ""))
}

func (s *Server) getAllClassrooms(c echo.Context) error {
	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	if offset < 0 {
		return c.JSON(400, utils.NewResponse(nil, "offset must not be negative"))
	}

	public_classrooms, err := s.DB.ListAllPublicClass(c.Request().Context(), int32(offset*10))

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(200, utils.NewResponse(public_classrooms, ""))
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

	jwt_token := c.Get("user")
	jwt_payload := utils.GetPayloadFromJwt(jwt_token.(*jwt.Token))

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

	new_member, err := s.DB.AddNewClassroomMember(c.Request().Context(), database.AddNewClassroomMemberParams{
		ID:      uuid.New(),
		ClassID: new_classroom.ID,
		UserID:  new_classroom.AdminID,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(ClassroomResponse{
		Classroom: &new_classroom,
		Members:   []database.ClassroomMember{new_member},
	}, ""))
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
		ID:          uid,
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
	id := c.Param("id")
	page := c.QueryParam("page")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)
	if offset < 0 {
		return c.JSON(400, "page should not be negative")
	}

	members, err := s.DB.GetAllClassroomMembers(c.Request().Context(), database.GetAllClassroomMembersParams{
		ClassID: uid,
		Offset:  int32(offset * 10),
	})

	if err != nil {
		return c.JSON(400, fmt.Sprintf("class id [%v] doesn't exist", uid))
	}
	return c.JSON(200, utils.NewResponse(members, ""))
}

func (s *Server) joinClassroom(c echo.Context) error {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)
	if err == sql.ErrNoRows && check_user.ID == uuid.Nil {
		return c.JSON(400, fmt.Sprintf("user [%v] doesn't exist", uid))
	}

	var payload ClassroomJoinRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(400, err.Error())
	}

	classroom_id, err := s.DB.GetClassroomWithInviteCode(c.Request().Context(), payload.InviteCode)
	if err == sql.ErrNoRows && classroom_id == uuid.Nil {
		return c.JSON(400, fmt.Sprintf("classroom with invite code [%v] doesn't exist", payload.InviteCode))
	}

	token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(token)

	if jwt_payload.ID != uid {
		return c.JSON(400, fmt.Sprintf("[%v] are not authorized to perform this action", jwt_payload.ID))
	}

	joined, err := s.DB.AddNewClassroomMember(c.Request().Context(), database.AddNewClassroomMemberParams{
		ID:      uuid.New(),
		ClassID: classroom_id,
		UserID:  uid,
	})
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, utils.NewResponse(joined, ""))
}

func (s *Server) leaveClassroom(c echo.Context) error {
	id := c.Param("id")
	class_id := c.Param("class_id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	parsed_classId, err := uuid.Parse(class_id)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)
	if err == sql.ErrNoRows && check_user.ID == uuid.Nil {
		return c.JSON(400, fmt.Sprintf("user [%v] doesn't exist", uid))
	}

	check_class, err := s.DB.GetClass(c.Request().Context(), parsed_classId)
	if err == sql.ErrNoRows && check_class.ID == uuid.Nil {
		return c.JSON(400, fmt.Sprintf("classroom [%v] doesn't exist", parsed_classId))
	}

	token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(token)

	if jwt_payload.ID != uid {
		return c.JSON(400, fmt.Sprintf("[%v] are not authorized to perform this action", jwt_payload.ID))
	}

	leaved, err := s.DB.LeaveClassroom(c.Request().Context(), database.LeaveClassroomParams{
		UserID:  jwt_payload.ID,
		ClassID: parsed_classId,
	})
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(200, utils.NewResponse(leaved, ""))
}

func (s *Server) getClassroomPosts(c echo.Context) error {
	return c.String(200, "TODO")
}
