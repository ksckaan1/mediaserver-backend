package main

import (
	"context"
	"season_service/config"
	"season_service/internal/core/app"
	"season_service/internal/infra/repository/couchbasedb"
	"shared/pb/seasonpb"
	"shared/service"
)

func main() {
	ctx := context.Background()
	err := service.Run(ctx, initializer)
	if err != nil {
		panic(err)
	}
}

func initializer(ctx context.Context, s *service.Service[config.Config]) error {
	repository := couchbasedb.New(s.CBBucket)
	appServer := app.New(repository, s.SeriesServiceClient, s.EpisodeServiceClient, s.IDGenerator)
	seasonpb.RegisterSeasonServiceServer(s.GrpcServer, appServer)
	return nil
}
