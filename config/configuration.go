package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var configuration map[string]interface{} = nil

func loadConfig(filename string) map[string]interface{} {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error %s\n", err)
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
