package domain

import (
	"context"
	"github.com/google/uuid"
	"server/database/postgresql"
)

type Classroom = postgresql.Classroom

type ClassroomRepository interface {
	Create(c context.Context, cl *postgresql.CreateClassParams) error
	Fetch(c context.Context, offsent int32) ([]Classroom, error)
	GetByID(c context.Context, id uuid.UUID) (Classroom, error)
	GetByInviteCode(c context.Context, inviteCode uuid.UUID) (uuid.UUID, error)
	GetClasroomByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]Classroom, error)
	UpdateClassroomName(c context.Context, cl *postgresql.UpdateClassroomNameParams) error
	UpdateClassroomDescription(c context.Context, cl *postgresql.UpdateClassroomDescriptionParams) error
	UpdateClassroomSection(c context.Context, cl *postgresql.UpdateClassroomSectionParams) error
	UpdateClassroomRoom(c context.Context, cl *postgresql.UpdateClassroomRoomParams) error
	UpdateClassroomSubject(c context.Context, cl *postgresql.UpdateClassroomSubjectParams) error
	UpdateClassroomInviteCode(c context.Context, cl *postgresql.UpdateClassroomInviteCodeParams) error
	Delete(c context.Context, id uuid.UUID) error
}
