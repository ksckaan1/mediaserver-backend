package models

import "time"

type Series struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TmdbInfo    *TMDB     `json:"tmdb_info"`
	Tags        []string  `json:"tags"`
}
