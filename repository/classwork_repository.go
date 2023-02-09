package repository

import (
	"context"
	"github.com/google/uuid"
	"server/bootstrap"
	"server/database/store"
	"server/domain"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

type classworkRepository struct {
	store    store.Store
	provider bootstrap.Storage
}

func NewClassworkRepository(s store.Store, p bootstrap.Storage) domain.ClassworkRepository {
	return &classworkRepository{
		store:    s,
		provider: p,
	}
}

// TODO!: implement methods

func (cw *classworkRepository) GetOne(c context.Context, id domain.AssetID) (admin.AssetResult, error) {
	return admin.AssetResult{}, nil
}

func (cw *classworkRepository) GetAllClassworksByUserID(c context.Context, id uuid.UUID) (admin.AssetsResult, error) {
	return admin.AssetsResult{}, nil
}
func (cw *classworkRepository) GetAllClassworksByClassID(c context.Context, id uuid.UUID) (admin.AssetsResult, error) {
	return admin.AssetsResult{}, nil
}
func (cw *classworkRepository) Upload(c context.Context, buffer []byte) error {
	return nil
}
func (cw *classworkRepository) Delete(c context.Context, id domain.AssetID) error {
	return nil
}
