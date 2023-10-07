package services

import (
	"fmt"
	"os"
	"songwhip_bot/models"

	"gopkg.in/yaml.v3"
)

func GetStreamingServices(configFiles ...string) (*models.Services, error) {
	var configFile string
	var services models.Services

	if len(configFiles) > 0 {
		configFile = configFiles[0]
	} else {
		configFile = "configs/services.yaml"
	}

	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open services file: %w", err)
	}

	decoder := yaml.NewDecoder(file)
	if decoder.Decode(&services) != nil {
		return nil, fmt.Errorf("failed to decode services: %w", err)
	}

	return &services, nil
}
