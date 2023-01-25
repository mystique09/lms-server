package usecase

import (
	"server/database/postgresql"
	"server/domain/classroom"

	"golang.org/x/net/context"
)

func NewCreateClassroomUsecase(repository domain.ClassroomRepository) domain.CreateClassroomUsecase {
	return &createClassroomUsecase{
		repository: repository,
	}
}

func (ccu *createClassroomUsecase) Create(c context.Context, cl *postgresql.CreateClassParams) error {
	return ccu.repository.Create(c, cl)
}
