package utils

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func SetupDB(DATABASE_URL string) *sql.DB {
	db, err := sql.Open("postgres", DATABASE_URL)

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatalf("unable to migrate: %v", err.Error())
	}

	log.Println("Starting database migration")
	m.Up()
	log.Println("Migration success")

	return db
}
