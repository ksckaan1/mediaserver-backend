package main

import (
	"common/idgen"
	"common/logger"
	"common/pb/episodepb"
	"common/pb/mediapb"
	"context"
	"episode_service/config"
	"episode_service/internal/domain/core/app"
	"episode_service/internal/domain/core/infra/repository/couchbasedb"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	cfg := initConfig()
	lg := initLogger()

	repo, err := initRepository(ctx, cfg)
	if err != nil {
		lg.Fatal(ctx, "failed to create repository", "error", err)
	}

	idGenerator, err := idgen.New()
	if err != nil {
		lg.Fatal(ctx, "failed to create id generator", "error", err)
	}

	mediaClient, err := initMediaClient(cfg)
	if err != nil {
		lg.Fatal(ctx, "failed to create media client", "error", err)
	}

	appServer := app.New(repo, idGenerator, mediaClient)

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

func initMediaClient(cfg *config.Config) (mediapb.MediaServiceClient, error) {
	client, err := grpc.NewClient(cfg.MediaServerHost)
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	mediaClient := mediapb.NewMediaServiceClient(client)
	return mediaClient, nil
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
		return scope.Name == "episode_service"
	})
	if !ok {
		err := bucket.CollectionsV2().CreateScope("episode_service", &gocb.CreateScopeOptions{
			Context: ctx,
		})
		if err != nil {
			return nil, fmt.Errorf("bucket.CollectionsV2().CreateScope: %w", err)
		}
	}

	// Check collection exists
	if _, ok := lo.Find(targetScope.Collections, func(collection gocb.CollectionSpec) bool {
		return collection.Name == "episodes"
	}); !ok {
		err := bucket.CollectionsV2().CreateCollection("episode_service", "episodes", &gocb.CreateCollectionSettings{}, &gocb.CreateCollectionOptions{
			Context: ctx,
		})
		if err != nil {
			return nil, fmt.Errorf("bucket.CollectionsV2().CreateCollection: %w", err)
		}
	}

	return bucket, nil
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

func registerServices(server *grpc.Server, appServer episodepb.EpisodeServiceServer) {
	episodepb.RegisterEpisodeServiceServer(server, appServer)
	reflection.Register(server)
}

func handleGracefulShutdown(server *grpc.Server, lg *logger.Logger) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	lg.Info(context.Background(), "shutting down gracefully")
	server.GracefulStop()
}
