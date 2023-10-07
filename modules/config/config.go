package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func GetConfig(configFiles ...string) (*AppConfig, error) {
	var configFile string

	if len(configFiles) > 0 {
		configFile = configFiles[0]
	} else {
		configFile = "configs/config.yaml"
	}

	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	config := &AppConfig{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	return config, nil
}
