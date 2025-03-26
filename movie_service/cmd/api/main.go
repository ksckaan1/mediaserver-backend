package main

import (
	"common/configer"
	"common/grpclog"
	"common/idgen"
	"common/logger"
	"common/pb/mediapb"
	"common/pb/moviepb"
	"common/pb/tmdbpb"
	"common/ports"
	"context"
	"fmt"
	"movie_service/config"
	"movie_service/internal/domain/core/app"
	"movie_service/internal/infrastructure/repository/mongodb"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	cfg, err := initConfig()
	if err != nil {
		panic(err)
	}

	lg, err := initLogger()
	if err != nil {
		panic(err)
	}

	mongoClient, err := initMongoDB(ctx, cfg)
	if err != nil {
		lg.Fatal(ctx, "initMongoDB", "error", err)
	}
	defer mongoClient.Disconnect(ctx)

	idGenerator, err := initIDGenerator()
	if err != nil {
		lg.Fatal(ctx, "initIDGenerator", "error", err)
	}

	mediaClient, err := initMediaClient(cfg)
	if err != nil {
		lg.Fatal(ctx, "initMediaClient", "error", err)
	}

	tmdbClient, err := initTMDBClient(cfg)
	if err != nil {
		lg.Fatal(ctx, "initTMDBClient", "error", err)
	}

	repo := initRepository(mongoClient, cfg)

	appServer, err := app.New(repo, idGenerator, mediaClient, tmdbClient)
	if err != nil {
		lg.Fatal(ctx, "app.New", "error", err)
	}

	server := initGRPCServer(lg)

	listener, err := initListener(cfg)
	if err != nil {
		lg.Fatal(ctx, "initListener", "error", err)
	}

	registerServices(server, appServer)

	go handleGracefulShutdown(server, lg)

	lg.Info(ctx, "server starting", "port", cfg.Data.Port)

	err = server.Serve(listener)
	if err != nil {
		lg.Fatal(ctx, "server.Serve", "error", err)
	}
}

func initConfig() (*configer.Configer[config.Config], error) {
	cfg := configer.New[config.Config]()
	if err := cfg.Load(); err != nil {
		return nil, fmt.Errorf("configer.New: %w", err)
	}
	return cfg, nil
}

func initLogger() (*logger.Logger, error) {
	lg, err := logger.New()
	if err != nil {
		return nil, fmt.Errorf("logger.New: %w", err)
	}
	return lg, nil
}

func initMongoDB(ctx context.Context, cfg *configer.Configer[config.Config]) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.Data.DatabaseURL).SetServerAPIOptions(serverAPI))
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect: %w", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo.Ping: %w", err)
	}
	return client, nil
}

func initIDGenerator() (ports.IDGenerator, error) {
	idGenerator, err := idgen.New()
	if err != nil {
		return nil, fmt.Errorf("idgen.New: %w", err)
	}
	return idGenerator, nil
}

func initMediaClient(cfg *configer.Configer[config.Config]) (mediapb.MediaServiceClient, error) {
	grpcClient, err := grpc.NewClient(cfg.Data.MediaServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	return mediapb.NewMediaServiceClient(grpcClient), nil
}

func initTMDBClient(cfg *configer.Configer[config.Config]) (tmdbpb.TMDBServiceClient, error) {
	grpcClient, err := grpc.NewClient(cfg.Data.TMDBServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	return tmdbpb.NewTMDBServiceClient(grpcClient), nil
}

func initRepository(client *mongo.Client, cfg *configer.Configer[config.Config]) *mongodb.Repository {
	db := client.Database(cfg.Data.DatabaseName)
	return mongodb.New(db)
}

func initGRPCServer(lg *logger.Logger) *grpc.Server {
	lmw := grpclog.New(lg)
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(lmw.UnaryInterceptor),
		grpc.StreamInterceptor(lmw.StreamInterceptor),
	}
	return grpc.NewServer(opts...)
}

func initListener(cfg *configer.Configer[config.Config]) (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Data.Port))
	if err != nil {
		return nil, fmt.Errorf("net.Listen: %w", err)
	}
	return listener, nil
}

func registerServices(server *grpc.Server, appServer moviepb.MovieServiceServer) {
	moviepb.RegisterMovieServiceServer(server, appServer)
	reflection.Register(server)
}

func handleGracefulShutdown(server *grpc.Server, lg *logger.Logger) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	lg.Info(context.Background(), "initiating graceful shutdown")
	server.GracefulStop()
	lg.Info(context.Background(), "server shutdown complete")
}
