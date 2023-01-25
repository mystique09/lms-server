package usecase

import (
	domain "server/domain/classroom"

	"context"
	"github.com/google/uuid"
)

func NewDeleteClassroomUsecase(repository domain.ClassroomRepository) domain.DeleteClassroomUsecase {
	return &deleteClassroomUsecase{
		repository: repository,
	}
}

func (dcu *deleteClassroomUsecase) GetByID(c context.Context, id uuid.UUID) (domain.Classroom, error) {
	return dcu.repository.GetByID(c, id)
}

func (dcu *deleteClassroomUsecase) Delete(c context.Context, id uuid.UUID) error {
	return dcu.repository.Delete(c, id)
}
