package main

import (
	"common/logger"
	"common/pb/tmdbpb"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
	"tmdb_service/config"
	"tmdb_service/internal/domain/core/app"
	"tmdb_service/internal/domain/core/infra/repository/couchbasedb"
	"tmdb_service/internal/domain/core/infra/tmdbclient"

	"github.com/couchbase/gocb/v2"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	cfg := initConfig()
	lg := initLogger()

	tmdbClient, err := initTMDBClient(cfg)
	if err != nil {
		lg.Fatal(ctx, "failed to create tmdb client", "error", err)
	}

	repo, err := initRepository(ctx, cfg)
	if err != nil {
		lg.Fatal(ctx, "failed to create repository", "error", err)
	}

	appServer, err := initAppServer(tmdbClient, repo)
	if err != nil {
		lg.Fatal(ctx, "failed to create app server", "error", err)
	}

	server := initGRPCServer(lg)

	listener, err := initListener(cfg)
	if err != nil {
		lg.Fatal(ctx, "failed to create listener", "error", err)
	}

	registerServices(server, appServer)

	go handleGracefulShutdown(server, lg)

	lg.Info(ctx, "server starting", "port", cfg.Port)

	err = server.Serve(listener)
	if err != nil {
		lg.Fatal(ctx, "failed to serve", "error", err)
	}

	lg.Info(ctx, "graceful shutdown successful")
}

func initConfig() *config.Config {
	cfg := config.New()
	if err := cfg.Load(); err != nil {
		panic(err)
	}
	return cfg
}

func initLogger() *logger.Logger {
	lg, err := logger.New()
	if err != nil {
		panic(err)
	}
	return lg
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

	// Wait for initialization
	err = cluster.WaitUntilReady(10*time.Second, &gocb.WaitUntilReadyOptions{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("cluster.WaitUntilReady: %w", err)
	}

	// Check bucket exists
	buckerManager := cluster.Buckets()
	buckets, err := buckerManager.GetAllBuckets(&gocb.GetAllBucketsOptions{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("bucketManager.GetAllBuckets: %w", err)
	}
	if _, ok := buckets[cfg.CouchbaseBucket]; !ok {
		err := buckerManager.CreateBucket(gocb.CreateBucketSettings{
			BucketSettings: gocb.BucketSettings{
				Name: cfg.CouchbaseBucket,
			},
		}, &gocb.CreateBucketOptions{
			Context: ctx,
		})
		if err != nil {
			return nil, fmt.Errorf("bucketManager.CreateBucket: %w", err)
		}
	}

	// Check scope exists
	bucket := cluster.Bucket(cfg.CouchbaseBucket)
	scopes, err := bucket.CollectionsV2().GetAllScopes(&gocb.GetAllScopesOptions{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("bucket.CollectionsV2().GetAllScopes: %w", err)
	}

	targetScope, ok := lo.Find(scopes, func(scope gocb.ScopeSpec) bool {
		return scope.Name == "tmdb_service"
	})
	if !ok {
		err := bucket.CollectionsV2().CreateScope("tmdb_service", &gocb.CreateScopeOptions{
			Context: ctx,
		})
		if err != nil {
			return nil, fmt.Errorf("bucket.CollectionsV2().CreateScope: %w", err)
		}
	}

	// Check collection exists
	if _, ok := lo.Find(targetScope.Collections, func(collection gocb.CollectionSpec) bool {
		return collection.Name == "infos"
	}); !ok {
		err := bucket.CollectionsV2().CreateCollection("tmdb_service", "infos", &gocb.CreateCollectionSettings{}, &gocb.CreateCollectionOptions{
			Context: ctx,
		})
		if err != nil {
			return nil, fmt.Errorf("bucket.CollectionsV2().CreateCollection: %w", err)
		}
	}

	return bucket, nil
}

func initAppServer(tmdbClient *tmdbclient.TMDBClient, repo *couchbasedb.Repository) (tmdbpb.TMDBServiceServer, error) {
	appServer, err := app.New(tmdbClient, repo)
	if err != nil {
		return nil, fmt.Errorf("app.New: %w", err)
	}
	return appServer, nil
}

func initGRPCServer(lg *logger.Logger) *grpc.Server {
	lmw := &loggerMiddleWare{
		logger: lg,
	}
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(lmw.unaryInterceptor),
		grpc.StreamInterceptor(lmw.streamInterceptor),
	}
	return grpc.NewServer(opts...)
}

func initListener(cfg *config.Config) (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return nil, fmt.Errorf("net.Listen: %w", err)
	}
	return listener, nil
}

func registerServices(server *grpc.Server, appServer tmdbpb.TMDBServiceServer) {
	tmdbpb.RegisterTMDBServiceServer(server, appServer)
	reflection.Register(server)
}

func handleGracefulShutdown(server *grpc.Server, lg *logger.Logger) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	lg.Info(context.Background(), "shutting down gracefully")
	server.GracefulStop()
}
