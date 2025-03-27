package models

import (
	"time"
)

type TMDBInfo struct {
	Id        string         `json:"id"`
	Data      map[string]any `json:"data"`
	UpdatedAt time.Time      `json:"updated_at"`
}
