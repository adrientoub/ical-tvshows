package main

import (
	"fmt"
	"github.com/adrientoub/ical-tvshows/config"
	"github.com/adrientoub/ical-tvshows/server"
	"log"
)

func main() {
	userFile := "config.json"
	config := config.LoadConfig(userFile)
	if config == nil {
		log.Fatal("Config file is not correct.")
	}
	port := config["uri"].(string)
	fmt.Println("Listening on " + port)
	server.Listen(port)
}
