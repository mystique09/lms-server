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
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserUpdateDTO struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

type UserClassrooms struct {
	*database.User
	Rooms []database.Classroom `json:"classrooms"`
}

func (s *Server) getUsers(c echo.Context) error {
	page := c.QueryParam("page")
	comment := c.QueryParam("comment_page")

	if page == "" {
		page = "0"
	}

	if comment == "" {
		comment = "0"
	}

	offset, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewResponse(nil, err.Error()))
	}

	comment_offset, err := strconv.Atoi(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewResponse(nil, err.Error()))
	}

	users, err := s.DB.GetUsers(c.Request().Context(), int32(offset*10))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewResponse(nil, err.Error()))
	}

	var users_wclassrooms []UserClassrooms = []UserClassrooms{}

	for i := range users {
		user_resp := UserClassrooms{
			User:  &users[i],
			Rooms: []database.Classroom{},
		}

		classrooms, err := s.DB.GetAllJoinedClassrooms(c.Request().Context(), database.GetAllJoinedClassroomsParams{
			UserID: users[i].ID,
			Offset: int32(comment_offset * 10),
		})

		if err != nil {
			user_resp.Rooms = []database.Classroom{}
		}
		user_resp.Rooms = append(user_resp.Rooms, classrooms...)
		users_wclassrooms = append(users_wclassrooms, user_resp)
	}

	return c.JSON(http.StatusOK, utils.NewResponse(&users_wclassrooms, ""))
}

func (s *Server) getUser(c echo.Context) error {
	id := c.Param("id")
	page := c.QueryParam("page")

	if page == "" {
		page = "0"
	}

	offset, err := strconv.Atoi(page)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Please provide an ID."))
	}

	user, err := s.DB.GetUser(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	classrooms, err := s.DB.GetAllJoinedClassrooms(c.Request().Context(), database.GetAllJoinedClassroomsParams{
		UserID: uid,
		Offset: int32(offset * 10),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	user_wclassrooms := UserClassrooms{
		User:  &user,
		Rooms: classrooms,
	}

	return c.JSON(http.StatusOK, utils.NewResponse(user_wclassrooms, ""))
}

func (s *Server) createUser(c echo.Context) error {
	user_data := new(UserCreateDTO)
	if err := c.Bind(user_data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	if user_data.Username == "" || user_data.Email == "" || user_data.Password == "" {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Please provide all inputs."))
	}

	if len(user_data.Username) < 8 {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Username must be at least 8 characters."))
	}

	if len(user_data.Password) < 8 {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Password must be at least 8 characters."))
	}

	if !utils.IsEmail(user_data.Email) {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Please provide a valid email."))
	}

	check_user, err := s.DB.GetUserByUsername(c.Request().Context(), user_data.Username)

	if check_user.ID != uuid.Nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "User already exist."))
	}

	hashed_password, err := utils.Encrypt(user_data.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewResponse(nil, err.Error()))
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
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(user, ""))
}

func (s *Server) updateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Please provide an ID."))
	}

	field := c.QueryParam("field")

	if field == "" {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Invalid query field, idk what to update."))
	}

	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	var updateDto UserUpdateDTO
	if err := (&echo.DefaultBinder{}).BindBody(c, &updateDto); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	// check if the current user is the one being updated
	token := c.Get("user").(*jwt.Token)
	var payload utils.JwtUserPayload = utils.GetPayloadFromJwt(token)

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)

	if err != nil || check_user.ID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "User not found."))
	}

	if check_user.Username != payload.Username || check_user.Email != payload.Email {
		return c.JSON(http.StatusUnauthorized, "You don't have the permission to update this user.")
	}

	if field == "username" {
		payload := database.UpdateUsernameParams{
			ID:       uid,
			Username: updateDto.Username,
		}

		if updateDto.Username == "" {
			return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Empty required field."))
		}

		new_user, err := s.DB.UpdateUsername(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, utils.NewResponse(new_user, ""))
	} else if field == "email" {
		payload := database.UpdateUserEmailParams{
			ID:    uid,
			Email: updateDto.Email,
		}

		if updateDto.Email == "" {
			return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Empty required field."))
		}

		new_user, err := s.DB.UpdateUserEmail(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, utils.NewResponse(new_user, ""))
	} else if field == "password" {
		hashed_password, err := utils.Encrypt(updateDto.Password)

		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
		}

		payload := database.UpdateUserPasswordParams{
			ID:       uid,
			Password: hashed_password,
		}

		if updateDto.Password == "" {
			return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Empty required field."))
		}

		new_user, err := s.DB.UpdateUserPassword(c.Request().Context(), payload)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, utils.NewResponse(new_user, ""))
	}

	return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Invalid query field, idk what to update."))
}

func (s *Server) deleteUser(c echo.Context) error {
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "Please provide an ID."))
	}

	jwt_token := c.Get("user").(*jwt.Token)
	var payload utils.JwtUserPayload = utils.GetPayloadFromJwt(jwt_token)

	check_user, err := s.DB.GetUser(c.Request().Context(), uid)

	if err != nil || check_user.ID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, "User not found."))
	}

	if check_user.Username != payload.Username || check_user.Email != payload.Email {
		return c.JSON(http.StatusUnauthorized, "You don't have the permission to delete this user.")
	}

	deleted_user, err := s.DB.DeleteUser(c.Request().Context(), uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(deleted_user, ""))
}
