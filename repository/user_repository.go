package repository

import (
	"server/database/store"
)

type userRepository struct {
	database *store.Store
}

func NewUserRepository(store *store.Store) userRepository {
	return userRepository{
		database: store,
	}
}
