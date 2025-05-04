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
	v1.Mount("/users", initUserRoutes(s.ServiceClients.UserServiceClient, authMW, userTypeMW))
	v1.Mount("/settings", initSettingRoutes(s.ServiceClients.SettingServiceClient, authMW, userTypeMW))
	v1.Use(authMW)
	v1.Mount("/medias", initMediaRoutes(s.ServiceClients.MediaServiceClient, userTypeMW))
	v1.Mount("/movies", initMovieRoutes(s.ServiceClients.MovieServiceClient, userTypeMW))
	v1.Mount("/series", initSeriesRoutes(s.ServiceClients.SeriesServiceClient, userTypeMW))
	v1.Mount("/seasons", initSeasonRoutes(s.ServiceClients.SeasonServiceClient, userTypeMW))
	v1.Mount("/episodes", initEpisodeRoutes(s.ServiceClients.EpisodeServiceClient, userTypeMW))
	return nil
}
