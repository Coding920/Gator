package config

import (
	"encoding/json"
	"os"
)

func write(c Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(&c, "", "    ")
	if err != nil {
		return err
	}

	stat, err := os.Stat(configPath)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, bytes, stat.Mode())
	if err != nil {
		return err
	}

	return nil
}
