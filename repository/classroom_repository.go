package repository

import (
	"context"
	"server/database/postgresql"
	"server/database/store"
	"server/domain"

	"github.com/google/uuid"
)

type classroomRepository struct {
	store store.Store
}

func NewClassroomRepository(st store.Store) domain.ClassroomRepository {
	return &classroomRepository{
		store: st,
	}
}

func (cr *classroomRepository) Create(c context.Context, crp *postgresql.CreateClassParams) error {
	_, err := cr.store.CreateClass(c, *crp)
	return err
}

func (cr *classroomRepository) Fetch(c context.Context, offset int32) ([]domain.Classroom, error) {
	classrooms, err := cr.store.ListAllPublicClass(c, offset*10)
	return classrooms, err
}

func (cr *classroomRepository) GetByID(c context.Context, id uuid.UUID) (domain.Classroom, error) {
	classroom, err := cr.store.GetClass(c, id)
	return classroom, err
}

func (cr *classroomRepository) GetByInviteCode(c context.Context, inviteCode uuid.UUID) (uuid.UUID, error) {
	classroom, err := cr.store.GetClassroomWithInviteCode(c, inviteCode)
	return classroom, err
}

func (cr *classroomRepository) GetClasroomByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]domain.Classroom, error) {
	classrooms, err := cr.store.GetAllClassFromUser(c, opts)
	return classrooms, err
}

func (cr *classroomRepository) UpdateClassroomName(c context.Context, u *postgresql.UpdateClassroomNameParams) error {
	_, err := cr.store.UpdateClassroomName(c, *u)
	return err
}

func (cr *classroomRepository) UpdateClassroomDescription(c context.Context, u *postgresql.UpdateClassroomDescriptionParams) error {
	_, err := cr.store.UpdateClassroomDescription(c, *u)
	return err
}

func (cr *classroomRepository) UpdateClassroomSection(c context.Context, u *postgresql.UpdateClassroomSectionParams) error {
	_, err := cr.store.UpdateClassroomSection(c, *u)
	return err
}

func (cr *classroomRepository) UpdateClassroomRoom(c context.Context, u *postgresql.UpdateClassroomRoomParams) error {
	_, err := cr.store.UpdateClassroomRoom(c, *u)
	return err
}

func (cr *classroomRepository) UpdateClassroomSubject(c context.Context, u *postgresql.UpdateClassroomSubjectParams) error {
	_, err := cr.store.UpdateClassroomSubject(c, *u)
	return err
}

func (cr *classroomRepository) UpdateClassroomInviteCode(c context.Context, u *postgresql.UpdateClassroomInviteCodeParams) error {
	_, err := cr.store.UpdateClassroomInviteCode(c, *u)
	return err
}

func (cr *classroomRepository) Delete(c context.Context, id uuid.UUID) error {
	_, err := cr.store.DeleteClass(c, id)
	return err
}
