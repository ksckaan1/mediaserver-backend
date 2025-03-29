package models

type TMDB struct {
	Id   string         `json:"id"`
	Data map[string]any `json:"data"`
}
