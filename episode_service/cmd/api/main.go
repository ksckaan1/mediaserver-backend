package main

import (
	"common/pb/episodepb"
	"common/service"
	"context"
	"episode_service/config"
	"episode_service/internal/core/app"
	"episode_service/internal/infra/repository/couchbasedb"
	"fmt"
)

func main() {
	ctx := context.Background()
	err := service.Run(ctx, initializer)
	if err != nil {
		panic(err)
	}
}

func initializer(ctx context.Context, s *service.Service[config.Config]) error {
	repo, err := couchbasedb.New(s.CBBucket)
	if err != nil {
		return fmt.Errorf("couchbasedb.New: %w", err)
	}
	appServer := app.New(repo, s.IDGenerator, s.MediaServiceClient)
	episodepb.RegisterEpisodeServiceServer(s.GrpcServer, appServer)
	return nil
}
