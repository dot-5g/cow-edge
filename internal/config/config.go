package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	NodeID    string `yaml:"nodeID"`
	Interface string `yaml:"interface"`
}

func ReadConfig(configPath string) (*Config, error) {
	var config Config

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
