package main

import (
	"log"

	"github.com/adrientoub/ical-tvshows/config"
	"github.com/adrientoub/ical-tvshows/server"
)

func main() {
	config := config.GetConfig()
	if config == nil {
		log.Fatal("Config file is not correct.")
	}
	port := config["uri"].(string)
	log.Println("Listening on " + port)
	server.Listen(port)
}
