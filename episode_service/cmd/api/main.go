package main

import (
	"common/pb/episodepb"
	"common/pb/mediapb"
	"common/service"
	"context"
	"episode_service/config"
	"episode_service/internal/core/app"
	"episode_service/internal/infra/repository/couchbasedb"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	err := service.Run(ctx, initializer)
	if err != nil {
		panic(err)
	}
}

func initializer(ctx context.Context, s *service.Service[config.Config]) error {
	s.Addr = fmt.Sprintf(":%d", s.Cfg.Port)

	repo, err := initRepository(ctx, s.Cfg)
	if err != nil {
		return fmt.Errorf("initRepository: %w", err)
	}

	mediaClient, err := initMediaClient(s.Cfg)
	if err != nil {
		return fmt.Errorf("initMediaClient: %w", err)
	}

	appServer := app.New(repo, s.IDGenerator, mediaClient)

	episodepb.RegisterEpisodeServiceServer(s.GrpcServer, appServer)

	s.Logger.Info(ctx, "service initialized")

	return nil
}

func initMediaClient(cfg *config.Config) (mediapb.MediaServiceClient, error) {
	client, err := grpc.NewClient(cfg.MediaServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	return mediapb.NewMediaServiceClient(client), nil
}

func initRepository(ctx context.Context, cfg *config.Config) (*couchbasedb.Repository, error) {
	bucket, err := initCouchbase(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("initCouchbase: %w", err)
	}
	repo, err := couchbasedb.New(bucket)
	if err != nil {
		return nil, fmt.Errorf("couchbasedb.New: %w", err)
	}
	return repo, nil
}

func initCouchbase(ctx context.Context, cfg *config.Config) (*gocb.Bucket, error) {
	cluster, err := gocb.Connect(cfg.CouchbaseURL, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: cfg.CouchbaseUser,
			Password: cfg.CouchbasePassword,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("gocb.Connect: %w", err)
	}

	err = cluster.WaitUntilReady(10*time.Second, &gocb.WaitUntilReadyOptions{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("cluster.WaitUntilReady: %w", err)
	}

	bucket := cluster.Bucket(cfg.CouchbaseBucket)

	return bucket, nil
}
