package model

type TMDBData map[string]any

type TMDBInfo struct {
	ID   string   `json:"id"`
	Data TMDBData `json:"data"`
}
