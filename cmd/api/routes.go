package main

func (s *Server) linkRoutes() {
	v1 := s.router.Group("/api/v1")

	// MOVIES
	v1.Post("/movie", H(s.movieApp.CreateMovie))
	v1.Get("/movie", H(s.movieApp.ListMovies))
	v1.Get("/movie/:id", H(s.movieApp.GetMovieByID))
	v1.Put("/movie/:id", H(s.movieApp.UpdateMovieByID))
	v1.Delete("/movie/:id", H(s.movieApp.DeleteMovieByID))

	// MEDIA OPS
	v1.Post("/media", H(s.mediaApp.UploadMedia))
	v1.Get("/media", H(s.mediaApp.ListMedias))
	v1.Get("/media/:id", H(s.mediaApp.GetMediaByID))
	v1.Delete("/media/:id", H(s.mediaApp.DeleteMediaByID))
}
