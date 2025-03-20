package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port            int    `env:"PORT" envDefault:"8081"`
	MediaServerHost string `env:"MEDIA_SERVER_HOST" envDefault:"localhost:8080"`
	DatabaseURL     string `env:"DATABASE_URL" envDefault:"mongodb://localhost:27017"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) Load() error {
	err := env.Parse(c)
	if err != nil {
		return fmt.Errorf("env.Parse: %w", err)
	}
	return nil
}
