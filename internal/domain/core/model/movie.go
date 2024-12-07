package model

import "time"

type Movie struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TMDBID      string    `json:"tmdb_id"`
}

type MovieList struct {
	Movies []*Movie `json:"movies"`
	Count  int64    `json:"count"`
	Limit  int64    `json:"limit"`
	Offset int64    `json:"offset"`
}
