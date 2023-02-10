package bootstrap

import (
	"bytes"
	"context"
	"log"
	"server/internal/stringutil"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type (
	Storage struct {
		Provider *cloudinary.Cloudinary
	}

	IStorage interface {
		GetAsset(c context.Context, id string, asset_type api.AssetType) (interface{}, error)
		GetAssets(c context.Context, asset_type api.AssetType) (interface{}, error)
		UploadAsset(c context.Context, data bytes.Buffer, asset_type api.AssetType, file_name, folder string) (interface{}, error)
		UpdateAsset(c context.Context, id string) (interface{}, error)
		DeleteAsset(c context.Context, id string) (interface{}, error)
	}
)

func NewStorageProvider(env *Env) IStorage {
	log.Println("initiating storage provider")
	provider, err := cloudinary.NewFromURL(env.CldUrl)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("storage provider connected")
	return Storage{
		Provider: provider,
	}
}

func (s Storage) GetAsset(c context.Context, id string, asset_type api.AssetType) (interface{}, error) {
	asset, err := s.Provider.Admin.Asset(c, admin.AssetParams{
		PublicID:  id,
		AssetType: asset_type,
	})
	return asset, err
}

func (s Storage) GetAssets(c context.Context, asset_type api.AssetType) (interface{}, error) {
	assets, err := s.Provider.Admin.Assets(c, admin.AssetsParams{
		AssetType: asset_type,
	})
	return assets, err
}

func (s Storage) UploadAsset(c context.Context, data bytes.Buffer, asset_type api.AssetType, file_name, folder string) (interface{}, error) {
	asset, err := s.Provider.Upload.Upload(c, data, uploader.UploadParams{
		PublicID:       stringutil.RandomString(12),
		PublicIDPrefix: file_name,
		UniqueFilename: api.Bool(false),
		DisplayName:    file_name,
		AssetFolder:    folder,
	})
	return asset, err
}

// TODO!:
// we need to have a method that can update an asset name/data
func (s Storage) UpdateAsset(c context.Context, id string) (interface{}, error) {
	return nil, nil
}

func (s Storage) DeleteAsset(c context.Context, id string) (interface{}, error) {
	asset, err := s.Provider.Admin.DeleteAssets(c, admin.DeleteAssetsParams{
		PublicIDs: api.CldAPIArray{id},
	})
	return asset, err
}
