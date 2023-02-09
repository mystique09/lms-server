package bootstrap

import (
	"database/sql"
	"log"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func NewPostgresqlClient(env *Env) *sql.DB {
	db, err := sql.Open("postgres", env.DBUrl)

	if err != nil {
		log.Fatal(err)
	}

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

	log.Println("starting migration")

	migrateErr := m.Up()
	if migrateErr != nil {
		if strings.Contains(migrateErr.Error(), "no change") {
			log.Println("migration: ", migrateErr.Error())
		} else {
			log.Println("migration err: ", migrateErr.Error())
		}
	}

	log.Println("migration done")

	return db
}

func ClosePostgresqlConnection(client *sql.DB) {
	if client == nil {
		return
	}

	err := client.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Panicln("connection to postgresql is now close.")
}
