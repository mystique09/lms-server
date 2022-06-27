package routes

import (
	"context"
	"net/http"
	database "server/database/sqlc"
	"server/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserCreateDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
A function to retrive all users.
*/
func (rt *Route) GetUsers(c echo.Context) error {
	ctx := context.Background()
	users, err := rt.DB.GetUsers(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewResponse(0, nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(1, &users, ""))
}

/*
A function to retrive a user by ID.
*/
func (rt *Route) GetUser(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, err.Error()))
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, "Please provide an ID.")
	}

	user, err := rt.DB.GetUser(ctx, uid)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(1, user, ""))
}

/*
A function to create new user.
*/
func (rt *Route) CreateUser(c echo.Context) error {
	ctx := context.Background()
	user_data := new(UserCreateDTO)
	if err := c.Bind(user_data); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, err.Error()))
	}

	// check if inputs is not empty
	if user_data.Username == "" || user_data.Email == "" || user_data.Password == "" {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, "Please provide all inputs."))
	}

	// check username length if greater than 8
	if len(user_data.Username) < 8 {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, "Username must be at least 8 characters."))
	}

	// check password length if greater than 8
	if len(user_data.Password) < 8 {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, "Password must be at least 8 characters."))
	}

	// check if email is valid
	if !utils.IsEmail(user_data.Email) {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, "Please provide a valid email."))
	}

	// check if user already exist
	check_user, err := rt.DB.GetUserByUsername(ctx, user_data.Username)

	if check_user.ID != uuid.Nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, "User already exist."))
	}

	var new_user_param database.CreateUserParams = database.CreateUserParams{
		ID:       uuid.New(),
		Username: user_data.Username,
		Email:    user_data.Email,
		Password: user_data.Password,
		UserRole: database.RoleSTUDENT,
	}

	user, err := rt.DB.CreateUser(ctx, new_user_param)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(1, user, ""))
}

/*
A function to update a user by ID.
*/
func (rt *Route) UpdateUser(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, err.Error()))
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, "Please provide an ID."))
	}

	user := new(database.UpdateUserParams)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, err.Error()))
	}

	user.ID = uid

	if err := rt.DB.UpdateUser(ctx, *user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(1, user, ""))
}

/*
A function to delete a user by ID.
*/
func (rt *Route) DeleteUser(c echo.Context) error {
	ctx := context.Background()
	id := c.Param("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, err.Error()))
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, "Please provide an ID."))
	}

	if err := rt.DB.DeleteUser(ctx, uid); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(0, nil, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(1, id, ""))
}
