package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port          int    `env:"PORT" envDefault:"8080"`
	DBPath        string `env:"DB_PATH" envDefault:"./db.sqlite"`
	DBAutoMigrate bool   `env:"DB_AUTO_MIGRATE" envDefault:"false"`
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
