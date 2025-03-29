package main

import (
	"fmt"
	"shared/pb/episodepb"
	"shared/pb/mediapb"
	"shared/pb/moviepb"
	"shared/pb/seasonpb"
	"shared/pb/seriespb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewMediaServiceClient(addr string) (mediapb.MediaServiceClient, error) {
	client, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	return mediapb.NewMediaServiceClient(client), nil
}

func NewMovieServiceClient(addr string) (moviepb.MovieServiceClient, error) {
	client, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	return moviepb.NewMovieServiceClient(client), nil
}

func NewSeriesServiceClient(addr string) (seriespb.SeriesServiceClient, error) {
	client, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	return seriespb.NewSeriesServiceClient(client), nil
}

func NewSeasonServiceClient(addr string) (seasonpb.SeasonServiceClient, error) {
	client, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	return seasonpb.NewSeasonServiceClient(client), nil
}

func NewEpisodeServiceClient(addr string) (episodepb.EpisodeServiceClient, error) {
	client, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}
	return episodepb.NewEpisodeServiceClient(client), nil
}
