package series

import (
	"bff-service/internal/core/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/seriespb"
	"strings"
)

type GetSeriesByID struct {
	seriesClient seriespb.SeriesServiceClient
}

func NewGetSeriesByID(seriesClient seriespb.SeriesServiceClient) *GetSeriesByID {
	return &GetSeriesByID{
		seriesClient: seriesClient,
	}
}

type GetSeriesByIDRequest struct {
	SeriesID string `params:"series_id"`
}

func (h *GetSeriesByID) Handle(ctx context.Context, req *GetSeriesByIDRequest) (*models.Series, int, error) {
	resp, err := h.seriesClient.GetSeriesByID(ctx, &seriespb.GetSeriesByIDRequest{
		SeriesId: req.SeriesID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "series not found") {
			return nil, http.StatusNotFound, errors.New("series not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("seriesClient.GetSeriesByID: %w", err)
	}
	var tmdbInfo *models.TMDB
	if resp.TmdbInfo != nil {
		tmdbInfo = &models.TMDB{
			Id:   resp.TmdbInfo.Id,
			Data: resp.TmdbInfo.Data.AsMap(),
		}
	}
	return &models.Series{
		ID:          resp.Id,
		CreatedAt:   resp.CreatedAt.AsTime(),
		UpdatedAt:   resp.UpdatedAt.AsTime(),
		Title:       resp.Title,
		Description: resp.Description,
		TmdbInfo:    tmdbInfo,
	}, http.StatusOK, nil
}
