package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port        int    `env:"PORT" envDefault:"8080"`
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://media_service:media_service@postgres:5432/media_service?sslmode=disable"`
	S3Endpoint  string `env:"S3_ENDPOINT" envDefault:"localhost:9000"`
	S3Region    string `env:"S3_REGION" envDefault:"eu-central-1"`
	S3Bucket    string `env:"S3_BUCKET" envDefault:"media"`
	S3AccessKey string `env:"S3_ACCESS_KEY" envDefault:"minioadmin"`
	S3SecretKey string `env:"S3_SECRET_KEY" envDefault:"minioadmin"`
	S3UseSSL    bool   `env:"S3_USE_SSL" envDefault:"false"`
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
