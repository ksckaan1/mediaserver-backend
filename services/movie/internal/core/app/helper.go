package app

import (
	"context"
	"fmt"
	"shared/pb/mediapb"
	"shared/pb/tmdbpb"
)

func (a *App) validateMedia(ctx context.Context, mediaID string) (*mediapb.Media, error) {
	if mediaID == "" {
		return nil, nil
	}
	media, err := a.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{
		MediaId: mediaID,
	})
	if err != nil {
		return nil, fmt.Errorf("mediaClient.GetMediaByID: %w", err)
	}
	return media, nil
}

func (a *App) validateTMDBInfo(ctx context.Context, tmdbID string) (*tmdbpb.TMDBInfo, error) {
	if tmdbID == "" {
		return nil, nil
	}
	tmdbInfo, err := a.tmdbClient.GetTMDBInfo(ctx, &tmdbpb.GetTMDBInfoRequest{
		Id: tmdbID,
	})
	if err != nil {
		return nil, fmt.Errorf("tmdbClient.GetTMDBInfo: %w", err)
	}
	return tmdbInfo, nil
}
