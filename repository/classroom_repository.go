package repository

import (
	"context"
	"server/database/postgresql"
	"server/database/store"
	"server/domain/classroom"

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

func (cr *classroomRepository) Create(c context.Context, cl *postgresql.CreateClassParams) error {
	_, err := cr.store.CreateClass(c, *cl)
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

func (cr *classroomRepository) UpdateClassroomName(c context.Context, cl *postgresql.UpdateClassroomNameParams) error {
	_, err := cr.store.UpdateClassroomName(c, *cl)
	return err
}

func (cr *classroomRepository) UpdateClassroomDescription(c context.Context, cl *postgresql.UpdateClassroomDescriptionParams) error {
	_, err := cr.store.UpdateClassroomDescription(c, *cl)
	return err
}

func (cr *classroomRepository) UpdateClassroomSection(c context.Context, cl *postgresql.UpdateClassroomSectionParams) error {
	_, err := cr.store.UpdateClassroomSection(c, *cl)
	return err
}

func (cr *classroomRepository) UpdateClassroomRoom(c context.Context, cl *postgresql.UpdateClassroomRoomParams) error {
	_, err := cr.store.UpdateClassroomRoom(c, *cl)
	return err
}

func (cr *classroomRepository) UpdateClassroomSubject(c context.Context, cl *postgresql.UpdateClassroomSubjectParams) error {
	_, err := cr.store.UpdateClassroomSubject(c, *cl)
	return err
}

func (cr *classroomRepository) UpdateClassroomInviteCode(c context.Context, cl *postgresql.UpdateClassroomInviteCodeParams) error {
	_, err := cr.store.UpdateClassroomInviteCode(c, *cl)
	return err
}

func (cr *classroomRepository) Delete(c context.Context, id uuid.UUID) error {
	_, err := cr.store.DeleteClass(c, id)
	return err
}
