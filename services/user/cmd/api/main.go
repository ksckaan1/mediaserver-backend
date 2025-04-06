package main

import (
	"context"
	"shared/password"
	"shared/pb/userpb"
	"shared/service"
	"user_service/config"
	"user_service/internal/core/app"
	"user_service/internal/infra/repository/couchbasedb"
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
	passwordUtil := password.New()
	appServer := app.New(repository, s.IDGenerator, passwordUtil)
	userpb.RegisterUserServiceServer(s.GrpcServer, appServer)
	return nil
}
