package utils

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func SetupDB(DATABASE_URL string) *sql.DB {
	db, err := sql.Open("postgres", DATABASE_URL)

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatalf("error while migrating: %v", err.Error())
	}

	log.Println("-- Migration started --")
	migrateErr := m.Up()
	if migrateErr != nil {
		log.Fatal(migrateErr.Error())
	}
	log.Println("-- Migration done --")

	return db
}
