package routes

import (
	"net/http"
	database "server/database/sqlc"
	"server/utils"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	UserCreateDTO struct {
		Username string `json:"username" validate:"required,gt=6"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gt=6"`
	}

	UserUpdateDTO struct {
		Username string `json:"username" validate:"required,gt=6"`
		Email    string `json:"email" validate:"required, email"`
		Password string `json:"password" validate:"required,gt-6"`
	}

	UserResponse struct {
		ID         uuid.UUID           `json:"id"`
		Username   string              `json:"username"`
		Email      string              `json:"email"`
		UserRole   database.Role       `json:"user_role"`
		Visibility database.Visibility `json:"visibility"`
	}
)

func newUserResponse(user database.User) UserResponse {
	return UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		UserRole:   user.UserRole,
		Visibility: user.Visibility,
	}
}

type UserClassrooms struct {
	*database.User
	Rooms []database.Classroom `json:"classrooms"`
}

func (s *Server) getUsers(c echo.Context) error {
	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	users, err := s.DB.GetUsers(c.Request().Context(), int32(offset*10))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.Render(200, "userPage", users)
	//return c.JSON(http.StatusOK, users)
}

func (s *Server) getUser(c echo.Context) error {
	id := c.Param("id")
	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, MISSING_ID_FIELD)
	}

	user, err := s.DB.GetUser(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	classrooms, err := s.DB.GetAllJoinedClassrooms(c.Request().Context(), database.GetAllJoinedClassroomsParams{
		UserID: uid,
		Offset: int32(offset * 10),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user_wclassrooms := UserClassrooms{
		User:  &user,
		Rooms: classrooms,
	}

	return c.JSON(http.StatusOK, user_wclassrooms)
}

func (s *Server) createUser(c echo.Context) error {
	user_data := new(UserCreateDTO)
	c.Bind(&user_data)

	if err := c.Validate(user_data); err != nil {
		return c.JSON(400, err)
	}

	check_user, err := s.DB.GetUserByUsername(c.Request().Context(), user_data.Username)

	if check_user.ID != uuid.Nil {
		return c.JSON(http.StatusBadRequest, USER_ALREADY_EXIST)
	}

	hashed_password, err := utils.Encrypt(user_data.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var new_user_param database.CreateUserParams = database.CreateUserParams{
		ID:         uuid.New(),
		Username:   user_data.Username,
		Email:      user_data.Email,
		Password:   hashed_password,
		UserRole:   database.RoleSTUDENT,
		Visibility: database.VisibilityPUBLIC,
	}

	user, err := s.DB.CreateUser(c.Request().Context(), new_user_param)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) updateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, MISSING_ID_FIELD)
	}

	field := c.QueryParam("field")

	if field == "" {
		return c.JSON(http.StatusBadRequest, EMPTY_QUERY_PARAM)
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var updateDto UserUpdateDTO
	c.Bind(&updateDto)
	if err := c.Validate(updateDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// check if the current user is the one being updated
	token := c.Get("user").(*jwt.Token)
	var payload utils.JwtUserPayload = utils.GetPayloadFromJwt(token)

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)

	if err != nil || check_user.ID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, USER_NOTFOUND)
	}

	if check_user.Username != payload.Username || check_user.Email != payload.Email {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	if field == "username" {
		payload := database.UpdateUsernameParams{
			ID:       uid,
			Username: updateDto.Username,
		}

		if updateDto.Username == "" {
			return c.JSON(http.StatusBadRequest, MISSING_FIELDS)
		}

		new_user, err := s.DB.UpdateUsername(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, new_user)
	} else if field == "email" {
		payload := database.UpdateUserEmailParams{
			ID:    uid,
			Email: updateDto.Email,
		}

		if updateDto.Email == "" {
			return c.JSON(http.StatusBadRequest, MISSING_FIELDS)
		}

		new_user, err := s.DB.UpdateUserEmail(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, new_user)
	} else if field == "password" {
		hashed_password, err := utils.Encrypt(updateDto.Password)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		payload := database.UpdateUserPasswordParams{
			ID:       uid,
			Password: hashed_password,
		}

		if updateDto.Password == "" {
			return c.JSON(http.StatusBadRequest, MISSING_FIELDS)
		}

		new_user, err := s.DB.UpdateUserPassword(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, new_user)
	}

	return c.JSON(http.StatusBadRequest, UNKNOWN_FIELD)
}

func (s *Server) deleteUser(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, MISSING_ID_FIELD)
	}

	jwt_token := c.Get("user").(*jwt.Token)
	var payload utils.JwtUserPayload = utils.GetPayloadFromJwt(jwt_token)

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)

	if err != nil || check_user.ID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, USER_NOTFOUND)
	}

	if check_user.Username != payload.Username || check_user.Email != payload.Email {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	deleted_user, err := s.DB.DeleteUser(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, deleted_user)
}
