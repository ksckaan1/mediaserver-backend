package model

type TMDBInfo struct {
	ID            int64   `json:"id"`
	OriginalTitle string  `json:"original_title"`
	PosterPath    string  `json:"poster_path"`
	BackdropPath  string  `json:"backdrop_path"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int64   `json:"vote_count"`
	Popularity    float64 `json:"popularity"`
	ReleaseDate   string  `json:"release_date"`
}
