package main

import (
	"flag"
	"log"
	"server/api"
	"server/database/seeder"
	"server/utils"
)

var (
	seedDatabase = flag.Bool("seed", false, "Seed the database with random data...")
	seedAmount   = flag.Int("amount", 10, "amount of data to put in the database.")
	tableName    = flag.String("table", "users", "the database table to seed the data")
)

func main() {
	flag.Parse()

	cfg, err := utils.LoadConfig(".", "app")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if *seedDatabase {
		seedRunner := seeder.New(*tableName, *seedAmount, &cfg)
		seedRunner.Run()
		log.Println("Seeding done")
		return
	}

	api.Launch(&cfg)
}
