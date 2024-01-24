package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/router"
)

func main() {
	dbConnection, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("Could not instantiate database: %s", err)
	}
	log.Printf("Database instantiated")

	userRepository := user.NewRepository(dbConnection.GetDB())
	userService := user.NewUserService(userRepository)
	userHandler := user.NewHandler(userService)

	r := router.NewRouter(userHandler)
	router.StartRouter("0.0.0.0:8080", r)
}
