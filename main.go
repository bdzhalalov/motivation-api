package main

import (
	"log"
	"motivations-api/config"
	"motivations-api/lib"
)

func main() {
	cfg := config.InitConfig()

	server := lib.Init(&cfg)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
