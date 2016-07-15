package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var configuration map[string]interface{}

func LoadConfig(filename string) map[string]interface{} {
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
	configuration = parsed

	return parsed
}

func GetConfig() map[string]interface{} {
	return configuration
}
