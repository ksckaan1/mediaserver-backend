package service

import (
	"context"
	"fmt"
	"shared/pb/authpb"
	"shared/pb/episodepb"
	"shared/pb/mediapb"
	"shared/pb/moviepb"
	"shared/pb/seasonpb"
	"shared/pb/seriespb"
	"shared/pb/settingpb"
	"shared/pb/tmdbpb"
	"shared/pb/userpb"
	"shared/ports"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClients struct {
	AuthServiceClient    authpb.AuthServiceClient
	UserServiceClient    userpb.UserServiceClient
	SettingServiceClient settingpb.SettingServiceClient
	MediaServiceClient   mediapb.MediaServiceClient
	TMDBServiceClient    tmdbpb.TMDBServiceClient
	MovieServiceClient   moviepb.MovieServiceClient
	SeriesServiceClient  seriespb.SeriesServiceClient
	SeasonServiceClient  seasonpb.SeasonServiceClient
	EpisodeServiceClient episodepb.EpisodeServiceClient
	logger               ports.Logger
}

func newServiceClient(logger ports.Logger) *ServiceClients {
	return &ServiceClients{
		logger: logger,
	}
}

func (s *ServiceClients) initServiceClients(cfg *ServiceConfig) error {
	err := s.initAuthService(cfg)
	if err != nil {
		return fmt.Errorf("initAuthService: %w", err)
	}
	err = s.initUserClient(cfg)
	if err != nil {
		return fmt.Errorf("initUserClient: %w", err)
	}
	err = s.initSettingService(cfg)
	if err != nil {
		return fmt.Errorf("initSettingService: %w", err)
	}
	err = s.initMediaClient(cfg)
	if err != nil {
		return fmt.Errorf("initMediaClient: %w", err)
	}
	err = s.initTMDBClient(cfg)
	if err != nil {
		return fmt.Errorf("initTMDBClient: %w", err)
	}
	err = s.initMovieClient(cfg)
	if err != nil {
		return fmt.Errorf("initMovieClient: %w", err)
	}
	err = s.initSeriesClient(cfg)
	if err != nil {
		return fmt.Errorf("initSeriesClient: %w", err)
	}
	err = s.initSeasonClient(cfg)
	if err != nil {
		return fmt.Errorf("initSeasonClient: %w", err)
	}
	err = s.initEpisodeClient(cfg)
	if err != nil {
		return fmt.Errorf("initEpisodeClient: %w", err)
	}
	return nil
}

func (s *ServiceClients) initAuthService(cfg *ServiceConfig) error {
	if cfg.AuthServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.AuthServiceAddr)
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.AuthServiceClient = authpb.NewAuthServiceClient(conn)

	s.logger.Info(context.Background(), "auth service client initialized")

	return nil
}

func (s *ServiceClients) initSettingService(cfg *ServiceConfig) error {
	if cfg.SettingServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.SettingServiceAddr)
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.SettingServiceClient = settingpb.NewSettingServiceClient(conn)

	s.logger.Info(context.Background(), "setting service client initialized")

	return nil
}

func (s *ServiceClients) initUserClient(cfg *ServiceConfig) error {
	if cfg.UserServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.UserServiceAddr)
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.UserServiceClient = userpb.NewUserServiceClient(conn)

	s.logger.Info(context.Background(), "user service client initialized")

	return nil
}

func (s *ServiceClients) initMediaClient(cfg *ServiceConfig) error {
	if cfg.MediaServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.MediaServiceAddr)
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.MediaServiceClient = mediapb.NewMediaServiceClient(conn)

	s.logger.Info(context.Background(), "media service client initialized")

	return nil
}

func (s *ServiceClients) initTMDBClient(cfg *ServiceConfig) error {
	if cfg.TMDBServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.TMDBServiceAddr)
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.TMDBServiceClient = tmdbpb.NewTMDBServiceClient(conn)

	s.logger.Info(context.Background(), "tmdb service client initialized")

	return nil
}

func (s *ServiceClients) initMovieClient(cfg *ServiceConfig) error {
	if cfg.MovieServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.MovieServiceAddr)
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.MovieServiceClient = moviepb.NewMovieServiceClient(conn)

	s.logger.Info(context.Background(), "movie service client initialized")

	return nil
}

func (s *ServiceClients) initSeriesClient(cfg *ServiceConfig) error {
	if cfg.SeriesServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.SeriesServiceAddr)
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.SeriesServiceClient = seriespb.NewSeriesServiceClient(conn)

	s.logger.Info(context.Background(), "series service client initialized")

	return nil
}

func (s *ServiceClients) initSeasonClient(cfg *ServiceConfig) error {
	if cfg.SeasonServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.SeasonServiceAddr)
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.SeasonServiceClient = seasonpb.NewSeasonServiceClient(conn)

	s.logger.Info(context.Background(), "season service client initialized")

	return nil
}

func (s *ServiceClients) initEpisodeClient(cfg *ServiceConfig) error {
	if cfg.EpisodeServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.EpisodeServiceAddr)
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.EpisodeServiceClient = episodepb.NewEpisodeServiceClient(conn)

	s.logger.Info(context.Background(), "episode service client initialized")

	return nil
}

func (s *ServiceClients) initClient(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor(
			otelgrpc.WithSpanOptions(trace.WithSpanKind(trace.SpanKindClient)),
		)),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor(
			otelgrpc.WithSpanOptions(trace.WithSpanKind(trace.SpanKindClient)),
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	conn.Connect()
	return conn, nil
}
