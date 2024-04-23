package main

import (
	"log"
	"motivations-api/config"
	"motivations-api/internal"
)

func main() {
	cfg := config.InitConfig()

	server := internal.Init(&cfg)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
