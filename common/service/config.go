package service

type ServiceConfig struct {
	Addr            string `env:"ADDR" envDefault:":8080"`
	IDGeneratorNode int64  `env:"ID_GENERATOR_NODE" envDefault:"1"`
	// Couchbase
	CouchbaseURL      string `env:"COUCHBASE_URL"`
	CouchbaseUser     string `env:"COUCHBASE_USER"`
	CouchbasePassword string `env:"COUCHBASE_PASSWORD"`
	CouchbaseBucket   string `env:"COUCHBASE_BUCKET"`

	// Clients
	MediaServiceAddr   string `env:"MEDIA_SERVICE_ADDR"`
	TMDBServiceAddr    string `env:"TMDB_SERVICE_ADDR"`
	MovieServiceAddr   string `env:"MOVIE_SERVICE_ADDR"`
	SeriesServiceAddr  string `env:"SERIES_SERVICE_ADDR"`
	EpisodeServiceAddr string `env:"EPISODE_SERVICE_ADDR"`
}
