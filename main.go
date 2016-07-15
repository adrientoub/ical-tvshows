package main

import (
	"fmt"
	"github.com/adrientoub/ical-tvshows/config"
	"github.com/adrientoub/ical-tvshows/server"
)

func main() {
	userFile := "config.txt"
	config := config.LoadConfig(userFile)
	port := config[0]
	fmt.Println("Listening on " + port)
	server.Listen(port)
}
