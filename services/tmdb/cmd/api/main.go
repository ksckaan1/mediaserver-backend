package main

import (
	"context"
	"fmt"
	"shared/pb/tmdbpb"
	"shared/service"
	"tmdb_service/config"
	"tmdb_service/internal/core/app"
	"tmdb_service/internal/infra/repository/couchbasedb"
	"tmdb_service/internal/infra/tmdbclient"
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
	tmdbClient, err := tmdbclient.New(s.Cfg.TMDBApiKey)
	if err != nil {
		return fmt.Errorf("tmdbclient.New: %w", err)
	}
	repo := couchbasedb.New(s.CBBucket)
	appServer := app.New(tmdbClient, repo)
	tmdbpb.RegisterTMDBServiceServer(s.GrpcServer, appServer)
	return nil
}
