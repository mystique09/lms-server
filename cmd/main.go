package main

import (
	"log"
	"server/api"
	"server/utils"
)

func main() {
	cfg, err := utils.LoadConfig(".", "app")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.Launch(&cfg)
}
