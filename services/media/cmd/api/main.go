package main

import (
	"context"
	"fmt"
	"media_service/config"
	"media_service/internal/core/app"
	"media_service/internal/infra/repository/couchbasedb"
	"media_service/internal/pkg/s3storage"
	"shared/pb/mediapb"
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
		"CREATE SCOPE IF NOT EXISTS `media_server`.media_service;",
		"CREATE COLLECTION IF NOT EXISTS `media_server`.media_service.medias;",
		"CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.media_service.medias;",
		"CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.media_service.medias(id);",
	)
	if err != nil {
		return fmt.Errorf("s.RunCouchbaseQueries: %w", err)
	}
	storage, err := s3storage.New(s.Cfg, s.IDGenerator)
	if err != nil {
		return fmt.Errorf("s3storage.New: %w", err)
	}
	repository := couchbasedb.New(s.CBBucket)
	appServer := app.New(repository, storage, s.IDGenerator)
	mediapb.RegisterMediaServiceServer(s.GrpcServer, appServer)
	return nil
}
