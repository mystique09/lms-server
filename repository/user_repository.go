package repository

import (
	"server/database/postgresql"
	"server/database/store"
	"server/domain"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type userRepository struct {
	store store.Store
}

func NewUserRepository(st store.Store) domain.UserRepository {
	return &userRepository{
		store: st,
	}
}

func (ur *userRepository) Create(c echo.Context, u *postgresql.CreateUserParams) error {
	_, err := ur.store.CreateUser(c.Request().Context(), *u)
	return err
}

func (ur *userRepository) Fetch(c echo.Context, offset int32) ([]domain.User, error) {
	users, err := ur.store.GetUsers(c.Request().Context(), offset*10)
	return users, err
}

func (ur *userRepository) GetByID(c echo.Context, id uuid.UUID) (domain.User, error) {
	user, err := ur.store.GetUser(c.Request().Context(), id)
	return user, err
}

func (ur *userRepository) GetByUsername(c echo.Context, username string) (domain.User, error) {
	user, err := ur.store.GetUserByUsername(c.Request().Context(), username)
	return user, err
}

func (ur *userRepository) UpdateUsername(c echo.Context, u *postgresql.UpdateUsernameParams) error {
	_, err := ur.store.UpdateUsername(c.Request().Context(), *u)
	return err
}

func (ur *userRepository) UpdateEmail(c echo.Context, u *postgresql.UpdateUserEmailParams) error {
	_, err := ur.store.UpdateUserEmail(c.Request().Context(), *u)
	return err
}

func (ur *userRepository) UpdatePassword(c echo.Context, u *postgresql.UpdateUserPasswordParams) error {
	_, err := ur.store.UpdateUserPassword(c.Request().Context(), *u)
	return err
}

func (ur *userRepository) Delete(c echo.Context, id uuid.UUID) error {
	_, err := ur.store.DeleteUser(c.Request().Context(), id)
	return err
}
