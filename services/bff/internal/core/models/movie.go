package models

import "time"

type Movie struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	MediaInfo   *Media    `json:"media_info"`
	TmdbInfo    *TMDB     `json:"tmdb_info"`
}
