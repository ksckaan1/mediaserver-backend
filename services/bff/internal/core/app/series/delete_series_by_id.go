package series

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/seriespb"
	"strings"
)

type DeleteSeriesByID struct {
	seriesClient seriespb.SeriesServiceClient
}

func NewDeleteSeriesByID(seriesClient seriespb.SeriesServiceClient) *DeleteSeriesByID {
	return &DeleteSeriesByID{
		seriesClient: seriesClient,
	}
}

type DeleteSeriesByIDRequest struct {
	SeriesID string `params:"series_id"`
}

func (h *DeleteSeriesByID) Handle(ctx context.Context, req *DeleteSeriesByIDRequest) (*any, int, error) {
	_, err := h.seriesClient.DeleteSeriesByID(ctx, &seriespb.DeleteSeriesByIDRequest{
		SeriesId: req.SeriesID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "series not found") {
			return nil, http.StatusNotFound, errors.New("series not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("seriesClient.DeleteSeriesByID: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
