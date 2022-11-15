package main

import (
	"log"
	"server/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	routes.Launch()
}
