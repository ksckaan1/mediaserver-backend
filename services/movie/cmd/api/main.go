package main

import (
	"context"
	"fmt"
	"movie_service/config"
	"movie_service/internal/core/app"
	"movie_service/internal/infra/repository/couchbasedb"
	"shared/pb/moviepb"
	"shared/searcher"
	"shared/service"
)

func main() {
	ctx := context.Background()

	s := service.NewGRPC(initializer)
	err := s.Run(ctx)
	if err != nil {
		panic(err)
	}
}

func initializer(ctx context.Context, s *service.GRPCService[config.Config]) error {
	err := s.RunCouchbaseQueries(
		ctx,
		"CREATE SCOPE IF NOT EXISTS `media_server`.movie_service;",
		"CREATE COLLECTION IF NOT EXISTS `media_server`.movie_service.movies;",
		"CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.movie_service.movies;",
		"CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.movie_service.movies(id);",
	)
	if err != nil {
		return fmt.Errorf("s.RunCouchbaseQueries: %w", err)
	}
	repository := couchbasedb.New(s.CBBucket)
	src := searcher.New(s.Cfg.TypesenseURL, s.Cfg.TypesenseAPIKey)
	err = src.Migrate(ctx, "movies")
	if err != nil {
		return fmt.Errorf("src.Migrate: %w", err)
	}
	appServer := app.New(
		repository,
		s.IDGenerator,
		s.ServiceClients.MediaServiceClient,
		s.ServiceClients.TMDBServiceClient,
		src,
	)
	moviepb.RegisterMovieServiceServer(s.GrpcServer, appServer)
	return nil
}
