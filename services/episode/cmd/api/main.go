package main

import (
	"context"
	"episode_service/config"
	"episode_service/internal/core/app"
	"episode_service/internal/infra/repository/couchbasedb"
	"shared/pb/episodepb"
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
	repo := couchbasedb.New(s.CBBucket)
	appServer := app.New(repo, s.IDGenerator, s.ServiceClients.MediaServiceClient, s.ServiceClients.SeasonServiceClient)
	episodepb.RegisterEpisodeServiceServer(s.GrpcServer, appServer)
	return nil
}
