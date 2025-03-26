package service

import (
	"common/configer"
	"common/grpclog"
	"common/idgen"
	"common/logger"
	"common/pb/episodepb"
	"common/pb/mediapb"
	"common/pb/moviepb"
	"common/pb/seriespb"
	"common/pb/tmdbpb"
	"common/ports"
	"context"
	"fmt"
	"net"

	"github.com/couchbase/gocb/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Service[CFG any] struct {
	Cfg         *CFG
	ServiceCfg  *ServiceConfig
	Logger      ports.Logger
	GrpcServer  *grpc.Server
	IDGenerator ports.IDGenerator
	CBBucket    *gocb.Bucket

	// Clients
	MediaServiceClient   mediapb.MediaServiceClient
	TMDBServiceClient    tmdbpb.TMDBServiceClient
	MovieServiceClient   moviepb.MovieServiceClient
	SeriesServiceClient  seriespb.SeriesServiceClient
	EpisodeServiceClient episodepb.EpisodeServiceClient

	// protected fields
	listener net.Listener
	cluster  *gocb.Cluster
}

func Run[CFG any](ctx context.Context, initializer func(context.Context, *Service[CFG]) error) error {
	s := new(Service[CFG])

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

	s.Logger, err = logger.New()
	if err != nil {
		return fmt.Errorf("logger.New: %w", err)
	}

	lmw := grpclog.New(s.Logger)
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(lmw.UnaryInterceptor),
		grpc.StreamInterceptor(lmw.StreamInterceptor),
	}

	s.GrpcServer = grpc.NewServer(opts...)
	reflection.Register(s.GrpcServer)

	s.IDGenerator, err = idgen.New(s.ServiceCfg.IDGeneratorNode)
	if err != nil {
		return fmt.Errorf("idgen.New: %w", err)
	}

	err = s.initCouchbaseBucket(ctx)
	if err != nil {
		return fmt.Errorf("initCouchbaseBucket: %w", err)
	}

	err = s.initServiceClients()
	if err != nil {
		return fmt.Errorf("initServiceClients: %w", err)
	}

	s.Logger.Info(ctx, "service initializing")

	err = initializer(ctx, s)
	if err != nil {
		return fmt.Errorf("initializer: %w", err)
	}
	s.Logger.Info(ctx, "service initialized")

	s.listener, err = net.Listen("tcp", s.ServiceCfg.Addr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	go handleGracefulShutdown(s.GrpcServer, s.listener, s.Logger)

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
