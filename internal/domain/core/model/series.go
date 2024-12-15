package model

import "time"

type Series struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TMDBID      int64     `json:"tmdb_id"`
}

type SeriesList struct {
	List   []*Series `json:"list"`
	Count  int64     `json:"count"`
	Limit  int64     `json:"limit"`
	Offset int64     `json:"offset"`
}

type Season struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SeriesID    string    `json:"series_id"`
	Order       int64     `json:"order"`
}

type SeasonList struct {
	List   []*Season `json:"list"`
	Count  int64     `json:"count"`
	Limit  int64     `json:"limit"`
	Offset int64     `json:"offset"`
}

type Episode struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SeasonID    string    `json:"season_id"`
	Order       int64     `json:"order"`
	MediaID     string    `json:"media_id"`
}

type EpisodeList struct {
	List   []*Episode `json:"list"`
	Count  int64      `json:"count"`
	Limit  int64      `json:"limit"`
	Offset int64      `json:"offset"`
}
