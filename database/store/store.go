package store

import (
	"database/sql"
	"server/database/postgresql"
)

type Store interface {
}

type SQLStore struct {
	*postgresql.Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: postgresql.New(db),
	}
}
