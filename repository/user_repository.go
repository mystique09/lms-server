package repository

import (
	"server/database/postgresql"
	"server/database/store"
	"server/domain"

	"context"
	"github.com/google/uuid"
)

type userRepository struct {
	store store.Store
}

func NewUserRepository(st store.Store) domain.UserRepository {
	return &userRepository{
		store: st,
	}
}

func (ur *userRepository) Create(c context.Context, u *postgresql.CreateUserParams) error {
	_, err := ur.store.CreateUser(c, *u)
	return err
}

func (ur *userRepository) Fetch(c context.Context, offset int32) ([]domain.User, error) {
	users, err := ur.store.GetUsers(c, offset*10)
	return users, err
}

func (ur *userRepository) GetByID(c context.Context, id uuid.UUID) (domain.User, error) {
	user, err := ur.store.GetUser(c, id)
	return user, err
}

func (ur *userRepository) GetByUsername(c context.Context, username string) (domain.User, error) {
	user, err := ur.store.GetUserByUsername(c, username)
	return user, err
}

func (ur *userRepository) UpdateUsername(c context.Context, u *postgresql.UpdateUsernameParams) error {
	_, err := ur.store.UpdateUsername(c, *u)
	return err
}

func (ur *userRepository) UpdateEmail(c context.Context, u *postgresql.UpdateUserEmailParams) error {
	_, err := ur.store.UpdateUserEmail(c, *u)
	return err
}

func (ur *userRepository) UpdatePassword(c context.Context, u *postgresql.UpdateUserPasswordParams) error {
	_, err := ur.store.UpdateUserPassword(c, *u)
	return err
}

func (ur *userRepository) Delete(c context.Context, id uuid.UUID) error {
	_, err := ur.store.DeleteUser(c, id)
	return err
}
