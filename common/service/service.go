package service

import (
	"common/configer"
	"common/grpclog"
	"common/idgen"
	"common/logger"
	"common/ports"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Service[CFG any] struct {
	Cfg         *CFG
	Logger      ports.Logger
	GrpcServer  *grpc.Server
	Addr        string
	IDGenerator ports.IDGenerator
	listener    net.Listener
}

func Run[CFG any](ctx context.Context, initializer func(context.Context, *Service[CFG]) error) error {
	cfg := configer.New[CFG]()
	err := cfg.Load()
	if err != nil {
		return fmt.Errorf("configer.Load: %w", err)
	}

	lg, err := logger.New()
	if err != nil {
		return fmt.Errorf("logger.New: %w", err)
	}

	lmw := grpclog.New(lg)
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(lmw.UnaryInterceptor),
		grpc.StreamInterceptor(lmw.StreamInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)

	idGen, err := idgen.New()
	if err != nil {
		return fmt.Errorf("idgen.New: %w", err)
	}

	s := &Service[CFG]{
		Cfg:         &cfg.Data,
		Logger:      lg,
		GrpcServer:  grpcServer,
		IDGenerator: idGen,
	}

	err = initializer(ctx, s)
	if err != nil {
		return fmt.Errorf("initializer: %w", err)
	}

	s.listener, err = net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	go handleGracefulShutdown(s.GrpcServer, s.listener, lg)

	lg.Info(ctx, "service started",
		"addr", s.Addr,
	)

	err = s.GrpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("grpcServer.Serve: %w", err)
	}

	lg.Info(ctx, "service stopped")

	return nil
}

func handleGracefulShutdown(server *grpc.Server, listener net.Listener, lg *logger.Logger) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	lg.Info(context.Background(), "shutting down gracefully")
	server.GracefulStop()
	listener.Close()
}
