package main

import (
	"bff-service/config"
	"context"
	"shared/service"
)

func main() {
	ctx := context.Background()

	s := service.NewREST(initializer)

	err := s.Run(ctx)
	if err != nil {
		panic(err)
	}
}

func initializer(ctx context.Context, s *service.RESTService[config.Config]) error {
	s.Router.Use(requestIDMW(s.IDGenerator))
	v1 := s.Router.Group("/api/v1")
	authRoutes, authMW, userTypeMW := initAuthRoutes(s.ServiceClients.AuthServiceClient, s.ServiceClients.UserServiceClient)
	v1.Mount("/auth", authRoutes)
	v1.Mount("/user", initUserRoutes(s.ServiceClients.UserServiceClient, authMW, userTypeMW))
	v1.Use(authMW)
	v1.Mount("/media", initMediaRoutes(s.ServiceClients.MediaServiceClient, userTypeMW))
	v1.Mount("/movie", initMovieRoutes(s.ServiceClients.MovieServiceClient, userTypeMW))
	v1.Mount("/series", initSeriesRoutes(s.ServiceClients.SeriesServiceClient, userTypeMW))
	v1.Mount("/season", initSeasonRoutes(s.ServiceClients.SeasonServiceClient, userTypeMW))
	v1.Mount("/episode", initEpisodeRoutes(s.ServiceClients.EpisodeServiceClient, userTypeMW))
	return nil
}
