package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port              int    `env:"PORT" envDefault:"8080"`
	TMDBApiKey        string `env:"TMDB_API_KEY"`
	CouchbaseURL      string `env:"COUCHBASE_URL"`
	CouchbaseUser     string `env:"COUCHBASE_USER"`
	CouchbasePassword string `env:"COUCHBASE_PASSWORD"`
	CouchbaseBucket   string `env:"COUCHBASE_BUCKET"`
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
