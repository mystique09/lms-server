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
	UserId uuid.UUID `json:"user_id"`
}

func (s *Server) getFollowers(c echo.Context) error {
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

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	if offset < 0 {
		return c.JSON(400, utils.NewResponse(nil, "offset must not be negative"))
	}

	followers, err := s.DB.GetAllFollowers(c.Request().Context(), database.GetAllFollowersParams{
		Following: uid,
		Offset:    int32(offset * 10),
	})

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	user_followers := UserFollowers{
		UserId:    uid,
		Followers: followers,
	}

	return c.JSON(200, utils.NewResponse(user_followers, ""))
}

func (s *Server) getFollowings(c echo.Context) error {
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

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	if offset < 0 {
		return c.JSON(400, utils.NewResponse(nil, "offset must not be negative"))
	}

	following, err := s.DB.GetAllFollowing(c.Request().Context(), database.GetAllFollowingParams{
		Follower: uid,
		Offset:   int32(offset * 10),
	})

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	user_following := UserFollowing{
		UserId:    uid,
		Following: following,
	}

	return c.JSON(200, utils.NewResponse(user_following, ""))
}

func (s *Server) addNewFollower(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	fmt.Println(id, uid)

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	token := c.Get("user").(*jwt.Token)
	user := utils.GetPayloadFromJwt(token)

	var payload FollowUserRequest
	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	check_userid, err := s.DB.GetUser(c.Request().Context(), payload.UserId)
	fmt.Println(check_userid.ID)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, fmt.Sprintf("user %v doesn't exist", payload.UserId)))
	}

	if payload.UserId == uid {
		return c.JSON(400, utils.NewResponse(nil, "you can't follow yourself"))
	}

	if user.ID != uid {
		return c.JSON(400, utils.NewResponse(nil, "you are not authorized to perform this action"))
	}

	check_follow, err := s.DB.GetOneFollower(c.Request().Context(), database.GetOneFollowerParams{
		Follower:  user.ID,
		Following: payload.UserId,
	})

	if err == nil || check_follow.ID != uuid.Nil {
		return c.JSON(400, utils.NewResponse(nil, "you already followed this user"))
	}

	new_follower, err := s.DB.FollowUser(c.Request().Context(), database.FollowUserParams{
		ID:        uuid.New(),
		Follower:  user.ID,
		Following: payload.UserId,
	})

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(200, utils.NewResponse(new_follower, ""))
}

func (s *Server) removeFollowing(c echo.Context) error {
	id := c.Param("id")
	following_id := c.Param("following_id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	follow_id, err := uuid.Parse(following_id)
	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)

	if check_user.ID == uuid.Nil || err != nil {
		return c.JSON(400, utils.NewResponse(nil, fmt.Sprintf("user [%v] doesn't exist.", uid)))
	}

	check_following, err := s.DB.GetFollowerById(c.Request().Context(), follow_id)

	if check_following.ID == uuid.Nil || err != nil {
		return c.JSON(400, utils.NewResponse(nil, fmt.Sprintf("[%v] is not in your followings list!", follow_id)))
	}

	// if check_user.ID == uuid.Nil {
	// 	return c.JSON(400, utils.NewResponse(nil, "are you sure the id you provided is correct?"))
	// }

	token := c.Get("user").(*jwt.Token)
	user := utils.GetPayloadFromJwt(token)

	if user.ID != uid {
		return c.JSON(403, utils.NewResponse(nil, "you are not authorize to perform this action"))
	}

	unfollowed, err := s.DB.UnfollowUser(c.Request().Context(), follow_id)

	if err != nil {
		return c.JSON(400, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(200, utils.NewResponse(unfollowed, ""))
}
