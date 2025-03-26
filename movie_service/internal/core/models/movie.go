package models

import "time"

type Movie struct {
	ID          string    `bson:"_id" json:"id"`
	CreatedAt   time.Time `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	MediaID     string    `bson:"media_id" json:"media_id"`
	TMDBID      string    `bson:"tmdb_id" json:"tmdb_id"`
}

type MovieList struct {
	List   []*Movie `json:"list"`
	Count  int64    `json:"count"`
	Limit  int64    `json:"limit"`
	Offset int64    `json:"offset"`
}

type MovieSearch struct {
	ID    string `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
}
