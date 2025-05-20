package config

import (
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const (
	configFileName = ".gatorconfig.json"
)

func getConfigPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := homePath + "/" + configFileName
	_, err = os.Stat(configPath)
	if err != nil {
		return "", nil
	}

	return configPath, nil
}
