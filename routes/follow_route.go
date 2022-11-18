package routes

import (
	"fmt"
	database "server/database/sqlc"
	"server/utils"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserFollowers struct {
	UserId    uuid.UUID                     `json:"user_id"`
	Followers []database.GetAllFollowersRow `json:"followers"`
}

type UserFollowing struct {
	UserId    uuid.UUID                     `json:"user_id"`
	Following []database.GetAllFollowingRow `json:"followings"`
}

type FollowUserRequest struct {
	UserId uuid.UUID `json:"user_id" validate:"required,uuid"`
}

func (s *Server) getFollowers(c echo.Context) error {
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

	if err != nil {
		return c.JSON(400, err)
	}

	if offset < 0 {
		return c.JSON(400, NEGATIVE_OFFSET)
	}

	followers, err := s.store.GetAllFollowers(c.Request().Context(), database.GetAllFollowersParams{
		Following: uid,
		Offset:    int32(offset * 10),
	})

	if err != nil {
		return c.JSON(400, err)
	}

	user_followers := UserFollowers{
		UserId:    uid,
		Followers: followers,
	}

	return c.JSON(200, user_followers)
}

func (s *Server) getFollowings(c echo.Context) error {
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

	if err != nil {
		return c.JSON(400, err)
	}

	if offset < 0 {
		return c.JSON(400, NEGATIVE_OFFSET)
	}

	following, err := s.store.GetAllFollowing(c.Request().Context(), database.GetAllFollowingParams{
		Follower: uid,
		Offset:   int32(offset * 10),
	})

	if err != nil {
		return c.JSON(400, err)
	}

	user_following := UserFollowing{
		UserId:    uid,
		Following: following,
	}

	return c.JSON(200, user_following)
}

func (s *Server) addNewFollower(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(400, err)
	}

	token := c.Get("user").(*jwt.Token)
	user := utils.GetPayloadFromJwt(token)

	var payload FollowUserRequest

	bindErr := c.Bind(&payload)
	if bindErr != nil {
		return c.JSON(400, bindErr)
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(400, err)
	}

	check_userid, err := s.store.GetUser(c.Request().Context(), payload.UserId)
	if err != nil || check_userid.ID == uuid.Nil {
		return c.JSON(400, USER_NOTFOUND)
	}

	if payload.UserId == uid {
		return c.JSON(400, newResponse[any](nil, "you can't follow yourself"))
	}

	if user.ID != uid {
		return c.JSON(400, UNAUTHORIZED)
	}

	check_follow, err := s.store.GetOneFollower(c.Request().Context(), database.GetOneFollowerParams{
		Follower:  user.ID,
		Following: payload.UserId,
	})

	if err == nil || check_follow.ID != uuid.Nil {
		return c.JSON(400, newResponse[any](nil, "you already followed this user"))
	}

	new_follower, err := s.store.FollowUser(c.Request().Context(), database.FollowUserParams{
		ID:        uuid.New(),
		Follower:  user.ID,
		Following: payload.UserId,
	})

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, new_follower)
}

func (s *Server) removeFollowing(c echo.Context) error {
	id := c.Param("id")
	following_id := c.Param("following_id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, err)
	}

	follow_id, err := uuid.Parse(following_id)
	if err != nil {
		return c.JSON(400, err)
	}

	check_user, err := s.store.GetUser(c.Request().Context(), uid)

	if check_user.ID == uuid.Nil || err != nil {
		return c.JSON(400, USER_NOTFOUND)
	}

	check_following, err := s.store.GetFollowerById(c.Request().Context(), follow_id)

	if check_following.ID == uuid.Nil || err != nil {
		return c.JSON(400, newResponse[any](nil, fmt.Sprintf("[%v] is not in your followings list!", follow_id)))
	}

	token := c.Get("user").(*jwt.Token)
	user := utils.GetPayloadFromJwt(token)

	if user.ID != uid {
		return c.JSON(403, UNAUTHORIZED)
	}

	unfollowed, err := s.store.UnfollowUser(c.Request().Context(), follow_id)

	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, unfollowed)
}
