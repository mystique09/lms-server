package domain

import (
	"bytes"
	"context"
	"server/database/postgresql"

	"github.com/google/uuid"
)

type AssetID = string
type Classwork = postgresql.ClassWork

type (
	ClassworkRepository interface {
		GetOne(c context.Context, id, user_id uuid.UUID) (Classwork, error)
		GetAllClassworksByUserID(c context.Context, user_id uuid.UUID, offset int32) ([]Classwork, error)
		GetAllClassworksByClassID(c context.Context, class_id uuid.UUID, offset int32) ([]Classwork, error)
		Upload(c context.Context, data bytes.Buffer) error
		Delete(c context.Context, id uuid.UUID) (*string, error)
	}

	// TODO!: implement usecase for classwork
	ClassworkUsecase interface {
		GetClassworks()
		ViewUserClassworks()
		ViewSubmittedClassworks()
		UploadClasswork()
		DeleteClasswork()
	}
)
