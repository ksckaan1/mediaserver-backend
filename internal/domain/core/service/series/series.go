package series

type Series struct{}

func New() (*Series, error) {
	return &Series{}, nil
}
