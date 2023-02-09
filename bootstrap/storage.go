package bootstrap

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"log"
)

type Storage struct {
	Provider *cloudinary.Cloudinary
}

func NewStorageProvider(env *Env) Storage {
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
