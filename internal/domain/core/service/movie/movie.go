package movie

type Movie struct{}

func New() (*Movie, error) {
	return &Movie{}, nil
}
