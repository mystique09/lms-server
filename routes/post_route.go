package routes

import (
	"fmt"
	database "server/database/sqlc"
	"server/utils"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	CreatePostRequest struct {
		Content string `json:"content"`
		ClassID string `json:"class_id"`
	}

	UpdatePostRequest struct {
		Content string `json:"content"`
	}
)

func (s *Server) getOnePost(c echo.Context) error {
	id := c.Param("id")
	pid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	post, err := s.DB.GetOnePost(c.Request().Context(), pid)
	if err != nil || post.ID == uuid.Nil {
		return c.JSON(404, utils.NewResponse(nil, fmt.Sprintf("post with id [%v] doesn't exist", pid)))
	}

	check_user_if_member, err := s.DB.GetClassroomMemberById(c.Request().Context(), database.GetClassroomMemberByIdParams{
		UserID:  jwt_payload.ID,
		ClassID: post.ClassID,
	})
	if err != nil || check_user_if_member.ID == uuid.Nil {
		return c.JSON(404, utils.NewResponse(nil, fmt.Sprintf("user with id [%v] is not a member of classroom with id [%v]", jwt_payload.ID, post.ClassID)))
	}

	return c.JSON(200, utils.NewResponse(post, ""))
}

func (s *Server) createNewPost(c echo.Context) error {
	var payload CreatePostRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	parsed_classId, err := uuid.Parse(payload.ClassID)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	check_class, err := s.DB.GetClass(c.Request().Context(), parsed_classId)
	if err != nil || check_class.ID == uuid.Nil {
		return c.JSON(404, utils.NewResponse("", fmt.Sprintf("classroom with id [%v] doesn't exist", parsed_classId)))
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	check_user_if_member, err := s.DB.GetClassroomMemberById(c.Request().Context(), database.GetClassroomMemberByIdParams{
		UserID:  jwt_payload.ID,
		ClassID: parsed_classId,
	})
	if err != nil || check_user_if_member.ID == uuid.Nil {
		return c.JSON(404, utils.NewResponse(nil, fmt.Sprintf("user with id [%v] is not a member of classroom with id [%v]", jwt_payload.ID, parsed_classId)))
	}

	new_postParam := database.InsertNewPostParams{
		ID:       uuid.New(),
		Content:  payload.Content,
		AuthorID: jwt_payload.ID,
		ClassID:  parsed_classId,
	}

	new_post, err := s.DB.InsertNewPost(c.Request().Context(), new_postParam)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(200, utils.NewResponse(new_post, ""))
}

func (s *Server) updatePost(c echo.Context) error {
	postId := c.Param("id")
	puid, err := uuid.Parse(postId)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	var payload UpdatePostRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	check_post, err := s.DB.GetOnePost(c.Request().Context(), puid)
	if err != nil || check_post.ID == uuid.Nil {
		return c.JSON(400, utils.NewResponse(nil, fmt.Sprintf("post with id [%v] doesn't exist", puid)))
	}

	if jwt_payload.ID != check_post.AuthorID {
		return c.JSON(401, utils.NewResponse(nil, "you can't perform this action"))
	}

	updated_post, err := s.DB.UpdatePostContent(c.Request().Context(), database.UpdatePostContentParams{
		Content: payload.Content,
	})

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(200, utils.NewResponse(updated_post, ""))
}

func (s *Server) deletePost(c echo.Context) error {
	post_id := c.Param("id")
	puid, err := uuid.Parse(post_id)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := utils.GetPayloadFromJwt(jwt_token)

	check_post, err := s.DB.GetOnePost(c.Request().Context(), puid)
	if err != nil || check_post.ID == uuid.Nil {
		return c.JSON(400, utils.NewResponse(nil, fmt.Sprintf("post with id [%v] doesn't exist", puid)))
	}

	if jwt_payload.ID != check_post.AuthorID {
		return c.JSON(401, utils.NewResponse(nil, "you can't perform this action"))
	}

	deleted_post, err := s.DB.DeletePostFromClass(c.Request().Context(), database.DeletePostFromClassParams{
		ClassID:  check_post.ClassID,
		AuthorID: check_post.AuthorID,
		ID:       check_post.ID,
	})
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(200, utils.NewResponse(deleted_post, ""))
}

func (s *Server) getAllPostLikes(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) likePost(c echo.Context) error {
	return c.String(200, "TODO")
}

func (s *Server) unlikePost(c echo.Context) error {
	return c.String(200, "TODO")
}
