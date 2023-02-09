package domain

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/google/uuid"
)

type AssetID = string

type (
	ClassworkRepository interface {
		GetOne(c context.Context, id AssetID) (admin.AssetResult, error)
		GetAllClassworksByUserID(c context.Context, id uuid.UUID) (admin.AssetsResult, error)
		GetAllClassworksByClassID(c context.Context, id uuid.UUID) (admin.AssetsResult, error)
		Upload(c context.Context, buffer []byte) error
		Delete(c context.Context, id AssetID) error
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
