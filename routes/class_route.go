package routes

import (
	"database/sql"
	"net/http"
	database "server/database/sqlc"
	"server/utils"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateClassroomDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Section     string `json:"section"`
	Room        string `json:"room"`
	Subject     string `json:"subject"`
}

type UpdateClassroomName struct {
	Name string `json:"name" validate:"required"`
}

type UpdateClassroomDescription struct {
	Description string `json:"description" validate:"required"`
}

type UpdateClassroomSubject struct {
	Subject string `json:"name" validate:"required"`
}

type UpdateClassroomSection struct {
	Section string `json:"section" validate:"required"`
}

type UpdateClassroomRoom struct {
	Room string `json:"room" validate:"required"`
}

type ClassroomResponse struct {
	*database.Classroom
	Members []database.ClassroomMember `json:"members"`
}

type ClassroomJoinRequest struct {
	InviteCode uuid.UUID `json:"invite_code" validate:"uuid"`
}

func (s *Server) getClassrooms(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, err)
	}
	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)
	if offset < 0 {
		return c.JSON(400, NEGATIVE_OFFSET)
	}

	if err != nil {
		return c.JSON(400, err)
	}

	param := database.GetAllClassFromUserParams{
		AdminID: uid,
		Offset:  int32(offset) * 10,
	}

	classes, err := s.DB.GetAllClassFromUser(c.Request().Context(), param)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, classes)
}

func (s *Server) getAllClassrooms(c echo.Context) error {
	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)

	if err != nil {
		return c.JSON(400, err)
	}

	if offset < 0 {
		return c.JSON(400, NEGATIVE_OFFSET)
	}

	public_classrooms, err := s.DB.ListAllPublicClass(c.Request().Context(), int32(offset*10))

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, public_classrooms)
}

