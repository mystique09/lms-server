package domain

import (
	"context"
	"server/database/postgresql"

	"github.com/google/uuid"
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

type CreateClassroomRequest struct {
	AdminID     uuid.UUID `json:"admin_id" validate:"required"`
	Name        string    `json:"name" validate:"required,gte=8"`
	Description string    `json:"description" validate:"required,gte=1"`
	Section     string    `json:"section" validate:"required,gte=1"`
	Room        string    `json:"room" validate:"required,gte=1"`
	Subject     string    `json:"subject" validate:"required,gte=1"`
}

type CreateClassroomResponse struct {
	Classroom Classroom `json:"classroom"`
}

type CreateClassroomUsecase interface {
	Create(c context.Context, u *postgresql.CreateClassParams) error
}

type DeleteClassroomUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (Classroom, error)
	Delete(c context.Context, id uuid.UUID) error
}

type GetClassroomUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (Classroom, error)
	GetByInviteCode(c context.Context, inviteCode uuid.UUID) (uuid.UUID, error)
	GetClasroomByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]Classroom, error)
}

type UpdateClassroomUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (Classroom, error)
	UpdateClassroomName(c context.Context, cl *postgresql.UpdateClassroomNameParams) error
	UpdateClassroomDescription(c context.Context, cl *postgresql.UpdateClassroomDescriptionParams) error
	UpdateClassroomSection(c context.Context, cl *postgresql.UpdateClassroomSectionParams) error
	UpdateClassroomRoom(c context.Context, cl *postgresql.UpdateClassroomRoomParams) error
	UpdateClassroomSubject(c context.Context, cl *postgresql.UpdateClassroomSubjectParams) error
	UpdateClassroomInviteCode(c context.Context, cl *postgresql.UpdateClassroomInviteCodeParams) error
}
