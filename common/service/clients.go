package service

import (
	"common/pb/episodepb"
	"common/pb/mediapb"
	"common/pb/moviepb"
	"common/pb/seriespb"
	"common/pb/tmdbpb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Service[CFG]) initServiceClients() error {
	err := s.initMediaClient()
	if err != nil {
		return fmt.Errorf("initMediaClient: %w", err)
	}
	err = s.initTMDBClient()
	if err != nil {
		return fmt.Errorf("initTMDBClient: %w", err)
	}
	err = s.initMovieClient()
	if err != nil {
		return fmt.Errorf("initMovieClient: %w", err)
	}
	err = s.initSeriesClient()
	if err != nil {
		return fmt.Errorf("initSeriesClient: %w", err)
	}
	err = s.initEpisodeClient()
	if err != nil {
		return fmt.Errorf("initEpisodeClient: %w", err)
	}
	return nil
}

func (s *Service[CFG]) initMediaClient() error {
	if s.ServiceCfg.MediaServiceAddr == "" {
		return nil
	}
	conn, err := grpc.NewClient(s.ServiceCfg.MediaServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("grpc.NewClient: %w", err)
	}
	s.MediaServiceClient = mediapb.NewMediaServiceClient(conn)
	return nil
}

func (s *Service[CFG]) initTMDBClient() error {
	if s.ServiceCfg.TMDBServiceAddr == "" {
		return nil
	}
	conn, err := grpc.NewClient(s.ServiceCfg.TMDBServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("grpc.NewClient: %w", err)
	}
	s.TMDBServiceClient = tmdbpb.NewTMDBServiceClient(conn)
	return nil
}

func (s *Service[CFG]) initMovieClient() error {
	if s.ServiceCfg.MovieServiceAddr == "" {
		return nil
	}
	conn, err := grpc.NewClient(s.ServiceCfg.MovieServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("grpc.NewClient: %w", err)
	}
	s.MovieServiceClient = moviepb.NewMovieServiceClient(conn)
	return nil
}

func (s *Service[CFG]) initSeriesClient() error {
	if s.ServiceCfg.SeriesServiceAddr == "" {
		return nil
	}
	conn, err := grpc.NewClient(s.ServiceCfg.SeriesServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("grpc.NewClient: %w", err)
	}
	s.SeriesServiceClient = seriespb.NewSeriesServiceClient(conn)
	return nil
}

func (s *Service[CFG]) initEpisodeClient() error {
	if s.ServiceCfg.EpisodeServiceAddr == "" {
		return nil
	}
	conn, err := grpc.NewClient(s.ServiceCfg.EpisodeServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("grpc.NewClient: %w", err)
	}
	s.EpisodeServiceClient = episodepb.NewEpisodeServiceClient(conn)
	return nil
}
