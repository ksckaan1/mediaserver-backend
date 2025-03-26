package main

import (
	"common/pb/seriespb"
	"common/service"
	"context"
	"series_service/config"
	"series_service/internal/core/app"
	"series_service/internal/infra/repository/couchbasedb"
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
	appServer := app.New(repository, s.TMDBServiceClient, s.IDGenerator)
	seriespb.RegisterSeriesServiceServer(s.GrpcServer, appServer)
	return nil
}
