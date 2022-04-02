package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

const defaultPath = "./config/config.json"

type (
	// Config -.
	Config struct {
		Redis Redis `json:"redis"`
	}

	// Redis -.
	Redis struct {
		Name     string `json-required:"true" json:"name"`
		Port     string `json-required:"true" json:"port"`
		Password string `json-required:"true" json:"password"`
		Db       int    `json:"db" json-default:"0"`
	}
)

// NewConfig returns app config.
func NewConfig(configPath string) (*Config, error) {

	if configPath == "" {
		configPath = defaultPath
	}
	cfg := &Config{}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
