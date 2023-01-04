package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Name      string
	FakerType string
	JsonType  string
}

func MarshalFileToJson(path string) []Config {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Failed to read file %v", err)
	}

	var result []Config

	err = json.Unmarshal(data, &result)

	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return result
}
