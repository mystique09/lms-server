package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

var DATABASE_URL string = "postgres://mystique09:mystique09@localhost:5432/class-manager?sslmode=disable"

func TestMain(m *testing.M) {
	conn, err := sql.Open("postgres", DATABASE_URL)

	if err != nil {
		log.Fatal(err.Error())
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
