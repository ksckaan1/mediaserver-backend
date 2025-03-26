package main

import (
	"common/configer"
	"common/grpclog"
	"common/idgen"
	"common/logger"
	"common/pb/mediapb"
	"common/ports"
	"context"
	"errors"
	"fmt"
	"media_service/config"
	"media_service/database/postgresql/migrations"
	"media_service/internal/core/app"
	"media_service/internal/infra/repository"
	"media_service/internal/pkg/s3storage"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	cfg, err := initConfig()
	if err != nil {
		panic(err)
	}

	lg, err := logger.New()
	if err != nil {
		panic(err)
	}

	conn, err := initDatabase(ctx, cfg)
	if err != nil {
		lg.Fatal(ctx, "initDatabase",
			"error", err,
		)
	}
	defer conn.Close(ctx)

	err = runMigrations(ctx, cfg, lg)
	if err != nil {
		lg.Fatal(ctx, "runMigrations",
			"error", err,
		)
	}

	repo, err := repository.New(conn)
	if err != nil {
		lg.Fatal(ctx, "repository.New",
			"error", err,
		)
	}

	idGen, err := idgen.New()
	if err != nil {
		lg.Fatal(ctx, "idgen.New",
			"error", err,
		)
	}

	s3Storage, err := s3storage.New(cfg, idGen)
	if err != nil {
		lg.Fatal(ctx, "s3storage.New",
			"error", err,
		)
	}

	appServer := app.New(repo, s3Storage, idGen)

	server := initGRPCServer(lg)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Data.Port))
	if err != nil {
		lg.Fatal(ctx, "net.Listen",
			"error", err,
		)
	}

	registerServices(server, appServer)

	go handleGracefulShutdown(server, lg)

	lg.Info(ctx, "server starting", "port", cfg.Data.Port)
	err = server.Serve(listener)
	if err != nil {
		lg.Fatal(ctx, "server.Serve",
			"error", err,
		)
	}
}

func initConfig() (*configer.Configer[config.Config], error) {
	cfg := configer.New[config.Config]()
	err := cfg.Load()
	if err != nil {
		return nil, fmt.Errorf("configer.Load: %w", err)
	}
	return cfg, nil
}

func initDatabase(ctx context.Context, cfg *configer.Configer[config.Config]) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, cfg.Data.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect: %w", err)
	}
	return conn, nil
}

func runMigrations(ctx context.Context, cfg *configer.Configer[config.Config], lg ports.Logger) error {
	lg.Info(ctx, "database migration starting")

	migrationSource, err := iofs.New(migrations.MigrationsFS, ".")
	if err != nil {
		return fmt.Errorf("iofs.New: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", migrationSource, cfg.Data.DatabaseURL)
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

func initGRPCServer(lg ports.Logger) *grpc.Server {
	lmw := grpclog.New(lg)
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(lmw.UnaryInterceptor),
		grpc.StreamInterceptor(lmw.StreamInterceptor),
	}
	return grpc.NewServer(opts...)
}

func registerServices(server *grpc.Server, appServer mediapb.MediaServiceServer) {
	mediapb.RegisterMediaServiceServer(server, appServer)
	reflection.Register(server)
}

func handleGracefulShutdown(server *grpc.Server, lg ports.Logger) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	lg.Info(context.Background(), "initiating graceful shutdown")
	server.GracefulStop()
	lg.Info(context.Background(), "server shutdown complete")
}
