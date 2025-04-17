package series

import (
	"context"
	"fmt"
	"net/http"
	"shared/pb/seriespb"
)

type CreateSeries struct {
	seriesClient seriespb.SeriesServiceClient
}

func NewCreateSeries(seriesClient seriespb.SeriesServiceClient) *CreateSeries {
	return &CreateSeries{
		seriesClient: seriesClient,
	}
}

type CreateSeriesRequest struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description"`
	TmdbId      string   `json:"tmdb_id"`
	Tags        []string `json:"tags"`
}

type CreateSeriesResponse struct {
	SeriesID string `json:"series_id"`
}

func (h *CreateSeries) Handle(ctx context.Context, req *CreateSeriesRequest) (*CreateSeriesResponse, int, error) {
	resp, err := h.seriesClient.CreateSeries(ctx, &seriespb.CreateSeriesRequest{
		Title:       req.Title,
		Description: req.Description,
		TmdbId:      req.TmdbId,
		Tags:        req.Tags,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("seriesClient.CreateSeries: %w", err)
	}
	return &CreateSeriesResponse{
		SeriesID: resp.SeriesId,
	}, http.StatusCreated, nil
}
