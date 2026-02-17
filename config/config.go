package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("no config file was found")
	}

	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("data corrupted")
	}
	return cfg, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(c)

	err = os.WriteFile(path, data, 0666)
	return err
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("no homeDirectory was found")
	}

	path := homeDir + "/.gatorconfig.json"
	return path, nil
}
