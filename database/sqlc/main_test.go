package database

import (
	"database/sql"
	"log"
	"os"
	"server/config"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	godotenv.Load("./.development.env")
	cfg := config.Init()
	conn, err := sql.Open("postgres", cfg.DATABASE_URL)

	if err != nil {
		log.Fatal(err.Error())
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
