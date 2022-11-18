package routes

import (
	"database/sql"
	"net/http"
	database "server/database/sqlc"
	"server/token"
	"strconv"

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
		Offset:  int32(offset * 10),
	}

	classes, err := s.store.GetAllClassFromUser(c.Request().Context(), param)

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

	public_classrooms, err := s.store.ListAllPublicClass(c.Request().Context(), int32(offset*10))

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, public_classrooms)
}

func (s *Server) getClassroom(c echo.Context) error {
	uid := c.Param("id")
	user_id, err := uuid.Parse(uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	class, err := s.store.GetClass(c.Request().Context(), user_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusBadRequest, class)
}

func (s *Server) createNewClassroom(c echo.Context) error {
	var req CreateClassroomDTO

	bindErr := c.Bind(&req)
	if bindErr != nil {
		return c.JSON(400, bindErr)
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(400, err)
	}

	payload := c.Get("user").(*token.Payload)

	new_classroom, err := s.store.CreateClass(c.Request().Context(), database.CreateClassParams{
		ID:          uuid.New(),
		AdminID:     payload.ID,
		Name:        req.Name,
		Description: req.Description,
		Section:     req.Section,
		Room:        req.Room,
		Subject:     req.Subject,
		InviteCode:  uuid.New(),
		Visibility:  database.VisibilityPUBLIC,
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	new_member, err := s.store.AddNewClassroomMember(c.Request().Context(), database.AddNewClassroomMemberParams{
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
		return c.JSON(http.StatusBadRequest, newResponse[any](nil, err.Error()))
	}

	classroom, err := s.store.GetClass(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, CLASSROOM_NOTFOUND)
	}

	payload := c.Get("user").(*token.Payload)

	if classroom.AdminID != payload.ID {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	switch field {
	case "name":
		var data UpdateClassroomName

		bindErr := c.Bind(&data)
		if bindErr != nil {
			return c.JSON(400, bindErr)
		}

		if err := c.Validate(data); err != nil {
			return c.JSON(400, newResponse[any](nil, err.Error()))
		}

		updated_class, err := s.store.UpdateClassroomName(c.Request().Context(), database.UpdateClassroomNameParams{
			ID:   uid,
			Name: data.Name,
		})

		if err != nil {
			return c.JSON(400, newResponse[any](nil, err.Error()))
		}
		return c.JSON(200, updated_class)

	case "description":
		var data UpdateClassroomDescription

		bindErr := c.Bind(&data)
		if bindErr != nil {
			return c.JSON(400, newResponse[any](nil, bindErr.Error()))
		}

		if err := c.Validate(data); err != nil {
			return c.JSON(400, newResponse[any](nil, err.Error()))
		}

		updated_class, err := s.store.UpdateClassroomDescription(c.Request().Context(), database.UpdateClassroomDescriptionParams{
			ID:          uid,
			Description: data.Description,
		})

		if err != nil {
			return c.JSON(400, newResponse[any](nil, err.Error()))
		}
		return c.JSON(200, updated_class)

	case "subject":
		var data UpdateClassroomSubject

		bindErr := c.Bind(&data)
		if bindErr != nil {
			return c.JSON(400, newResponse[any](nil, bindErr.Error()))
		}
		if err := c.Validate(data); err != nil {
			return c.JSON(400, newResponse[any](nil, err.Error()))
		}

		updated_class, err := s.store.UpdateClassroomSubject(c.Request().Context(), database.UpdateClassroomSubjectParams{
			ID:      uid,
			Subject: data.Subject,
		})

		if err != nil {
			return c.JSON(400, newResponse[any](nil, err.Error()))
		}
		return c.JSON(200, updated_class)

	case "section":
		var data UpdateClassroomSection

		bindErr := c.Bind(&data)
		if bindErr != nil {
			return c.JSON(400, newResponse[any](nil, bindErr.Error()))
		}
		if err := c.Validate(data); err != nil {
			return c.JSON(400, newResponse[any](nil, err.Error()))
		}

		updated_class, err := s.store.UpdateClassroomSection(c.Request().Context(), database.UpdateClassroomSectionParams{
			ID:      uid,
			Section: data.Section,
		})

		if err != nil {
			return c.JSON(400, newResponse[any](nil, err.Error()))
		}
		return c.JSON(200, updated_class)

	case "room":
		var data UpdateClassroomRoom

		bindErr := c.Bind(&data)
		if bindErr != nil {
			return c.JSON(400, newResponse[any](nil, bindErr.Error()))
		}
		if err := c.Validate(data); err != nil {
			return c.JSON(400, newResponse[any](nil, err.Error()))
		}

		updated_class, err := s.store.UpdateClassroomRoom(c.Request().Context(), database.UpdateClassroomRoomParams{
			ID:   uid,
			Room: data.Room,
		})

		if err != nil {
			return c.JSON(400, newResponse[any](nil, err.Error()))
		}
		return c.JSON(200, updated_class)

	case "invite_code":
		updated_class, err := s.store.UpdateClassroomInviteCode(c.Request().Context(), database.UpdateClassroomInviteCodeParams{
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
		return c.JSON(http.StatusBadRequest, newResponse[any](nil, err.Error()))
	}

	payload := c.Get("user").(*token.Payload)

	classroom, err := s.store.GetClass(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, newResponse[any](nil, err.Error()))
	}

	if payload.ID != classroom.AdminID {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	deleted_classroom, err := s.store.DeleteClass(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, newResponse[any](nil, err.Error()))
	}

	return c.JSON(http.StatusOK, newResponse(deleted_classroom, ""))
}

func (s *Server) getClassroomUsers(c echo.Context) error {
	id := c.Param("id")
	page := c.QueryParam("page")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)

	if err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	if offset < 0 {
		return c.JSON(400, NEGATIVE_OFFSET)
	}

	members, err := s.store.GetAllClassroomMembers(c.Request().Context(), database.GetAllClassroomMembersParams{
		ClassID: uid,
		Offset:  int32(offset * 10),
	})

	if err != nil {
		return c.JSON(400, CLASSROOM_NOTFOUND)
	}
	return c.JSON(200, newResponse(members, ""))
}

func (s *Server) joinClassroom(c echo.Context) error {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	check_user, err := s.store.GetUser(c.Request().Context(), uid)
	if err == sql.ErrNoRows && check_user.ID == uuid.Nil {
		return c.JSON(400, USER_NOTFOUND)
	}

	var req ClassroomJoinRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &req); err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	classroom_id, err := s.store.GetClassroomWithInviteCode(c.Request().Context(), req.InviteCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(400, CLASSROOM_NOTFOUND)
		}
		return c.JSON(http.StatusInternalServerError, newResponse[any](nil, err.Error()))
	}

	payload := c.Get("user").(*token.Payload)

	if payload.ID != uid {
		return c.JSON(400, UNAUTHORIZED)
	}

	joined, err := s.store.AddNewClassroomMember(c.Request().Context(), database.AddNewClassroomMemberParams{
		ID:      uuid.New(),
		ClassID: classroom_id,
		UserID:  uid,
	})
	if err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}
	return c.JSON(200, joined)
}

func (s *Server) leaveClassroom(c echo.Context) error {
	id := c.Param("id")
	class_id := c.Param("class_id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	parsed_classId, err := uuid.Parse(class_id)
	if err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	check_user, err := s.store.GetUser(c.Request().Context(), uid)
	if err == sql.ErrNoRows && check_user.ID == uuid.Nil {
		return c.JSON(400, USER_NOTFOUND)
	}

	check_class, err := s.store.GetClass(c.Request().Context(), parsed_classId)
	if err == sql.ErrNoRows && check_class.ID == uuid.Nil {
		return c.JSON(400, CLASSROOM_NOTFOUND)
	}

	payload := c.Get("user").(*token.Payload)

	if payload.ID != uid {
		return c.JSON(400, UNAUTHORIZED)
	}

	leaved, err := s.store.LeaveClassroom(c.Request().Context(), database.LeaveClassroomParams{
		UserID:  payload.ID,
		ClassID: parsed_classId,
	})
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, newResponse(leaved, ""))
}

func (s *Server) getClassroomPosts(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)

	if err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	if offset < 0 {
		return c.JSON(400, NEGATIVE_OFFSET)
	}

	posts, err := s.store.ListAllPostsFromClass(c.Request().Context(), database.ListAllPostsFromClassParams{
		ClassID: uid,
		Offset:  int32(offset * 10),
	})

	if err != nil {
		return c.JSON(400, newResponse[any](nil, err.Error()))
	}

	return c.JSON(200, newResponse(posts, ""))
}
