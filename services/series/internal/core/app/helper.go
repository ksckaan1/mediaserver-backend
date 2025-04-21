package app

import (
	"context"
	"fmt"
	"shared/pb/tmdbpb"
)

func (s *App) validateTMDBInfo(ctx context.Context, tmdbId string) (*tmdbpb.TMDBInfo, error) {
	if tmdbId == "" {
		return nil, nil
	}
	tmdbInfo, err := s.tmdbClient.GetTMDBInfo(ctx, &tmdbpb.GetTMDBInfoRequest{
		Id: tmdbId,
	})
	if err != nil {
		return nil, fmt.Errorf("tmdbClient.GetTMDBInfo: %w", err)
	}
	return tmdbInfo, nil
}
