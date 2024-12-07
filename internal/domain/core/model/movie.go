package model

import "time"

type Movie struct {
	ID          string
	CreatedAt   time.Time
	Title       string
	Description string
	TMDBID      string
}
