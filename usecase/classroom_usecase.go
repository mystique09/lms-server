package usecase

import (
	"context"
	"github.com/google/uuid"
	"server/database/postgresql"
	"server/domain"
)

type (
	classroomUsecase struct {
		repository domain.ClassroomRepository
	}
)

func NewClassroomUsecase(repository domain.ClassroomRepository) domain.ClassroomUsecase {
	return &classroomUsecase{
		repository: repository,
	}
}

func (clu *classroomUsecase) Create(c context.Context, cl *postgresql.CreateClassParams) error {
	return clu.repository.Create(c, cl)
}

func (clu *classroomUsecase) Delete(c context.Context, id uuid.UUID) error {
	return clu.repository.Delete(c, id)
}

func (clu *classroomUsecase) GetByInviteCode(c context.Context, inviteCode uuid.UUID) (uuid.UUID, error) {
	return clu.repository.GetByInviteCode(c, inviteCode)
}

func (clu *classroomUsecase) GetClasroomsByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]domain.Classroom, error) {
	return clu.repository.GetClasroomsByUser(c, opts)
}

func (clu *classroomUsecase) GetClassroomMembers(c context.Context, id uuid.UUID) ([]domain.ClassroomMember, error) {
	return clu.repository.GetClassroomMembers(c, id)
}

func (clu *classroomUsecase) GetByID(c context.Context, id uuid.UUID) (domain.Classroom, error) {
	return clu.repository.GetByID(c, id)
}

func (clu *classroomUsecase) UpdateClassroomName(c context.Context, cl *postgresql.UpdateClassroomNameParams) error {
	return clu.repository.UpdateClassroomName(c, cl)
}

func (clu *classroomUsecase) UpdateClassroomDescription(c context.Context, cl *postgresql.UpdateClassroomDescriptionParams) error {
	return clu.repository.UpdateClassroomDescription(c, cl)
}

func (clu *classroomUsecase) UpdateClassroomSection(c context.Context, cl *postgresql.UpdateClassroomSectionParams) error {
	return clu.repository.UpdateClassroomSection(c, cl)
}

func (clu *classroomUsecase) UpdateClassroomRoom(c context.Context, cl *postgresql.UpdateClassroomRoomParams) error {
	return clu.repository.UpdateClassroomRoom(c, cl)
}

func (clu *classroomUsecase) UpdateClassroomSubject(c context.Context, cl *postgresql.UpdateClassroomSubjectParams) error {
	return clu.repository.UpdateClassroomSubject(c, cl)
}

func (clu *classroomUsecase) UpdateClassroomInviteCode(c context.Context, cl *postgresql.UpdateClassroomInviteCodeParams) error {
	return clu.repository.UpdateClassroomInviteCode(c, cl)
}
