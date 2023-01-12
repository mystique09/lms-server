package bootstrap

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func NewPostgresqlClient(env *Env) *sql.DB {
	databaseUrl := env.DBUrl
	db, err := sql.Open("postgres", databaseUrl)

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

	log.Println("-- Migration started --")

	migrateErr := m.Up()
	if migrateErr != nil {
		log.Println("Migration err: ", migrateErr.Error())
	}

	log.Println("-- Migration done --")

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

	log.Panicln("Connection to Postgresql is now close.")
}
