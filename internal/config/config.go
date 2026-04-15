package config

import (
	"fmt"
	"os"
	"encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return config, err
	}

	return config, nil 
}

func (config Config) SetUser(newUserName string) error {
	config, err := Read() // get latest config struct based on current json file config
	if err != nil {
		return err
	}

	config.CurrentUserName = newUserName
	// convert config struct to bytes
	bytes, err := json.Marshal(config) 
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	// write bytes into json file
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf("%s/%s", homePath, configFileName)
	return filePath, nil
}