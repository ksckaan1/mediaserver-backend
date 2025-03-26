package config

type Config struct {
	Port            int    `env:"PORT" envDefault:"8081"`
	MediaServerHost string `env:"MEDIA_SERVER_HOST"`
	TMDBServerHost  string `env:"TMDB_SERVER_HOST"`
	DatabaseURL     string `env:"DATABASE_URL"`
	DatabaseName    string `env:"DATABASE_NAME"`
}
