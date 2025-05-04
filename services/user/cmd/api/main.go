package main

import (
	"context"
	"fmt"
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
	err := s.RunCouchbaseQueries(
		ctx,
		"CREATE SCOPE IF NOT EXISTS `media_server`.user_service;",
		"CREATE COLLECTION IF NOT EXISTS `media_server`.user_service.users;",
		"CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.user_service.users;",
		"CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.user_service.users(id);",
		"CREATE INDEX IF NOT EXISTS idx_username ON `media_server`.user_service.users(username);",
	)
	if err != nil {
		return fmt.Errorf("s.RunCouchbaseQueries: %w", err)
	}
	repository := couchbasedb.New(s.CBBucket)
	passwordUtil := password.New()
	appServer := app.New(repository, s.IDGenerator, passwordUtil)
	userpb.RegisterUserServiceServer(s.GrpcServer, appServer)
	return nil
}
