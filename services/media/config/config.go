package config

type Config struct {
	Port        int    `env:"PORT" envDefault:"8080"`
	S3Endpoint  string `env:"S3_ENDPOINT"`
	S3Region    string `env:"S3_REGION"`
	S3Bucket    string `env:"S3_BUCKET"`
	S3AccessKey string `env:"S3_ACCESS_KEY"`
	S3SecretKey string `env:"S3_SECRET_KEY"`
	S3UseSSL    bool   `env:"S3_USE_SSL"`
}
