package season

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/seasonpb"
	"strings"
)

type UpdateSeasonByID struct {
	seasonClient seasonpb.SeasonServiceClient
}

func NewUpdateSeasonByID(seasonClient seasonpb.SeasonServiceClient) *UpdateSeasonByID {
	return &UpdateSeasonByID{
		seasonClient: seasonClient,
	}
}

type UpdateSeasonByIDRequest struct {
	SeasonID    string `params:"season_id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

func (h *UpdateSeasonByID) Handle(ctx context.Context, req *UpdateSeasonByIDRequest) (*any, int, error) {
	_, err := h.seasonClient.UpdateSeasonByID(ctx, &seasonpb.UpdateSeasonByIDRequest{
		SeasonId:    req.SeasonID,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		if strings.Contains(err.Error(), "season not found") {
			return nil, http.StatusNotFound, errors.New("season not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("seasonClient.UpdateSeasonByID: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
