package domain

import (
	"context"

	"github.com/google/uuid"
)

type ProfileUsecase interface {
	GetProfile(c context.Context, id uuid.UUID) (User, error)
	GetClassrooms(c context.Context, id uuid.UUID) ([]Classroom, error)
}
