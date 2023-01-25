package usecase

import (
	"context"
	"server/database/postgresql"
	"server/domain/classroom"

	"github.com/google/uuid"
)

func NewUpdateClassroomUsecase(repository domain.ClassroomRepository) domain.UpdateClassroomUsecase {
	return &updateClassroomUsecase{
		repository: repository,
	}
}

func (ucu *updateClassroomUsecase) GetByID(c context.Context, id uuid.UUID) (domain.Classroom, error) {
	return ucu.repository.GetByID(c, id)
}

func (ucu *updateClassroomUsecase) UpdateClassroomName(c context.Context, cl *postgresql.UpdateClassroomNameParams) error {
	return ucu.repository.UpdateClassroomName(c, cl)
}

func (ucu *updateClassroomUsecase) UpdateClassroomDescription(c context.Context, cl *postgresql.UpdateClassroomDescriptionParams) error {
	return ucu.repository.UpdateClassroomDescription(c, cl)
}

func (ucu *updateClassroomUsecase) UpdateClassroomSection(c context.Context, cl *postgresql.UpdateClassroomSectionParams) error {
	return ucu.repository.UpdateClassroomSection(c, cl)
}

func (ucu *updateClassroomUsecase) UpdateClassroomRoom(c context.Context, cl *postgresql.UpdateClassroomRoomParams) error {
	return ucu.repository.UpdateClassroomRoom(c, cl)
}

func (ucu *updateClassroomUsecase) UpdateClassroomSubject(c context.Context, cl *postgresql.UpdateClassroomSubjectParams) error {
	return ucu.repository.UpdateClassroomSubject(c, cl)
}

func (ucu *updateClassroomUsecase) UpdateClassroomInviteCode(c context.Context, cl *postgresql.UpdateClassroomInviteCodeParams) error {
	return ucu.repository.UpdateClassroomInviteCode(c, cl)
}
