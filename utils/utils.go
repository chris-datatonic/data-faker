package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Name      string
	FakerType string
	JsonType  string
}

func UnmarshalFileToJson(path string) []Config {
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

func MarshalStructToJsonFile(path string, data []interface{}) {
	b, err := json.Marshal(data)

	if err != nil {
		log.Fatalf("Unable to marshall data %v", data)
	}

	os.WriteFile(path, b, 0644)
}
