package main

import (
	"log"
	"server/db"
)

func main() {
	_, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("Could not instantiate database: %s", err)
	} else {
		log.Printf("Database instantiated successfully")
	}
}
