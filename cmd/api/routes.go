package main

func (s *Server) linkRoutes() {
	v1 := s.router.Group("/api/v1")

	v1.Post("/movie", H(s.movieApp.CreateMovie))
	v1.Get("/movie/:id", H(s.movieApp.GetMovieByID))
}
