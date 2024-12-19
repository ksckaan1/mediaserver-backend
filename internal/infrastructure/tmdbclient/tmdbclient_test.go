//go:build integration || tmdbclient
// +build integration tmdbclient

package tmdbclient

import (
	"context"
	"os"
	"testing"
)

func TestTMDBClient_GetMovieDetail(t *testing.T) {
	apiKey := os.Getenv("TEST_TMDB_API_KEY")
	tmdbClient, err := New(apiKey)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	result, err := tmdbClient.GetMovieDetail(context.Background(), 603) // The Matrix
	if err != nil {
		t.Fatalf("GetMovieDetail: %v", err)
	}
	t.Logf("result: %v", result.Data)
}

func TestTMDBClient_GetSeriesDetail(t *testing.T) {
	apiKey := os.Getenv("TEST_TMDB_API_KEY")
	tmdbClient, err := New(apiKey)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	result, err := tmdbClient.GetSeriesDetail(context.Background(), 70523) // Dark
	if err != nil {
		t.Fatalf("GetSeriesDetail: %v", err)
	}
	t.Logf("result: %v", result.Data)
}
