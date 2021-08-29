package config

import (
	"encoding/json"
	"log"
	"os"
)

// JSONValues implements json struct from json config file
type JSONValues struct {
	Environment string                 `json:"environment"`
	Development map[string]interface{} `json:"development"`
	Production  map[string]interface{} `json:"production"`
}

type Config map[string]interface{}

// Configuration read json file from ./config.json and return map[string]interface{}
func Configuration() Config {
	var values JSONValues
	var config Config

	// Read file from path
	configJsonFile, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer configJsonFile.Close()

	var configDecoder *json.Decoder = json.NewDecoder(configJsonFile)
	err = configDecoder.Decode(&values)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch values.Environment {
	case "development":
		config = values.Development
		break
	case "production":
		config = values.Production
		break
	default:
		log.Fatal(`error: "Json config file no contains development or production config"`)
		break
	}

	return config
}
