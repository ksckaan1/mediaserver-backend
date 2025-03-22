package models

import "time"

type Series struct {
	ID          string    `json:"id" bson:"_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	TMDBID      string    `json:"tmdb_id" bson:"tmdb_id"`
}

type SeriesList struct {
	List   []*Series `json:"list" bson:"list"`
	Count  int64     `json:"count" bson:"count"`
	Limit  int64     `json:"limit" bson:"limit"`
	Offset int64     `json:"offset" bson:"offset"`
}

type SeriesSearch struct {
	ID    string `json:"id" bson:"_id"`
	Title string `json:"title" bson:"title"`
}
