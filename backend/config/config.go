package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseUrl string `json:"database-url" envconfig:"DATABASE_URL" required:"true"`
	Port        int    `json:"port" envconfig:"PORT" required:"true"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)

	return &cfg, err
}
