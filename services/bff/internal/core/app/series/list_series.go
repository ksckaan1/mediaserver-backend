package series

import (
	"bff-service/internal/core/models"
	"context"
	"fmt"
	"net/http"
	"shared/pb/seriespb"

	"github.com/samber/lo"
)

type ListSeries struct {
	seriesClient seriespb.SeriesServiceClient
}

func NewListSeries(seriesClient seriespb.SeriesServiceClient) *ListSeries {
	return &ListSeries{
		seriesClient: seriesClient,
	}
}

type ListSeriesRequest struct {
	Limit  int64 `query:"limit"`
	Offset int64 `query:"offset"`
}

type ListSeriesResponse struct {
	List   []*models.Series `json:"list"`
	Count  int64            `json:"count"`
	Limit  int64            `json:"limit"`
	Offset int64            `json:"offset"`
}

func (h *ListSeries) Handle(ctx context.Context, req *ListSeriesRequest) (*ListSeriesResponse, int, error) {
	resp, err := h.seriesClient.ListSeries(ctx, &seriespb.ListSeriesRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("seriesClient.ListSeries: %w", err)
	}
	return &ListSeriesResponse{
		List: lo.Map(resp.List, func(s *seriespb.Series, _ int) *models.Series {
			var tmdbInfo *models.TMDB
			if s.TmdbInfo != nil {
				tmdbInfo = &models.TMDB{
					Id:   s.TmdbInfo.Id,
					Data: s.TmdbInfo.Data.AsMap(),
				}
			}
			return &models.Series{
				ID:          s.Id,
				CreatedAt:   s.CreatedAt.AsTime(),
				UpdatedAt:   s.UpdatedAt.AsTime(),
				Title:       s.Title,
				Description: s.Description,
				TmdbInfo:    tmdbInfo,
			}
		}),
		Count:  resp.Count,
		Limit:  resp.Limit,
		Offset: resp.Offset,
	}, http.StatusOK, nil
}
