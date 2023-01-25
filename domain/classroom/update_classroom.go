package domain

import (
	"context"
	"server/database/postgresql"

	"github.com/google/uuid"
)

type UpdateClassroomUsecase interface {
	GetByID(c context.Context, id uuid.UUID) (Classroom, error)
	UpdateClassroomName(c context.Context, cl *postgresql.UpdateClassroomNameParams) error
	UpdateClassroomDescription(c context.Context, cl *postgresql.UpdateClassroomDescriptionParams) error
	UpdateClassroomSection(c context.Context, cl *postgresql.UpdateClassroomSectionParams) error
	UpdateClassroomRoom(c context.Context, cl *postgresql.UpdateClassroomRoomParams) error
	UpdateClassroomSubject(c context.Context, cl *postgresql.UpdateClassroomSubjectParams) error
	UpdateClassroomInviteCode(c context.Context, cl *postgresql.UpdateClassroomInviteCodeParams) error
}
