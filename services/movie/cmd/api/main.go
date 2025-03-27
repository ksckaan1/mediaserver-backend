package main

import (
	"context"
	"movie_service/config"
	"movie_service/internal/core/app"
	"movie_service/internal/infra/repository/couchbasedb"
	"shared/pb/moviepb"
	"shared/service"
)

func main() {
	ctx := context.Background()
	err := service.Run(ctx, initializer)
	if err != nil {
		panic(err)
	}
}

func initializer(ctx context.Context, s *service.Service[config.Config]) error {
	repo := couchbasedb.New(s.CBBucket)
	appServer := app.New(repo, s.IDGenerator, s.MediaServiceClient, s.TMDBServiceClient)
	moviepb.RegisterMovieServiceServer(s.GrpcServer, appServer)
	return nil
}
