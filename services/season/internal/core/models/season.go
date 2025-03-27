package models

import "time"

type Season struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Order       int32     `json:"order"`
	SeriesID    string    `json:"series_id"`
}
