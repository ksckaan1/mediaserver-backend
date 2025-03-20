package main

import (
	"common/idgen"
	"common/logger"
	"common/pb/mediapb"
	"context"
	"errors"
	"fmt"
	"media_service/config"
	"media_service/database/postgresql/migrations"
	"media_service/internal/domain/core/app"
	"media_service/internal/domain/core/infrastructure/repository"
	"media_service/internal/pkg/s3storage"
	"media_service/internal/port"
	"net"
	"os"
	"os/signal"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	cfg := initConfig()
	lg := initLogger()

	conn := initDatabase(ctx, cfg, lg)
	defer conn.Close(ctx)

	if err := runMigrations(ctx, cfg, lg); err != nil {
		lg.Fatal(ctx, "failed to migrate database", "error", err)
	}

	repo := initRepository(ctx, conn, lg)
	idGen := initIDGenerator(ctx, lg)
	s3Storage := initS3Storage(ctx, cfg, idGen, lg)
	appServer := initAppServer(ctx, repo, s3Storage, idGen, lg)

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

func initDatabase(ctx context.Context, cfg *config.Config, lg *logger.Logger) *pgx.Conn {
	conn, err := pgx.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		lg.Fatal(ctx, "failed to connect to database", "error", err)
	}
	return conn
}

func runMigrations(ctx context.Context, cfg *config.Config, lg port.Logger) error {
	lg.Info(ctx, "database migration starting")

	migrationSource, err := iofs.New(migrations.MigrationsFS, ".")
	if err != nil {
		return fmt.Errorf("iofs.New: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", migrationSource, cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("migrate.NewWithSourceInstance: %w", err)
	}

	if err = m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			lg.Info(ctx, "database migration completed with no change")
			return nil
		}
		return fmt.Errorf("m.Up: %w", err)
	}

	lg.Info(ctx, "database migration completed with change")
	return nil
}

func initRepository(ctx context.Context, conn *pgx.Conn, lg *logger.Logger) app.Repository {
	repo, err := repository.New(conn)
	if err != nil {
		lg.Fatal(ctx, "error when init repository", "error", err)
	}
	return repo
}

func initIDGenerator(ctx context.Context, lg *logger.Logger) port.IDGenerator {
	idGen, err := idgen.New()
	if err != nil {
		lg.Fatal(ctx, "error when init idgen", "error", err)
	}
	return idGen
}

func initS3Storage(ctx context.Context, cfg *config.Config, idGen port.IDGenerator, lg *logger.Logger) port.Storage {
	s3Storage, err := s3storage.New(cfg, idGen)
	if err != nil {
		lg.Fatal(ctx, "error when init s3 storage", "error", err)
	}
	return s3Storage
}

func initAppServer(ctx context.Context, repo app.Repository, storage port.Storage, idGen port.IDGenerator, lg *logger.Logger) mediapb.MediaServiceServer {
	appServer, err := app.New(repo, storage, idGen)
	if err != nil {
		lg.Fatal(ctx, "error when init app server", "error", err)
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

func registerServices(server *grpc.Server, appServer mediapb.MediaServiceServer) {
	mediapb.RegisterMediaServiceServer(server, appServer)
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
