package main

import (
	"context"
	"series_service/config"
	"series_service/internal/core/app"
	"series_service/internal/infra/repository/couchbasedb"
	"shared/pb/seriespb"
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
	appServer := app.New(repository, s.ServiceClients.TMDBServiceClient, s.IDGenerator)
	seriespb.RegisterSeriesServiceServer(s.GrpcServer, appServer)
	return nil
}
