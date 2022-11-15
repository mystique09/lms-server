package routes

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load(".development.env")
	if err != nil {
		log.Fatal(err.Error())
	}

	m.Run()
}
