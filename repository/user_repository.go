package repository

import (
	"server/database/store"
)

type userRepository struct {
	userStore *store.Store
}

func NewUserRepository(st store.Store) userRepository {
	return userRepository{
		userStore: &st,
	}
}
