package main

import (
	"motivations-api/config"
	"motivations-api/internal"
)

func main() {
	cfg := config.InitConfig()

	logger := internal.Logger(&cfg)

	server := internal.Init(&cfg, logger)

	db, err := internal.ConnectToDB(&cfg, logger)

	if err != nil {
		logger.Fatalf("Error connecting to database: %v\n", err)
	}

	defer db.Close()

	if err := server.Run(); err != nil {
		logger.Fatalf("Can't start server: %v", err)
	}
}
