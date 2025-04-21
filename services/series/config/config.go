package config

type Config struct {
	TypesenseURL    string `env:"TYPESENSE_URL"`
	TypesenseAPIKey string `env:"TYPESENSE_API_KEY"`
}
