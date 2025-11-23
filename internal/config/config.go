package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	fileName = ".gatorconfig.json"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, fileName)
	return path, nil
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	content, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUsername = username
	configJson, err := json.Marshal(c)
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, configJson, 0644)
	if err != nil {
		return err
	}

	return nil
}
