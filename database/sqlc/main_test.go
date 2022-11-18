package database

import (
	"database/sql"
	"log"
	"os"
	"server/utils"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

var testQuesries *Queries

func TestMain(m *testing.M) {
	cfg, err := utils.LoadConfig("../..", "app.sample")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conn, err := sql.Open(dbDriver, cfg.DBUrl)

	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	testQuesries = New(conn)
	os.Exit(m.Run())
}
