package config

import (
	"os"
	"encoding/json"
	"fmt"
)

type Configuration struct {
	Port int
	Message string
	Database string
}

func LoadConfig() (Configuration, error) {
	env := os.Getenv("GO_ENV")
	file, err := os.Open(fmt.Sprintf("./config/config.%s.json", env))
	configuration := Configuration{}
	if err != nil { return configuration, err }
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil { return configuration, err }
	return configuration, nil
}
