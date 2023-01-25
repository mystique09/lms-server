package domain

import (
	"context"
	"server/database/postgresql"

	"github.com/google/uuid"
)

type GetClassroomUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (Classroom, error)
	GetByInviteCode(c context.Context, inviteCode uuid.UUID) (uuid.UUID, error)
	GetClasroomByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]Classroom, error)
}
