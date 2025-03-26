package main

import (
	"common/pb/mediapb"
	"common/service"
	"context"
	"fmt"
	"media_service/config"
	"media_service/internal/core/app"
	"media_service/internal/infra/repository/couchbasedb"
	"media_service/internal/pkg/s3storage"
)

func main() {
	ctx := context.Background()
	err := service.Run(ctx, initializer)
	if err != nil {
		panic(err)
	}
}

func initializer(ctx context.Context, s *service.Service[config.Config]) error {
	storage, err := s3storage.New(s.Cfg, s.IDGenerator)
	if err != nil {
		return fmt.Errorf("s3storage.New: %w", err)
	}
	repository := couchbasedb.New(s.CBBucket)
	appServer := app.New(repository, storage, s.IDGenerator)
	mediapb.RegisterMediaServiceServer(s.GrpcServer, appServer)
	return nil
}
