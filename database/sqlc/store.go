package database

import (
	"database/sql"

	"server/utils"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db  *sql.DB
	cfg utils.Config
	*Queries
}

func NewStore(db *sql.DB, cfg *utils.Config) Store {
	return &SQLStore{
		db:      db,
		cfg:     *cfg,
		Queries: New(db),
	}
}
