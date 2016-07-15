package config

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func LoadConfig(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error %s reading %s", err, filename)
		// TODO: Find a cleaner way to create an empty []string
		return strings.Split("", "\n")
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
