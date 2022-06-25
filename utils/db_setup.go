package utils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func SetupDB(DATABASE_URL string) *sql.DB {
	db, err := sql.Open("postgres", DATABASE_URL)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return db
}
