package domain

import "github.com/labstack/echo/v4"

type SignupRequest struct {
	Username string `json:"username" validate:"required,gte=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SignupResponse struct {
	User NewUser `json:"user"`
}

type SignupUsecase interface {
	GetUserByUsername(c echo.Context, username string) (User, error)
	CreateUser(c echo.Context, username, email, password string) error
}
