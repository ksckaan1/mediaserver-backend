package main

func (s *Server) linkRoutes() {
	v1 := s.router.Group("/api/v1")

	// MEDIA OPS
	v1.Post("/media", H(s.mediaApp.UploadMedia))
	v1.Get("/media", H(s.mediaApp.ListMedias))
	v1.Get("/media/:id", H(s.mediaApp.GetMediaByID))
	v1.Delete("/media/:id", H(s.mediaApp.DeleteMediaByID))

	// MOVIES
	v1.Post("/movie", H(s.movieApp.CreateMovie))
	v1.Get("/movie", H(s.movieApp.ListMovies))
	v1.Get("/movie/:id", H(s.movieApp.GetMovieByID))
	v1.Put("/movie/:id", H(s.movieApp.UpdateMovieByID))
	v1.Delete("/movie/:id", H(s.movieApp.DeleteMovieByID))

	// SERIES
	v1.Post("/series", H(s.seriesApp.CreateSeries))
	v1.Get("/series", H(s.seriesApp.ListSeries))
	v1.Get("/series/:id", H(s.seriesApp.GetSeriesByID))
	v1.Put("/series/:id", H(s.seriesApp.UpdateSeriesByID))
	v1.Delete("/series/:id", H(s.seriesApp.DeleteSeriesByID))

	// SEASONS
	v1.Post("/season", H(s.seriesApp.CreateSeason))
	v1.Get("/series/:series_id/season", H(s.seriesApp.ListSeasonsBySeriesID))
	v1.Get("/season/:id", H(s.seriesApp.GetSeasonByID))
	v1.Put("/season/:id", H(s.seriesApp.UpdateSeasonByID))
	v1.Delete("/season/:id", H(s.seriesApp.DeleteSeasonByID))

	// EPISODE
	v1.Post("/episode", H(s.seriesApp.CreateEpisode))
	v1.Get("/season/:season_id/episode", H(s.seriesApp.ListEpisodesBySeasonID))
	v1.Get("/episode/:id", H(s.seriesApp.GetEpisodeByID))
	v1.Put("/episode/:id", H(s.seriesApp.UpdateEpisodeByID))
	v1.Delete("/episode/:id", H(s.seriesApp.DeleteEpisodeByID))
}
