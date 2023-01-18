package domain

import (
	"context"
	"server/database/postgresql"

	"github.com/google/uuid"
)

type Classroom = postgresql.Classroom

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

type ClassroomRepository interface {
	Create(c context.Context, u *postgresql.CreateClassParams) error
	Fetch(c context.Context, offsent int32) ([]Classroom, error)
	GetByID(c context.Context, id uuid.UUID) (Classroom, error)
	GetByInviteCode(c context.Context, inviteCode uuid.UUID) (uuid.UUID, error)
	GetClasroomByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]Classroom, error)
	UpdateClassroomName(c context.Context, u *postgresql.UpdateClassroomNameParams) error
	UpdateClassroomDescription(c context.Context, u *postgresql.UpdateClassroomDescriptionParams) error
	UpdateClassroomSection(c context.Context, u *postgresql.UpdateClassroomSectionParams) error
	UpdateClassroomRoom(c context.Context, u *postgresql.UpdateClassroomRoomParams) error
	UpdateClassroomSubject(c context.Context, u *postgresql.UpdateClassroomSubjectParams) error
	UpdateClassroomInviteCode(c context.Context, u *postgresql.UpdateClassroomInviteCodeParams) error
	Delete(c context.Context, id uuid.UUID) error
}