func (s *Server) getClassroom(c echo.Context) error {
	uid := c.Param("id")
	uuid, err := uuid.Parse(uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	class, err := s.DB.GetClass(c.Request().Context(), uuid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusBadRequest, class)
}

func (s *Server) createNewClassroom(c echo.Context) error {
	var payload CreateClassroomDTO

	c.Bind(&payload)
	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err)
	}

	jwt_token := c.Get("user")
	jwt_payload := utils.GetPayloadFromJwt(jwt_token.(*jwt.Token))

	new_classroom, err := s.DB.CreateClass(c.Request().Context(), database.CreateClassParams{
		ID:          uuid.New(),
		AdminID:     jwt_payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
		Section:     payload.Section,
		Room:        payload.Room,
		Subject:     payload.Subject,
		InviteCode:  uuid.New(),
		Visibility:  database.VisibilityPUBLIC,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	new_member, err := s.DB.AddNewClassroomMember(c.Request().Context(), database.AddNewClassroomMemberParams{
		ID:      uuid.New(),
		ClassID: new_classroom.ID,
		UserID:  new_classroom.AdminID,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ClassroomResponse{
		Classroom: &new_classroom,
		Members:   []database.ClassroomMember{new_member},
	})
}

func (s *Server) updateClassroom(c echo.Context) error {
	id := c.Param("id")
	field := c.QueryParam("field")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	classroom, err := s.DB.GetClass(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, CLASSROOM_NOTFOUND)
	}

	token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(token)

	if classroom.AdminID != jwt_payload.ID {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	switch field {
	case "name":
		var payload UpdateClassroomName

		c.Bind(&payload)
		if err := c.Validate(payload); err != nil {
			return c.JSON(400, err)
		}

		updated_class, err := s.DB.UpdateClassroomName(c.Request().Context(), database.UpdateClassroomNameParams{
			ID:   uid,
			Name: payload.Name,
		})

		if err != nil {
			return c.JSON(400, err)
		}
		return c.JSON(200, updated_class)

	case "description":
		var payload UpdateClassroomDescription

		c.Bind(&payload)
		if err := c.Validate(payload); err != nil {
			return c.JSON(400, err)
		}

		updated_class, err := s.DB.UpdateClassroomDescription(c.Request().Context(), database.UpdateClassroomDescriptionParams{
			ID:          uid,
			Description: payload.Description,
		})

		if err != nil {
			return c.JSON(400, err)
		}
		return c.JSON(200, updated_class)

	case "subject":
		var payload UpdateClassroomSubject

		c.Bind(&payload)
		if err := c.Validate(payload); err != nil {
			return c.JSON(400, err)
		}

		updated_class, err := s.DB.UpdateClassroomSubject(c.Request().Context(), database.UpdateClassroomSubjectParams{
			ID:      uid,
			Subject: payload.Subject,
		})

		if err != nil {
			return c.JSON(400, err)
		}
		return c.JSON(200, updated_class)

	case "section":
		var payload UpdateClassroomSection

		c.Bind(&payload)
		if err := c.Validate(payload); err != nil {
			return c.JSON(400, err)
		}

		updated_class, err := s.DB.UpdateClassroomSection(c.Request().Context(), database.UpdateClassroomSectionParams{
			ID:      uid,
			Section: payload.Section,
		})

		if err != nil {
			return c.JSON(400, err)
		}
		return c.JSON(200, updated_class)

	case "room":
		var payload UpdateClassroomRoom

		c.Bind(&payload)
		if err := c.Validate(payload); err != nil {
			return c.JSON(400, err)
		}

		updated_class, err := s.DB.UpdateClassroomRoom(c.Request().Context(), database.UpdateClassroomRoomParams{
			ID:   uid,
			Room: payload.Room,
		})

		if err != nil {
			return c.JSON(400, err)
		}
		return c.JSON(200, updated_class)

	case "invite_code":
		updated_class, err := s.DB.UpdateClassroomInviteCode(c.Request().Context(), database.UpdateClassroomInviteCodeParams{
			ID:         uid,
			InviteCode: uuid.New(),
		})

		if err != nil {
			return c.JSON(400, err)
		}
		return c.JSON(200, updated_class)
	}

	return c.JSON(400, UNKNOWN_FIELD)
}

func (s *Server) deleteClassroom(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(token)

	classroom, err := s.DB.GetClass(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if jwt_payload.ID != classroom.AdminID {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	deleted_classroom, err := s.DB.DeleteClass(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, deleted_classroom)
}

func (s *Server) getClassroomUsers(c echo.Context) error {
	id := c.Param("id")
	page := c.QueryParam("page")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, err)
	}

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)
	if offset < 0 {
		return c.JSON(400, NEGATIVE_OFFSET)
	}

	members, err := s.DB.GetAllClassroomMembers(c.Request().Context(), database.GetAllClassroomMembersParams{
		ClassID: uid,
		Offset:  int32(offset * 10),
	})

	if err != nil {
		return c.JSON(400, CLASSROOM_NOTFOUND)
	}
	return c.JSON(200, members)
}

func (s *Server) joinClassroom(c echo.Context) error {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, err)
	}

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)
	if err == sql.ErrNoRows && check_user.ID == uuid.Nil {
		return c.JSON(400, USER_NOTFOUND)
	}

	var payload ClassroomJoinRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(400, err)
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err)
	}

	classroom_id, err := s.DB.GetClassroomWithInviteCode(c.Request().Context(), payload.InviteCode)
	if err == sql.ErrNoRows && classroom_id == uuid.Nil {
		return c.JSON(400, CLASSROOM_NOTFOUND)
	}

	token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(token)

	if jwt_payload.ID != uid {
		return c.JSON(400, UNAUTHORIZED)
	}

	joined, err := s.DB.AddNewClassroomMember(c.Request().Context(), database.AddNewClassroomMemberParams{
		ID:      uuid.New(),
		ClassID: classroom_id,
		UserID:  uid,
	})
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, joined)
}

func (s *Server) leaveClassroom(c echo.Context) error {
	id := c.Param("id")
	class_id := c.Param("class_id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, err)
	}

	parsed_classId, err := uuid.Parse(class_id)
	if err != nil {
		return c.JSON(400, err)
	}

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)
	if err == sql.ErrNoRows && check_user.ID == uuid.Nil {
		return c.JSON(400, USER_NOTFOUND)
	}

	check_class, err := s.DB.GetClass(c.Request().Context(), parsed_classId)
	if err == sql.ErrNoRows && check_class.ID == uuid.Nil {
		return c.JSON(400, CLASSROOM_NOTFOUND)
	}

	token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(token)

	if jwt_payload.ID != uid {
		return c.JSON(400, UNAUTHORIZED)
	}

	leaved, err := s.DB.LeaveClassroom(c.Request().Context(), database.LeaveClassroomParams{
		UserID:  jwt_payload.ID,
		ClassID: parsed_classId,
	})
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, leaved)
}

func (s *Server) getClassroomPosts(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, err)
	}

	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)

	if err != nil {
		return c.JSON(400, err)
	}

	if offset < 0 {
		return c.JSON(400, NEGATIVE_OFFSET)
	}

	posts, err := s.DB.ListAllPostsFromClass(c.Request().Context(), database.ListAllPostsFromClassParams{
		ClassID: uid,
		Offset:  int32(offset * 10),
	})

	return c.JSON(200, posts)
}
