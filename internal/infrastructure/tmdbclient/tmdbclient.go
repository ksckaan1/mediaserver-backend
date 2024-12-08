package tmdbclient

import (
	"context"
	"encoding/json"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"net/http"
	"time"
)

type TMDBClient struct {
	apiKey string
	client *http.Client
}

func New(apiKey string) (*TMDBClient, error) {
	return &TMDBClient{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}, nil
}

const movieDetailURL = "https://api.themoviedb.org/3/movie/%d"

func (t *TMDBClient) GetMovieDetail(ctx context.Context, id int64) (*model.TMDBInfo, error) {
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
		return nil, fmt.Errorf("check status code: %d", customerrors.ErrUnexpectedStatusCode{
			StatusCode: resp.StatusCode,
		})
	}

	var result model.TMDBInfo
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	return &result, nil
}
