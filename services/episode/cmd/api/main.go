package main

import (
	"context"
	"episode_service/config"
	"episode_service/internal/core/app"
	"episode_service/internal/infra/repository/couchbasedb"
	"fmt"
	"shared/pb/episodepb"
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
		"CREATE SCOPE IF NOT EXISTS `media_server`.episode_service;",
		"CREATE COLLECTION IF NOT EXISTS `media_server`.episode_service.episodes;",
		"CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.episode_service.episodes;",
		"CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.episode_service.episodes(id);",
		"CREATE INDEX IF NOT EXISTS idx_season_id ON `media_server`.episode_service.episodes(season_id);",
	)
	if err != nil {
		return fmt.Errorf("s.RunCouchbaseQueries: %w", err)
	}
	repo := couchbasedb.New(s.CBBucket)
	appServer := app.New(repo, s.IDGenerator, s.ServiceClients.MediaServiceClient, s.ServiceClients.SeasonServiceClient)
	episodepb.RegisterEpisodeServiceServer(s.GrpcServer, appServer)
	return nil
}
