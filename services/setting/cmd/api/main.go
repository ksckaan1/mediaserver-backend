package main

import (
	"context"
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
	repository := couchbasedb.New(s.CBBucket)
	appServer := app.New(repository)
	settingpb.RegisterSettingServiceServer(s.GrpcServer, appServer)
	return nil
}
