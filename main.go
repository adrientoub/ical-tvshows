package main

import (
	"fmt"
	"github.com/adrientoub/ical-tvshows/config"
	"github.com/adrientoub/ical-tvshows/server"
	"log"
)

func main() {
	config := config.GetConfig()
	if config == nil {
		log.Fatal("Config file is not correct.")
	}
	port := config["uri"].(string)
	fmt.Println("Listening on " + port)
	server.Listen(port)
}
