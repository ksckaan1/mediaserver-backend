package season

import (
	"bff-service/internal/core/models"
	"context"
	"fmt"
	"net/http"
	"shared/pb/seasonpb"

	"github.com/samber/lo"
)

type ListSeasons struct {
	seasonClient seasonpb.SeasonServiceClient
}

func NewListSeasons(seasonClient seasonpb.SeasonServiceClient) *ListSeasons {
	return &ListSeasons{seasonClient: seasonClient}
}

type ListSeasonsRequest struct {
	SeriesID string `query:"series_id" validate:"required"`
}

type ListSeasonsResponse struct {
	List []*models.Season `json:"list"`
}

func (h *ListSeasons) Handle(ctx context.Context, req *ListSeasonsRequest) (*ListSeasonsResponse, int, error) {
	resp, err := h.seasonClient.ListSeasonsBySeriesID(ctx, &seasonpb.ListSeasonsBySeriesIDRequest{
		SeriesId: req.SeriesID,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("seasonClient.ListSeasons: %w", err)
	}
	return &ListSeasonsResponse{
		List: lo.Map(resp.List, func(s *seasonpb.Season, _ int) *models.Season {
			return &models.Season{
				ID:          s.Id,
				CreatedAt:   s.CreatedAt.AsTime(),
				UpdatedAt:   s.UpdatedAt.AsTime(),
				Title:       s.Title,
				Description: s.Description,
				Order:       s.Order,
				SeriesID:    s.SeriesId,
			}
		}),
	}, http.StatusOK, nil
}
