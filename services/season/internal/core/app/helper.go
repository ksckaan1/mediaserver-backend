package app

import (
	"context"
	"fmt"
	"season_service/internal/core/models"

	"github.com/samber/lo"
)

func (a *App) generateOrderNumber(ctx context.Context, seriesID string) (int32, error) {
	if seriesID == "" {
		return 1, nil
	}
	seasons, err := a.repository.ListSeasonsBySeriesID(ctx, seriesID)
	if err != nil {
		return 0, fmt.Errorf("repository.ListSeasonsBySeriesID: %w", err)
	}
	ids := lo.Map(seasons, func(e *models.Season, _ int) int32 {
		return e.Order
	})
	return lo.Max(ids) + 1, nil
}
