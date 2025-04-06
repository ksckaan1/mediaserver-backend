package main

import (
	"auth_service/config"
	"auth_service/internal/core/app"
	"auth_service/internal/infra/repository/couchbasedb"
	"context"
	"shared/password"
	"shared/pb/authpb"
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
	appServer := app.New(s.ServiceClients.UserServiceClient, repository, password.New(), s.IDGenerator)
	authpb.RegisterAuthServiceServer(s.GrpcServer, appServer)
	return nil
}
