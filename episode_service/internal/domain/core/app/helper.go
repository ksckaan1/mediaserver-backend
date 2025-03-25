package app

import (
	"common/pb/mediapb"
	"context"
	"episode_service/internal/domain/core/models"
	"fmt"

	"github.com/samber/lo"
)

func (h *Episode) validateMediaID(ctx context.Context, mediaID string) error {
	if mediaID == "" {
		return nil
	}
	_, err := h.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{
		MediaId: mediaID,
	})
	if err != nil {
		return fmt.Errorf("mediaClient.GetMediaByID: %w", err)
	}
	return nil
}

func (h *Episode) generateOrderNumber(ctx context.Context, seasonID string) (int32, error) {
	if seasonID == "" {
		return 1, nil
	}
	episodes, err := h.repository.ListEpisodesBySeasonID(ctx, seasonID)
	if err != nil {
		return 0, fmt.Errorf("repository.ListEpisodesBySeasonID: %w", err)
	}
	ids := lo.Map(episodes, func(e *models.Episode, _ int) int32 {
		return e.Order
	})
	return lo.Max(ids) + 1, nil
}
