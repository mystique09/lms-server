package main

import (
	"log"
	"server/routes"
	"server/utils"
)

func main() {
	cfg, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.Launch(&cfg)
}
