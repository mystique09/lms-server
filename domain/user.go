package domain

import (
	"server/database/postgresql"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type User = postgresql.User

type UserRepository interface {
	Create(c echo.Context, u *postgresql.CreateUserParams) error
	Fetch(c echo.Context, offsent int32) ([]User, error)
	GetByID(c echo.Context, id uuid.UUID) (User, error)
	GetByUsername(c echo.Context, email string) (User, error)
	UpdateUsername(c echo.Context, u *postgresql.UpdateUsernameParams) error
	UpdateEmail(c echo.Context, u *postgresql.UpdateUserEmailParams) error
	UpdatePassword(c echo.Context, u *postgresql.UpdateUserPasswordParams) error
	Delete(c echo.Context, id uuid.UUID) error
}
