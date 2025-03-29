package season

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/seasonpb"
	"strings"
)

type CreateSeason struct {
	seasonClient seasonpb.SeasonServiceClient
}

func NewCreateSeason(seasonClient seasonpb.SeasonServiceClient) *CreateSeason {
	return &CreateSeason{seasonClient: seasonClient}
}

type CreateSeasonRequest struct {
	SeriesId    string `json:"series_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type CreateSeasonResponse struct {
	SeasonID string `json:"season_id"`
}

func (h *CreateSeason) Handle(ctx context.Context, req *CreateSeasonRequest) (*CreateSeasonResponse, int, error) {
	resp, err := h.seasonClient.CreateSeason(ctx, &seasonpb.CreateSeasonRequest{
		SeriesId:    req.SeriesId,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		if strings.Contains(err.Error(), "series not found") {
			return nil, http.StatusNotFound, errors.New("series not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("seasonClient.CreateSeason: %w", err)
	}
	return &CreateSeasonResponse{SeasonID: resp.SeasonId}, http.StatusCreated, nil
}
