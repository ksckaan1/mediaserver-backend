package main

import (
	"context"
	"fmt"
	"season_service/config"
	"season_service/internal/core/app"
	"season_service/internal/infra/repository/couchbasedb"
	"shared/pb/seasonpb"
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
		"CREATE SCOPE IF NOT EXISTS `media_server`.season_service;",
		"CREATE COLLECTION IF NOT EXISTS `media_server`.season_service.seasons;",
		"CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.season_service.seasons;",
		"CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.season_service.seasons(id);",
		"CREATE INDEX IF NOT EXISTS idx_series_id ON `media_server`.season_service.seasons(series_id);",
	)
	if err != nil {
		return fmt.Errorf("s.RunCouchbaseQueries: %w", err)
	}
	repository := couchbasedb.New(s.CBBucket)
	appServer := app.New(repository, s.ServiceClients.SeriesServiceClient, s.ServiceClients.EpisodeServiceClient, s.IDGenerator)
	seasonpb.RegisterSeasonServiceServer(s.GrpcServer, appServer)
	return nil
}
