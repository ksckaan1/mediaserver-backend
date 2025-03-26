package config

type Config struct {
	Port              int    `env:"PORT" envDefault:"8080"`
	MediaServerHost   string `env:"MEDIA_SERVER_HOST"`
	CouchbaseURL      string `env:"COUCHBASE_URL"`
	CouchbaseUser     string `env:"COUCHBASE_USER"`
	CouchbasePassword string `env:"COUCHBASE_PASSWORD"`
	CouchbaseBucket   string `env:"COUCHBASE_BUCKET"`
}
