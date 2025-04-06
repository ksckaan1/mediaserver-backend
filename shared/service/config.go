package service

type ServiceConfig struct {
	ServiceName     string `env:"SERVICE_NAME" envDefault:"noname-service"`
	TracerEndpoint  string `env:"TRACER_ENDPOINT"`
	Addr            string `env:"ADDR" envDefault:":8080"`
	IDGeneratorNode int64  `env:"ID_GENERATOR_NODE" envDefault:"1"`

	// Couchbase
	CouchbaseHost     string `env:"COUCHBASE_HOST"`
	CouchbaseUser     string `env:"COUCHBASE_USER"`
	CouchbasePassword string `env:"COUCHBASE_PASSWORD"`
	CouchbaseBucket   string `env:"COUCHBASE_BUCKET"`

	// Clients
	MediaServiceAddr   string `env:"MEDIA_SERVICE_ADDR"`
	TMDBServiceAddr    string `env:"TMDB_SERVICE_ADDR"`
	MovieServiceAddr   string `env:"MOVIE_SERVICE_ADDR"`
	SeriesServiceAddr  string `env:"SERIES_SERVICE_ADDR"`
	SeasonServiceAddr  string `env:"SEASON_SERVICE_ADDR"`
	EpisodeServiceAddr string `env:"EPISODE_SERVICE_ADDR"`
	UserServiceAddr    string `env:"USER_SERVICE_ADDR"`
	AuthServiceAddr    string `env:"AUTH_SERVICE_ADDR"`
}
