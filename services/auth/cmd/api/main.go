package main

import (
	"auth_service/config"
	"auth_service/internal/core/app"
	"auth_service/internal/infra/repository/couchbasedb"
	"context"
	"fmt"
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
	err := s.RunCouchbaseQueries(
		ctx,
		"CREATE SCOPE IF NOT EXISTS `media_server`.auth_service;",
		"CREATE COLLECTION IF NOT EXISTS `media_server`.auth_service.sessions;",
		"CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.auth_service.sessions;",
		"CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.auth_service.sessions(session_id);",
		"CREATE INDEX IF NOT EXISTS idx_user_id ON `media_server`.auth_service.sessions(user_id);",
	)
	if err != nil {
		return fmt.Errorf("s.RunCouchbaseQueries: %w", err)
	}
	repository := couchbasedb.New(s.CBBucket)
	appServer := app.New(s.ServiceClients.UserServiceClient, repository, password.New(), s.IDGenerator)
	authpb.RegisterAuthServiceServer(s.GrpcServer, appServer)
	return nil
}
