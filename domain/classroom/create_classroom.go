package domain

import (
	"github.com/google/uuid"
	"server/database/postgresql"

	"context"
)

type CreateClassroomRequest struct {
	AdminID     uuid.UUID `json:"admin_id" validate:"required"`
	Name        string    `json:"name" validate:"required,gte=8"`
	Description string    `json:"description" validate:"required,gte=1"`
	Section     string    `json:"section" validate:"required,gte=1"`
	Room        string    `json:"room" validate:"required,gte=1"`
	Subject     string    `json:"subject" validate:"required,gte=1"`
}

type CreateClassroomResponse struct {
	Classroom Classroom `json:"classroom"`
}

type CreateClassroomUsecase interface {
	Create(c context.Context, u *postgresql.CreateClassParams) error
}
