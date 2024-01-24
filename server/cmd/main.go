package main

import (
	"log"
	"server/config"
	"server/db"
	"server/internal/user"
	"server/internal/websocket"
	"server/router"
)

func main() {

	// instantiate database
	dbConnection, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not instantiate database: %s", err)
	}
	log.Printf("Database instantiated")

	// load environment variables from config
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Could not load environment variables: %s", err)
	}

	// instantiate repository, service and handler
	userRepository := user.NewRepository(dbConnection.GetDB())
	userService := user.NewUserService(userRepository)
	userHandler := user.NewHandler(userService)

	webSocketHub := websocket.NewHub()
	webSocketHandler := websocket.NewHandler(webSocketHub)

	// create router and start server
	r := router.NewRouter(userHandler, webSocketHandler)
	router.StartRouter("0.0.0.0:8080", r)
}
