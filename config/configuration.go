package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

var configuration map[string]interface{} = nil

func loadConfig(filename string) map[string]interface{} {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		absPath, errFilePath := filepath.Abs(filename)
		if errFilePath != nil {
			log.Printf("Error getting absolute filepath for %s: %s\n", filename, errFilePath)
			return nil
		}
		log.Printf("Error reading %s %s\n", absPath, err)
		return nil
	}
	var parsed map[string]interface{}
	err = json.Unmarshal(content, &parsed)
	if err != nil {
		return nil
	}

	return parsed
}

func GetConfig() map[string]interface{} {
	if configuration == nil {
		userFile := "config.json"
		configuration = loadConfig(userFile)
	}
	return configuration
}
