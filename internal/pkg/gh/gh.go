package gh

import (
	"fmt"
	"net/http"
	"strconv"
)

var ErrNotFound = fmt.Errorf("not found")

func NewRequestContainer[T any](body T, params, queries map[string]string, headers http.Header) *Request[T] {
	return &Request[T]{
		Body:    body,
		Params:  params,
		Headers: headers,
		queries: queries,
	}
}

type Request[T any] struct {
	Body    T
	Params  map[string]string
	Headers http.Header
	queries map[string]string
}

func (r *Request[T]) GetQueryString(key string, defaultValue ...string) (string, error) {
	v, ok := r.queries[key]
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return "", ErrNotFound
	}
	return v, nil
}

func (r *Request[T]) GetQueryInt64(key string, defaultValue ...int64) (int64, error) {
	v, ok := r.queries[key]
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return 0, ErrNotFound
	}
	result, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseInt: %w", err)
	}
	return result, nil
}

func (r *Request[T]) GetQueryFloat64(key string, defaultValue ...float64) (float64, error) {
	v, ok := r.queries[key]
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return 0, ErrNotFound
	}
	result, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	return result, nil
}

func (r *Request[T]) GetQueryBool(key string, defaultValue ...bool) (bool, error) {
	v, ok := r.queries[key]
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return false, ErrNotFound
	}
	result, err := strconv.ParseBool(v)
	if err != nil {
		return false, fmt.Errorf("strconv.ParseBool: %w", err)
	}
	return result, nil
}

type Response[T any] struct {
	Body       T
	StatusCode int
}
