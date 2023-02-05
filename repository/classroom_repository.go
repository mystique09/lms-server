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

func (cr *classroomRepository) Create(c context.Context, cl *postgresql.CreateClassParams) error {
	_, err := cr.store.CreateClass(c, *cl)
	return err
}

func (cr *classroomRepository) JoinClassroom(c context.Context, opt postgresql.AddNewClassroomMemberParams) (domain.ClassroomMember, error) {
	return cr.store.AddNewClassroomMember(c, opt)
}

func (cr *classroomRepository) Fetch(c context.Context, offset int32) ([]domain.Classroom, error) {
	return cr.store.ListAllPublicClass(c, offset*10)
}

func (cr *classroomRepository) GetByID(c context.Context, id uuid.UUID) (domain.Classroom, error) {
	return cr.store.GetClass(c, id)
}

func (cr *classroomRepository) GetByInviteCode(c context.Context, inviteCode uuid.UUID) (uuid.UUID, error) {
	return cr.store.GetClassroomWithInviteCode(c, inviteCode)
}

func (cr *classroomRepository) GetClasroomsByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]domain.Classroom, error) {
	return cr.store.GetAllClassFromUser(c, opts)
}

func (cr *classroomRepository) GetClassroomMembers(c context.Context, id uuid.UUID) ([]domain.ClassroomMember, error) {
	return cr.store.GetAllClassroomMembers(c, postgresql.GetAllClassroomMembersParams{
		ClassID: id,
		Offset:  0,
	})
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

func (cr *classroomRepository) Leave(c context.Context, opt postgresql.LeaveClassroomParams) (domain.ClassroomMember, error) {
	return cr.store.LeaveClassroom(c, opt)
}
