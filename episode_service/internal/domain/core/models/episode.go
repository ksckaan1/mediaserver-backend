package models

import "time"

type Episode struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SeasonID    string    `json:"season_id"`
	MediaID     string    `json:"media_id"`
	Order       int32     `json:"order"`
}
