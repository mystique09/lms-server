package domain

import (
	"context"

	"github.com/google/uuid"
)

type DeleteClassroomUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (Classroom, error)
	Delete(c context.Context, id uuid.UUID) error
}
