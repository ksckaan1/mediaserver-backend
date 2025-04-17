package models

import "time"

type Series struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TMDBID      string    `json:"tmdb_id"`
	Tags        []string  `json:"tags"`
}

type SeriesList struct {
	List   []*Series `json:"list"`
	Count  int64     `json:"count"`
	Limit  int64     `json:"limit"`
	Offset int64     `json:"offset"`
}
