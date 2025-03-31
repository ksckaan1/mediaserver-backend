package tmdbclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"tmdb_service/internal/core/customerrors"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type TMDBClient struct {
	apiKey string
	client *http.Client
}

func New(apiKey string) (*TMDBClient, error) {
	return &TMDBClient{
		apiKey: apiKey,
		client: &http.Client{
			Timeout:   5 * time.Second,
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		},
	}, nil
}

const movieDetailURL = "https://api.themoviedb.org/3/movie/%s"

func (t *TMDBClient) GetMovieDetail(ctx context.Context, id string) (map[string]any, error) {
	uri := fmt.Sprintf(movieDetailURL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.apiKey))

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("check status code: %w", customerrors.ErrUnexpectedStatusCode{
			StatusCode: resp.StatusCode,
		})
	}

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return result, nil
}

const seriesDetailURL = "https://api.themoviedb.org/3/tv/%s"

func (t *TMDBClient) GetSeriesDetail(ctx context.Context, id string) (map[string]any, error) {
	uri := fmt.Sprintf(seriesDetailURL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.apiKey))

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("check status code: %w", customerrors.ErrUnexpectedStatusCode{
			StatusCode: resp.StatusCode,
		})
	}

	var result map[string]any
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return result, nil
}
