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
	repository := couchbasedb.New(s.CBBucket)
	src := searcher.New(s.Cfg.TypesenseURL, s.Cfg.TypesenseAPIKey)
	err := src.Migrate(ctx, "movies")
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
