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

	s := service.NewGRPC(initializer)
	err := s.Run(ctx)
	if err != nil {
		panic(err)
	}
}

func initializer(ctx context.Context, s *service.GRPCService[config.Config]) error {
	repository := couchbasedb.New(s.CBBucket)
	appServer := app.New(repository, s.ServiceClients.SeriesServiceClient, s.ServiceClients.EpisodeServiceClient, s.IDGenerator)
	seasonpb.RegisterSeasonServiceServer(s.GrpcServer, appServer)
	return nil
}
