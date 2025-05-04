package service

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"shared/configer"
	"shared/idgen"
	"shared/logger"
	"shared/ports"
	"syscall"

	"github.com/couchbase/gocb/v2"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCService[CFG any] struct {
	Cfg         *CFG
	ServiceCfg  *ServiceConfig
	Logger      ports.Logger
	GrpcServer  *grpc.Server
	IDGenerator ports.IDGenerator
	CBBucket    *gocb.Bucket

	ServiceClients *ServiceClients

	// protected fields
	listener net.Listener
	cluster  *gocb.Cluster

	initializer func(context.Context, *GRPCService[CFG]) error
}

func NewGRPC[CFG any](initializer func(context.Context, *GRPCService[CFG]) error) *GRPCService[CFG] {
	lg := logger.New()
	return &GRPCService[CFG]{
		Logger:         lg,
		initializer:    initializer,
		ServiceClients: newServiceClient(lg),
	}
}

func (s *GRPCService[CFG]) Run(ctx context.Context) error {
	serviceCfg := configer.New[ServiceConfig]()
	err := serviceCfg.Load()
	if err != nil {
		return fmt.Errorf("configer[ServiceConfig].Load: %w", err)
	}
	s.ServiceCfg = serviceCfg.Data

	cfg := configer.New[CFG]()
	err = cfg.Load()
	if err != nil {
		return fmt.Errorf("configer[%T].Load: %w", cfg, err)
	}
	s.Cfg = cfg.Data

	err = s.startTracing()
	if err != nil {
		return fmt.Errorf("startTracing: %w", err)
	}

	var grpcOpts []grpc.ServerOption

	if s.ServiceCfg.TracerEndpoint != "" {
		opts := []grpc.ServerOption{
			grpc.ChainUnaryInterceptor(
				otelgrpc.UnaryServerInterceptor(),
			),
			grpc.ChainStreamInterceptor(
				otelgrpc.StreamServerInterceptor(),
			),
		}

		grpcOpts = append(grpcOpts, opts...)
	}

	s.GrpcServer = grpc.NewServer(grpcOpts...)
	reflection.Register(s.GrpcServer)

	s.IDGenerator, err = idgen.New(s.ServiceCfg.IDGeneratorNode)
	if err != nil {
		return fmt.Errorf("idgen.New: %w", err)
	}

	err = s.initCouchbaseBucket(ctx)
	if err != nil {
		return fmt.Errorf("initCouchbaseBucket: %w", err)
	}

	err = s.ServiceClients.initServiceClients(s.ServiceCfg)
	if err != nil {
		return fmt.Errorf("initServiceClients: %w", err)
	}

	s.Logger.Info(ctx, "service initializing")

	err = s.initializer(ctx, s)
	if err != nil {
		return fmt.Errorf("initializer: %w", err)
	}
	s.Logger.Info(ctx, "service initialized")

	s.listener, err = net.Listen("tcp", s.ServiceCfg.Addr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	go s.handleGracefulShutdown()

	s.Logger.Info(ctx, "service started",
		"addr", s.ServiceCfg.Addr,
	)

	err = s.GrpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("grpcServer.Serve: %w", err)
	}

	s.Logger.Info(ctx, "service stopped")

	return nil
}

func (s *GRPCService[CFG]) handleGracefulShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	s.Logger.Info(context.Background(), "shutting down gracefully")
	s.GrpcServer.GracefulStop()
	s.Logger.Info(context.Background(), "grpc server gracefully stopped")
	s.listener.Close()
	s.Logger.Info(context.Background(), "listener closed")
}

func (s *GRPCService[CFG]) RunCouchbaseQueries(ctx context.Context, queries ...string) error {
	for _, query := range queries {
		_, err := s.CBBucket.DefaultScope().Query(query, &gocb.QueryOptions{
			Context: ctx,
		})
		if err != nil {
			return fmt.Errorf("bucket.DefaultScope().Query: %w", err)
		}
	}
	return nil
}
