package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

const defaultPath = "./config/config.yml"
type (
	// Config -.
	Config struct {
		Redis  `json:"redis"`
	}

	// Redis -.
	Redis struct {
		Name string `json-required:"true" json:"name"`
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

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
