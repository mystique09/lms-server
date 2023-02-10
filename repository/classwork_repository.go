package repository

import (
	"bytes"
	"context"
	"server/bootstrap"
	"server/database/postgresql"
	"server/database/store"
	"server/domain"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/google/uuid"
)

type classworkRepository struct {
	store   store.Store
	storage bootstrap.Storage
}

func NewClassworkRepository(s store.Store, storage bootstrap.Storage) domain.ClassworkRepository {
	return &classworkRepository{
		store:   s,
		storage: storage,
	}
}

// TODO!: implement methods

func (cw *classworkRepository) GetOne(c context.Context, id, user_id uuid.UUID) (domain.Classwork, error) {
	class_work, err := cw.store.GetClassWork(c, postgresql.GetClassWorkParams{
		ID:     id,
		UserID: user_id,
	})
	return class_work, err
}

func (cw *classworkRepository) GetAllClassworksByUserID(c context.Context, user_id uuid.UUID, offset int32) ([]domain.Classwork, error) {
	class_works, err := cw.store.ListSubmittedClassworks(c, postgresql.ListSubmittedClassworksParams{
		UserID: user_id,
		Offset: offset,
	})
	return class_works, err
}

func (cw *classworkRepository) GetAllClassworksByClassID(c context.Context, class_id uuid.UUID, offset int32) ([]domain.Classwork, error) {
	class_works, err := cw.store.ListClassworkAdmin(c, postgresql.ListClassworkAdminParams{
		ClassID: class_id,
		Offset:  offset,
	})
	return class_works, err
}

// TODO!: implement upload method
func (cw *classworkRepository) Upload(c context.Context, data bytes.Buffer) error {
	return nil
}

func (cw *classworkRepository) Delete(c context.Context, id uuid.UUID) (*string, error) {
	class_work, err := cw.store.DeleteClassworkFromClass(c, postgresql.DeleteClassworkFromClassParams{
		ID: id,
	})

	if err != nil {
		return nil, err
	}

	asset, err := cw.storage.DeleteAsset(c, class_work.Url)
	if err != nil {
		return nil, err
	}
	deleted_asset := asset.(admin.AssetResult)

	return &deleted_asset.AssetID, nil
}
