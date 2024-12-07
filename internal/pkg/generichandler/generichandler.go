package generichandler

import "net/http"

type GenericHandler struct{}

func New() (*GenericHandler, error) {
	return &GenericHandler{}, nil
}

type Request[T any] struct {
	Body    T
	Params  map[string]string
	Headers http.Header
	Queries map[string]string
}

type Response[T any] struct {
	Body       T
	StatusCode int
}
