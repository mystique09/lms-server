package usecase

import (
	"context"
	"github.com/google/uuid"
	"server/database/postgresql"
	"server/domain"
)

type (
	getClassroomUsecase struct {
		repository domain.ClassroomRepository
	}
	createClassroomUsecase struct {
		repository domain.ClassroomRepository
	}
	updateClassroomUsecase struct {
		repository domain.ClassroomRepository
	}
	deleteClassroomUsecase struct {
		repository domain.ClassroomRepository
	}
)

func NewCreateClassroomUsecase(repository domain.ClassroomRepository) domain.CreateClassroomUsecase {
	return &createClassroomUsecase{
		repository: repository,
	}
}

func (ccu *createClassroomUsecase) Create(c context.Context, cl *postgresql.CreateClassParams) error {
	return ccu.repository.Create(c, cl)
}

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

func (gcu *getClassroomUsecase) GetClasroomsByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]domain.Classroom, error) {
	return gcu.repository.GetClasroomsByUser(c, opts)
}

func (gcu *getClassroomUsecase) GetClassroomMembers(c context.Context, id uuid.UUID) ([]domain.ClassroomMember, error) {
	return gcu.repository.GetClassroomMembers(c, id)
}

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
