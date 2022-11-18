package routes

import (
	database "server/database/sqlc"
	"server/token"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	CreatePostRequest struct {
		Content string `json:"content" validate:"required,gt=0"`
		ClassID string `json:"class_id" validate:"required,uuid"`
	}

	UpdatePostRequest struct {
		Content string `json:"content" validate:"required,gt=0"`
	}

	LikeRequest struct {
		ClassID uuid.UUID `json:"class_id" validate:"required,uuid"`
	}

	UnlikeRequest struct {
		ID uuid.UUID `json:"id" validate:"required,uuid"`
	}
)

func (s *Server) getOnePost(c echo.Context) error {
	id := c.Param("id")
	pid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, err)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := token.GetPayloadFromJwt(jwt_token)

	post, err := s.store.GetOnePost(c.Request().Context(), pid)
	if err != nil || post.ID == uuid.Nil {
		return c.JSON(404, POST_NOTFOUND)
	}

	check_user_if_member, err := s.store.GetClassroomMemberById(c.Request().Context(), database.GetClassroomMemberByIdParams{
		UserID:  jwt_payload.ID,
		ClassID: post.ClassID,
	})
	if err != nil || check_user_if_member.ID == uuid.Nil {
		return c.JSON(404, NOT_A_MEMBER)
	}

	return c.JSON(200, post)
}

func (s *Server) createNewPost(c echo.Context) error {
	var payload CreatePostRequest

	bindErr := c.Bind(&payload)
	if bindErr != nil {
		return c.JSON(400, bindErr)
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err)
	}

	parsed_classId, err := uuid.Parse(payload.ClassID)
	if err != nil {
		return c.JSON(400, err)
	}

	check_class, err := s.store.GetClass(c.Request().Context(), parsed_classId)
	if err != nil || check_class.ID == uuid.Nil {
		return c.JSON(404, CLASSROOM_NOTFOUND)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := token.GetPayloadFromJwt(jwt_token)

	check_user_if_member, err := s.store.GetClassroomMemberById(c.Request().Context(), database.GetClassroomMemberByIdParams{
		UserID:  jwt_payload.ID,
		ClassID: parsed_classId,
	})
	if err != nil || check_user_if_member.ID == uuid.Nil {
		return c.JSON(404, NOT_A_MEMBER)
	}

	new_postParam := database.InsertNewPostParams{
		ID:       uuid.New(),
		Content:  payload.Content,
		AuthorID: jwt_payload.ID,
		ClassID:  parsed_classId,
	}

	new_post, err := s.store.InsertNewPost(c.Request().Context(), new_postParam)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, new_post)
}

func (s *Server) updatePost(c echo.Context) error {
	postId := c.Param("id")
	puid, err := uuid.Parse(postId)
	if err != nil {
		return c.JSON(400, err)
	}

	var payload UpdatePostRequest

	bindErr := c.Bind(&payload)
	if bindErr != nil {
		return c.JSON(400, bindErr)
	}
	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := token.GetPayloadFromJwt(jwt_token)

	check_post, err := s.store.GetOnePost(c.Request().Context(), puid)
	if err != nil || check_post.ID == uuid.Nil {
		return c.JSON(400, POST_NOTFOUND)
	}

	if jwt_payload.ID != check_post.AuthorID {
		return c.JSON(401, UNAUTHORIZED)
	}

	updated_post, err := s.store.UpdatePostContent(c.Request().Context(), database.UpdatePostContentParams{
		Content: payload.Content,
	})

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, updated_post)
}

func (s *Server) deletePost(c echo.Context) error {
	post_id := c.Param("id")
	puid, err := uuid.Parse(post_id)
	if err != nil {
		return c.JSON(400, err)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := token.GetPayloadFromJwt(jwt_token)

	check_post, err := s.store.GetOnePost(c.Request().Context(), puid)
	if err != nil || check_post.ID == uuid.Nil {
		return c.JSON(400, POST_NOTFOUND)
	}

	if jwt_payload.ID != check_post.AuthorID {
		return c.JSON(401, UNAUTHORIZED)
	}

	deleted_post, err := s.store.DeletePostFromClass(c.Request().Context(), database.DeletePostFromClassParams{
		ClassID:  check_post.ClassID,
		AuthorID: check_post.AuthorID,
		ID:       check_post.ID,
	})
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, deleted_post)
}

func (s *Server) getAllPostLikes(c echo.Context) error {
	post_id := c.Param("id")

	post_uuid, err := uuid.Parse(post_id)
	if err != nil {
		return c.JSON(400, err)
	}

	likes, err := s.store.GetAllPostLikes(c.Request().Context(), post_uuid)
	if err != nil {
		return c.JSON(400, err)
	}
	// todo
	return c.JSON(200, likes)
}

func (s *Server) likePost(c echo.Context) error {
	post_id := c.Param("id")

	post_uuid, err := uuid.Parse(post_id)
	if err != nil {
		return c.JSON(400, err)
	}

	var payload LikeRequest
	bindErr := c.Bind(&payload)
	if bindErr != nil {
		return c.JSON(400, bindErr)
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := token.GetPayloadFromJwt(jwt_token)

	check_member, err := s.store.GetClassroomMemberById(c.Request().Context(), database.GetClassroomMemberByIdParams{
		UserID:  jwt_payload.ID,
		ClassID: payload.ClassID,
	})
	if err != nil || check_member.ID == uuid.Nil {
		return c.JSON(401, NOT_A_MEMBER)
	}

	liked_post, err := s.store.LikePost(c.Request().Context(), database.LikePostParams{
		ID:     uuid.New(),
		UserID: jwt_payload.ID,
		PostID: post_uuid,
	})
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, liked_post)
}

func (s *Server) unlikePost(c echo.Context) error {
	post_id := c.Param("id")
	post_uid, err := uuid.Parse(post_id)

	if err != nil {
		return c.JSON(400, err)
	}

	var payload UnlikeRequest
	bindErr := c.Bind(&payload)
	if bindErr != nil {
		return c.JSON(400, bindErr)
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	jwt_payload := token.GetPayloadFromJwt(jwt_token)

	unliked_post, err := s.store.UnlikePost(c.Request().Context(), database.UnlikePostParams{
		ID:     post_uid,
		PostID: payload.ID,
		UserID: jwt_payload.ID,
	})
	if err != nil || unliked_post.ID == uuid.Nil {
		return c.JSON(400, POST_NOTFOUND)
	}

	return c.JSON(200, unliked_post)
}
