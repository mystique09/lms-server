package routes

import (
	"database/sql"
	"net/http"
	database "server/database/sqlc"
	"server/token"
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

type UserClassrooms struct {
	*database.User
	Rooms []database.Classroom `json:"classrooms"`
}

func (s *Server) getUsers(c echo.Context) error {
	ofst := c.QueryParam("offset")

	if ofst == "" || ofst == "0" {
		ofst = "0"
	}

	offset, err := strconv.Atoi(ofst)

	if err != nil || offset < 0 {
		return c.JSON(400, "Invalid page, must be a number!")
	}

	users, err := s.store.GetUsers(c.Request().Context(), int32(offset*10))

	if err != nil {
		return c.JSON(http.StatusBadRequest, newError(err.Error()))
	}

	return c.JSON(200, newResponse(users))
}

func (s *Server) getUser(c echo.Context) error {
	id := c.Param("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, newError(err.Error()))
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, MISSING_ID_FIELD)
	}

	user, err := s.store.GetUser(c.Request().Context(), uid)
	user.Password = ""

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, USER_NOTFOUND)
	}

	if err == sql.ErrConnDone {
		return c.JSON(http.StatusInternalServerError, newError(err.Error()))
	}

	return c.JSON(http.StatusOK, newResponse(user))
}

func (s *Server) createUser(c echo.Context) error {
	user_data := new(UserCreateDTO)

	bindErr := c.Bind(&user_data)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, "Missing required field.")
	}

	if err := c.Validate(user_data); err != nil {
		return c.JSON(400, newError(err.Error()))
	}

	hashed_password, err := utils.Encrypt(user_data.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newError(err.Error()))
	}

	var new_user_param database.CreateUserParams = database.CreateUserParams{
		ID:         uuid.New(),
		Username:   user_data.Username,
		Email:      user_data.Email,
		Password:   hashed_password,
		UserRole:   database.RoleSTUDENT,
		Visibility: database.VisibilityPUBLIC,
	}

	user, err := s.store.CreateUser(c.Request().Context(), new_user_param)

	if err != nil {
		return c.JSON(http.StatusBadRequest, USER_ALREADY_EXIST)
	}

	return c.JSON(http.StatusOK, newResponse(user))
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

	bindErr := c.Bind(&updateDto)
	if bindErr != nil {
		return c.JSON(400, bindErr)
	}

	if err := c.Validate(updateDto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// check if the current user is the one being updated
	payload := c.Get("user").(*token.Payload)

	check_user, err := s.store.GetUser(c.Request().Context(), uid)

	if err != nil || check_user.ID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, USER_NOTFOUND)
	}

	if check_user.Username != payload.Username || check_user.ID != payload.ID {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	switch field {
	case "username":
		payload := database.UpdateUsernameParams{
			ID:       uid,
			Username: updateDto.Username,
		}

		if updateDto.Username == "" {
			return c.JSON(http.StatusBadRequest, MISSING_FIELDS)
		}

		new_user, err := s.store.UpdateUsername(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, new_user)
	case "email":
		payload := database.UpdateUserEmailParams{
			ID:    uid,
			Email: updateDto.Email,
		}

		if updateDto.Email == "" {
			return c.JSON(http.StatusBadRequest, MISSING_FIELDS)
		}

		new_user, err := s.store.UpdateUserEmail(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, new_user)
	case "password":
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

		new_user, err := s.store.UpdateUserPassword(c.Request().Context(), payload)
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
	payload := token.GetPayloadFromJwt(jwt_token)

	check_user, err := s.store.GetUser(c.Request().Context(), uid)

	if err != nil || check_user.ID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, USER_NOTFOUND)
	}

	if check_user.Username != payload.Username || check_user.Email != payload.Email {
		return c.JSON(http.StatusUnauthorized, UNAUTHORIZED)
	}

	deleted_user, err := s.store.DeleteUser(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, deleted_user)
}
