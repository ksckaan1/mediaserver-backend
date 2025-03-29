package season

import (
	"bff-service/internal/core/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/seasonpb"
	"strings"
)

type GetSeasonByID struct {
	seasonClient seasonpb.SeasonServiceClient
}

func NewGetSeasonByID(seasonClient seasonpb.SeasonServiceClient) *GetSeasonByID {
	return &GetSeasonByID{
		seasonClient: seasonClient,
	}
}

type GetSeasonByIDRequest struct {
	SeasonID string `json:"season_id" validate:"required"`
}

func (h *GetSeasonByID) Handle(ctx context.Context, req *GetSeasonByIDRequest) (*models.Season, int, error) {
	resp, err := h.seasonClient.GetSeasonByID(ctx, &seasonpb.GetSeasonByIDRequest{
		SeasonId: req.SeasonID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "season not found") {
			return nil, http.StatusNotFound, errors.New("season not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("seasonClient.GetSeasonByID: %w", err)
	}
	return &models.Season{
		ID:          resp.Id,
		CreatedAt:   resp.CreatedAt.AsTime(),
		UpdatedAt:   resp.UpdatedAt.AsTime(),
		Title:       resp.Title,
		Description: resp.Description,
		Order:       resp.Order,
		SeriesID:    resp.SeriesId,
	}, http.StatusOK, nil
}
