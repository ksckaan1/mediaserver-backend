package models

import "time"

type Movie struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	MediaID     string    `json:"media_id"`
	TMDBID      string    `json:"tmdb_id"`
	Tags        []string  `json:"tags"`
}

type MovieList struct {
	List   []*Movie `json:"list"`
	Count  int64    `json:"count"`
	Limit  int64    `json:"limit"`
	Offset int64    `json:"offset"`
}
