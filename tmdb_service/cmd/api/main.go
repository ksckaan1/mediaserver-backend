package main

import (
	"common/pb/tmdbpb"
	"common/service"
	"context"
	"fmt"
	"time"
	"tmdb_service/config"
	"tmdb_service/internal/domain/core/app"
	"tmdb_service/internal/domain/core/infra/repository/couchbasedb"
	"tmdb_service/internal/domain/core/infra/tmdbclient"

	"github.com/couchbase/gocb/v2"
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

	tmdbClient, err := initTMDBClient(s.Cfg)
	if err != nil {
		return fmt.Errorf("initTMDBClient: %w", err)
	}

	appServer := app.New(tmdbClient, repo)

	tmdbpb.RegisterTMDBServiceServer(s.GrpcServer, appServer)

	s.Logger.Info(ctx, "service initialized")

	return nil
}

func initTMDBClient(cfg *config.Config) (*tmdbclient.TMDBClient, error) {
	tmdbClient, err := tmdbclient.New(cfg.TMDBApiKey)
	if err != nil {
		return nil, fmt.Errorf("tmdbclient.New: %w", err)
	}
	return tmdbClient, nil
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
