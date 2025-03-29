package season

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/seasonpb"
	"strings"
)

type DeleteSeasonByID struct {
	seasonClient seasonpb.SeasonServiceClient
}

func NewDeleteSeasonByID(seasonClient seasonpb.SeasonServiceClient) *DeleteSeasonByID {
	return &DeleteSeasonByID{
		seasonClient: seasonClient,
	}
}

type DeleteSeasonByIDRequest struct {
	SeasonID string `params:"season_id"`
}

func (h *DeleteSeasonByID) Handle(ctx context.Context, req *DeleteSeasonByIDRequest) (*any, int, error) {
	_, err := h.seasonClient.DeleteSeasonByID(ctx, &seasonpb.DeleteSeasonByIDRequest{
		SeasonId: req.SeasonID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "season not found") {
			return nil, http.StatusNotFound, errors.New("season not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("seasonClient.DeleteSeasonByID: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
