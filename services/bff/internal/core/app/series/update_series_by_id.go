package series

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/seriespb"
	"strings"
)

type UpdateSeriesByID struct {
	seriesClient seriespb.SeriesServiceClient
}

func NewUpdateSeriesByID(seriesClient seriespb.SeriesServiceClient) *UpdateSeriesByID {
	return &UpdateSeriesByID{
		seriesClient: seriesClient,
	}
}

type UpdateSeriesByIDRequest struct {
	SeriesID    string `params:"series_id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	TmdbId      string `json:"tmdb_id"`
}

func (h *UpdateSeriesByID) Handle(ctx context.Context, req *UpdateSeriesByIDRequest) (*any, int, error) {
	_, err := h.seriesClient.UpdateSeriesByID(ctx, &seriespb.UpdateSeriesByIDRequest{
		SeriesId:    req.SeriesID,
		Title:       req.Title,
		Description: req.Description,
		TmdbId:      req.TmdbId,
	})
	if err != nil {
		if strings.Contains(err.Error(), "series not found") {
			return nil, http.StatusNotFound, errors.New("series not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("seriesClient.UpdateSeriesByID: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
