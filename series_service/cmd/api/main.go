package main

import (
	"common/idgen"
	"common/logger"
	"common/pb/seriespb"
	"common/pb/tmdbpb"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"series_service/config"
	"series_service/internal/domain/core/app"
	"series_service/internal/infrastructure/repository/mongodb"
	"series_service/internal/port"

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

	tmdbClient, err := initTMDBClient(cfg)
	if err != nil {
		lg.Fatal(ctx, "initTMDBClient", "error", err)
	}

	repo := initRepository(mongoClient, cfg)

	appServer, err := app.New(repo, tmdbClient, idGenerator)
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

	lg.Info(ctx, "server starting", "port", cfg.Port)

	err = server.Serve(listener)
	if err != nil {
		lg.Fatal(ctx, "server.Serve", "error", err)
	}
}

func initConfig() (*config.Config, error) {
	cfg := config.New()
	if err := cfg.Load(); err != nil {
		return nil, fmt.Errorf("config.Load: %w", err)
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

func initMongoDB(ctx context.Context, cfg *config.Config) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	fmt.Println("cfg.DatabaseURL", cfg.DatabaseURL)
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.DatabaseURL).SetServerAPIOptions(serverAPI))
	if err != nil {
		return nil, fmt.Errorf("mongo.Connect: %w", err)
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo.Ping: %w", err)
	}
	return client, nil
}

func initIDGenerator() (port.IDGenerator, error) {
	idGenerator, err := idgen.New()
	if err != nil {
		return nil, fmt.Errorf("idgen.New: %w", err)
	}
	return idGenerator, nil
}

func initTMDBClient(cfg *config.Config) (tmdbpb.TMDBServiceClient, error) {
	grpcClient, err := grpc.NewClient(cfg.TMDBServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	return tmdbpb.NewTMDBServiceClient(grpcClient), nil
}

func initRepository(client *mongo.Client, cfg *config.Config) *mongodb.Repository {
	db := client.Database(cfg.DatabaseName)
	return mongodb.New(db)
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

func registerServices(server *grpc.Server, appServer seriespb.SeriesServiceServer) {
	seriespb.RegisterSeriesServiceServer(server, appServer)
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
