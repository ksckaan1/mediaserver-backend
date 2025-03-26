package app

import (
	"common/pb/mediapb"
	"common/pb/tmdbpb"
	"context"
	"fmt"
)

func (a *App) validateMedia(ctx context.Context, mediaID string) error {
	if mediaID == "" {
		return nil
	}
	_, err := a.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{
		MediaId: mediaID,
	})
	if err != nil {
		return fmt.Errorf("mediaClient.GetMediaByID: %w", err)
	}
	return nil
}

func (a *App) validateTMDBInfo(ctx context.Context, tmdbID string) error {
	if tmdbID == "" {
		return nil
	}
	_, err := a.tmdbClient.GetTMDBInfo(ctx, &tmdbpb.GetTMDBInfoRequest{
		Id: tmdbID,
	})
	if err != nil {
		return fmt.Errorf("tmdbClient.GetTMDBInfo: %w", err)
	}
	return nil
}
