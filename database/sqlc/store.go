package database

import (
	"database/sql"

	"server/config"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db  *sql.DB
	cfg config.Config
	*Queries
}

func NewStore(db *sql.DB, cfg config.Config) Store {
	return &SQLStore{
		db:      db,
		cfg:     cfg,
		Queries: New(db),
	}
}
