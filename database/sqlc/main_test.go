package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://mystique09:mystique09@localhost/class-manager?sslmode=disable"
)

var testQuesries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	testQuesries = New(conn)
	os.Exit(m.Run())
}
