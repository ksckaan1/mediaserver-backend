package service

import (
	"fmt"
	"shared/pb/episodepb"
	"shared/pb/mediapb"
	"shared/pb/moviepb"
	"shared/pb/seasonpb"
	"shared/pb/seriespb"
	"shared/pb/tmdbpb"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClients struct {
	MediaServiceClient   mediapb.MediaServiceClient
	TMDBServiceClient    tmdbpb.TMDBServiceClient
	MovieServiceClient   moviepb.MovieServiceClient
	SeriesServiceClient  seriespb.SeriesServiceClient
	SeasonServiceClient  seasonpb.SeasonServiceClient
	EpisodeServiceClient episodepb.EpisodeServiceClient
}

func (s *ServiceClients) initServiceClients(cfg *ServiceConfig) error {
	err := s.initMediaClient(cfg)
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

func (s *ServiceClients) initMediaClient(cfg *ServiceConfig) error {
	if cfg.MediaServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.MediaServiceAddr, "media")
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.MediaServiceClient = mediapb.NewMediaServiceClient(conn)
	return nil
}

func (s *ServiceClients) initTMDBClient(cfg *ServiceConfig) error {
	if cfg.TMDBServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.TMDBServiceAddr, "tmdb")
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.TMDBServiceClient = tmdbpb.NewTMDBServiceClient(conn)
	return nil
}

func (s *ServiceClients) initMovieClient(cfg *ServiceConfig) error {
	if cfg.MovieServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.MovieServiceAddr, "movie")
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.MovieServiceClient = moviepb.NewMovieServiceClient(conn)
	return nil
}

func (s *ServiceClients) initSeriesClient(cfg *ServiceConfig) error {
	if cfg.SeriesServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.SeriesServiceAddr, "series")
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.SeriesServiceClient = seriespb.NewSeriesServiceClient(conn)
	return nil
}

func (s *ServiceClients) initSeasonClient(cfg *ServiceConfig) error {
	if cfg.SeasonServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.SeasonServiceAddr, "season")
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.SeasonServiceClient = seasonpb.NewSeasonServiceClient(conn)
	return nil
}

func (s *ServiceClients) initEpisodeClient(cfg *ServiceConfig) error {
	if cfg.EpisodeServiceAddr == "" {
		return nil
	}

	conn, err := s.initClient(cfg.EpisodeServiceAddr, "episode")
	if err != nil {
		return fmt.Errorf("initClient: %w", err)
	}

	s.EpisodeServiceClient = episodepb.NewEpisodeServiceClient(conn)
	return nil
}

func (s *ServiceClients) initClient(addr, name string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor(
			otelgrpc.WithSpanAttributes(
				attribute.String("service.name", name+"_client"),
			),
			otelgrpc.WithSpanOptions(trace.WithSpanKind(trace.SpanKindClient)),
		)),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor(
			otelgrpc.WithSpanAttributes(
				attribute.String("service.name", name+"_client"),
			),
			otelgrpc.WithSpanOptions(trace.WithSpanKind(trace.SpanKindClient)),
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	conn.Connect()
	return conn, nil
}
