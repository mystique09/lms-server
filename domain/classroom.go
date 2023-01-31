package domain

import (
	"context"
	"server/database/postgresql"

	"github.com/google/uuid"
)

type (
	Classroom       = postgresql.Classroom
	ClassroomMember = postgresql.ClassroomMember
)

type (
	CreateClassroomRequest struct {
		AdminID     uuid.UUID `json:"admin_id" validate:"required"`
		Name        string    `json:"name" validate:"required,gte=8"`
		Description string    `json:"description" validate:"required,gte=1"`
		Section     string    `json:"section" validate:"required,gte=1"`
		Room        string    `json:"room" validate:"required,gte=1"`
		Subject     string    `json:"subject" validate:"required,gte=1"`
	}

	UpdateClassroomRequest struct {
		Name        string `json:"name" validate:"gte=0"`
		Description string `json:"description" validate:"gte=0"`
		Section     string `json:"section" validate:"gte=0"`
		Room        string `json:"room" validate:"gte=0"`
		Subject     string `json:"subject" validate:"gte=0"`
	}
)

type (
	ClassroomRepository interface {
		Create(c context.Context, cl *postgresql.CreateClassParams) error
		Fetch(c context.Context, offsent int32) ([]Classroom, error)
		GetByID(c context.Context, id uuid.UUID) (Classroom, error)
		GetByInviteCode(c context.Context, inviteCode uuid.UUID) (uuid.UUID, error)
		GetClasroomsByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]Classroom, error)
		GetClassroomMembers(c context.Context, id uuid.UUID) ([]ClassroomMember, error)
		UpdateClassroomName(c context.Context, cl *postgresql.UpdateClassroomNameParams) error
		UpdateClassroomDescription(c context.Context, cl *postgresql.UpdateClassroomDescriptionParams) error
		UpdateClassroomSection(c context.Context, cl *postgresql.UpdateClassroomSectionParams) error
		UpdateClassroomRoom(c context.Context, cl *postgresql.UpdateClassroomRoomParams) error
		UpdateClassroomSubject(c context.Context, cl *postgresql.UpdateClassroomSubjectParams) error
		UpdateClassroomInviteCode(c context.Context, cl *postgresql.UpdateClassroomInviteCodeParams) error
		Delete(c context.Context, id uuid.UUID) error
	}

	ClassroomUsecase interface {
		GetByID(c context.Context, id uuid.UUID) (Classroom, error)
		GetByInviteCode(c context.Context, inviteCode uuid.UUID) (uuid.UUID, error)
		GetClasroomsByUser(c context.Context, opts postgresql.GetAllClassFromUserParams) ([]Classroom, error)
		GetClassroomMembers(c context.Context, id uuid.UUID) ([]ClassroomMember, error)
		Create(c context.Context, u *postgresql.CreateClassParams) error
		UpdateClassroomName(c context.Context, cl *postgresql.UpdateClassroomNameParams) error
		UpdateClassroomDescription(c context.Context, cl *postgresql.UpdateClassroomDescriptionParams) error
		UpdateClassroomSection(c context.Context, cl *postgresql.UpdateClassroomSectionParams) error
		UpdateClassroomRoom(c context.Context, cl *postgresql.UpdateClassroomRoomParams) error
		UpdateClassroomSubject(c context.Context, cl *postgresql.UpdateClassroomSubjectParams) error
		UpdateClassroomInviteCode(c context.Context, cl *postgresql.UpdateClassroomInviteCodeParams) error
		Delete(c context.Context, id uuid.UUID) error
	}
)
