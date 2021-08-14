package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DatabaseUrl string `envconfig:"DATABASE_URL" required:"true"`
}

func LoadConfig() (*Config, error) {
	var cfg *Config
	err := envconfig.Process("", cfg)

	return cfg, err
}
