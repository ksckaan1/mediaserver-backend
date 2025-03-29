package models

import (
	"time"
)

type Episode struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Order       int32     `json:"order"`
	SeasonID    string    `json:"season_id"`
	MediaInfo   *Media    `json:"media_info"`
}
