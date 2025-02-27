package main

import (
	"motivations-api/config"
	"motivations-api/internal/database"
	"motivations-api/internal/logger"
	"motivations-api/internal/server"
)

func main() {
	cfg := config.InitConfig()

	log := logger.Logger(&cfg)

	db, err := database.ConnectToDB(&cfg, log)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer db.Close()

	apiServer := server.Init(&cfg, db, log)

	if err := apiServer.Run(); err != nil {
		log.Fatalf("Can't start server: %v", err)
	}
}
