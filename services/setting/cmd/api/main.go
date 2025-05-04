package main

import (
	"context"
	"fmt"
	"setting_service/config"
	"setting_service/internal/core/app"
	"setting_service/internal/infra/repository/couchbasedb"
	"shared/pb/settingpb"
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
		"CREATE SCOPE IF NOT EXISTS `media_server`.setting_service;",
		"CREATE COLLECTION IF NOT EXISTS `media_server`.setting_service.settings;",
		"CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.setting_service.settings;",
		"CREATE INDEX IF NOT EXISTS idx_key ON `media_server`.setting_service.settings(`key`);",
	)
	if err != nil {
		return fmt.Errorf("s.RunCouchbaseQueries: %w", err)
	}
	repository := couchbasedb.New(s.CBBucket)
	appServer := app.New(repository)
	settingpb.RegisterSettingServiceServer(s.GrpcServer, appServer)
	return nil
}
