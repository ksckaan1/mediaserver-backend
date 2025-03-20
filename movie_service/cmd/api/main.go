package main

import (
	"common/idgen"
	"common/logger"
	"common/pb/mediapb"
	"common/pb/moviepb"
	"context"
	"fmt"
	"movie_service/config"
	"movie_service/internal/domain/core/app"
	"movie_service/internal/infrastructure/repository/mongodb"
	"movie_service/internal/port"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	cfg := initConfig()
	lg := initLogger()
	
	mongoClient := initMongoDB(ctx, cfg, lg)
	defer mongoClient.Disconnect(ctx)
	
	idGenerator := initIDGenerator(ctx, lg)
	grpcClient := initGRPCClient(ctx, cfg, lg)
	defer grpcClient.Close()
	
	repo := initRepository(mongoClient)
	mediaClient := mediapb.NewMediaServiceClient(grpcClient)
	appServer := initAppServer(ctx, repo, idGenerator, mediaClient, lg)
	
	server := initGRPCServer(lg)
	listener := initListener(ctx, cfg, lg)
	
	registerServices(server, appServer)
	
	go handleGracefulShutdown(server, lg)
	
	lg.Info(ctx, "server starting", "port", cfg.Port)
	if err := server.Serve(listener); err != nil {
		lg.Fatal(ctx, "failed to serve", "error", err)
	}
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

func initMongoDB(ctx context.Context, cfg *config.Config, lg *logger.Logger) *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.DatabaseURL).SetServerAPIOptions(serverAPI))
	if err != nil {
		lg.Fatal(ctx, "failed to connect to database", "error", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		lg.Fatal(ctx, "failed to ping database", "error", err)
	}
	lg.Info(ctx, "connected to database")
	return client
}

func initIDGenerator(ctx context.Context, lg *logger.Logger) port.IDGenerator {
	idGenerator, err := idgen.New()
	if err != nil {
		lg.Fatal(ctx, "failed to create id generator", "error", err)
	}
	return idGenerator
}

func initGRPCClient(ctx context.Context, cfg *config.Config, lg *logger.Logger) *grpc.ClientConn {
	grpcClient, err := grpc.Dial(cfg.MediaServerHost, grpc.WithInsecure())
	if err != nil {
		lg.Fatal(ctx, "failed to create grpc client", "error", err)
	}
	return grpcClient
}

func initRepository(client *mongo.Client) *mongodb.Repository {
	db := client.Database("movie_service")
	return mongodb.New(db)
}

func initAppServer(ctx context.Context, repo *mongodb.Repository, idGenerator port.IDGenerator, mediaClient mediapb.MediaServiceClient, lg *logger.Logger) moviepb.MovieServiceServer {
	appServer, err := app.New(repo, idGenerator, mediaClient)
	if err != nil {
		lg.Fatal(ctx, "failed to create app server", "error", err)
	}
	return appServer
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

func initListener(ctx context.Context, cfg *config.Config, lg *logger.Logger) net.Listener {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		lg.Fatal(ctx, "failed to listen", "error", err)
	}
	return listener
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
