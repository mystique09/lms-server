package domain

import (
	"context"
	"server/database/postgresql"

	"github.com/google/uuid"
)

type User = postgresql.User

type UserRepository interface {
	Create(c context.Context, u *postgresql.CreateUserParams) error
	Fetch(c context.Context, offsent int32) ([]User, error)
	GetByID(c context.Context, id uuid.UUID) (User, error)
	GetByUsername(c context.Context, email string) (User, error)
	UpdateUsername(c context.Context, u *postgresql.UpdateUsernameParams) error
	UpdateEmail(c context.Context, u *postgresql.UpdateUserEmailParams) error
	UpdatePassword(c context.Context, u *postgresql.UpdateUserPasswordParams) error
	Delete(c context.Context, id uuid.UUID) error
}
