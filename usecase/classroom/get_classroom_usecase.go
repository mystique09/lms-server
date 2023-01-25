package usecase

import (
	"context"
	"server/database/postgresql"
	"server/domain/classroom"

	"github.com/google/uuid"
)

func NewGetClassroomUsecase(repository domain.ClassroomRepository) domain.GetClassroomUsecase {
	return &getClassroomUsecase{
		repository: repository,
	}
}

func (gcu *getClassroomUsecase) GetByID(c context.Context, id uuid.UUID) (domain.Classroom, error) {
	return gcu.repository.GetByID(c, id)
}

func (gcu *getClassroomUsecase) GetByInviteCode(c context.Context, inviteCode uuid.UUID) (uuid.UUID, error) {
	return gcu.repository.GetByInviteCode(c, inviteCode)
}

func (gcu *getClassroomUsecase) GetClasroomByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]domain.Classroom, error) {
	return gcu.repository.GetClasroomByUser(c, opts)
}
