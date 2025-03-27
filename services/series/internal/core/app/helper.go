package app

import (
	"context"
	"fmt"
	"shared/pb/tmdbpb"
)

func (s *App) validateTMDBInfo(ctx context.Context, tmdbId string) error {
	if tmdbId == "" {
		return nil
	}
	_, err := s.tmdbClient.GetTMDBInfo(ctx, &tmdbpb.GetTMDBInfoRequest{
		Id: tmdbId,
	})
	if err != nil {
		return fmt.Errorf("tmdbClient.GetTMDBInfo: %w", err)
	}
	return nil
}
