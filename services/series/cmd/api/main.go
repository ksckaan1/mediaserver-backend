package main

import (
	"context"
	"fmt"
	"series_service/config"
	"series_service/internal/core/app"
	"series_service/internal/infra/repository/couchbasedb"
	"shared/pb/seriespb"
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
	err := src.Migrate(ctx, "series")
	if err != nil {
		return fmt.Errorf("src.Migrate: %w", err)
	}
	appServer := app.New(
		repository,
		s.ServiceClients.TMDBServiceClient,
		s.IDGenerator,
		src,
	)
	seriespb.RegisterSeriesServiceServer(s.GrpcServer, appServer)
	return nil
}
