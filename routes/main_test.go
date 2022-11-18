package routes

import (
	"log"
	"os"
	"server/utils"
	"testing"
)

var cfg utils.Config

func TestMain(m *testing.M) {

	conf, err := utils.LoadConfig("../", "app.sample")

	if err != nil {
		log.Fatal(err)
	}

	cfg = conf

	os.Exit(m.Run())
}
