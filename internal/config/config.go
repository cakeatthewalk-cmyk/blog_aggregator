package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	jsonFilePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Could not get user home: %w", err)
	}
	jsonFilePath = filepath.Join(jsonFilePath, configFileName)
	return jsonFilePath, nil
}

func Read() (Config, error) {
	jsonFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Could not get config file %w", err)
	}
	fileContents, err := os.ReadFile(jsonFilePath)

	if err != nil {
		return Config{}, fmt.Errorf("Could not read file: %w", err)
	}

	cfg := Config{}

	err = json.Unmarshal(fileContents, &cfg)

	if err != nil {
		return Config{}, fmt.Errorf("Could not unmarshal json! (poor Jason): %w", err)
	}

	return cfg, nil
}

func write(cfg *Config) error {
	pathToConfig, err := getConfigFilePath()
	if err != nil {
		return err
	}
	jsonCfg, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Error occured: %w", err)
	}
	err = os.WriteFile(pathToConfig, jsonCfg, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	err := write(c)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	return nil
}
