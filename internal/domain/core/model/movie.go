package model

import (
	"time"
)

type Movie struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TMDBID      string    `json:"tmdb_id"`
	MediaID     string    `json:"media_id"`
}

type MovieList struct {
	Movies []*Movie `json:"movies"`
	Count  int64    `json:"count"`
	Limit  int64    `json:"limit"`
	Offset int64    `json:"offset"`
}

type GetMovieByIDResponse struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TMDBInfo    *TMDBInfo `json:"tmdb_info"`
	MediaInfo   *Media    `json:"media_info"`
}
