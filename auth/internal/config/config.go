package config

import (
	"github.com/himdhiman/unified-framework-go/core/config"
	"github.com/himdhiman/unified-framework-go/core/config/models")

func GetConfiguration(configFilePath string) (*models.Config, error) {
	config, err := config.LoadConfig(configFilePath, nil)
	if err != nil {
		return nil, err
	}

	return config, nil
}
