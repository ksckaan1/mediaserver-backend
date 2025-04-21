package series

import (
	"bff-service/internal/core/models"
	"context"
	"fmt"
	"net/http"
	"shared/pb/seriespb"

	"github.com/samber/lo"
)

type SearchSeries struct {
	seriesClient seriespb.SeriesServiceClient
}

func NewSearchSeries(seriesClient seriespb.SeriesServiceClient) *SearchSeries {
	return &SearchSeries{
		seriesClient: seriesClient,
	}
}

type SearchSeriesRequest struct {
	Limit   *int64 `query:"limit"`
	Offset  int64  `query:"offset"`
	Query   string `query:"query" validate:"required"`
	QueryBy string `query:"query_by" validate:"required,oneof=title tags"`
}

type SearchSeriesResponse struct {
	List   []*models.Series `json:"list"`
	Count  int64            `json:"count"`
	Limit  int64            `json:"limit"`
	Offset int64            `json:"offset"`
}

func (h *SearchSeries) Handle(ctx context.Context, req *SearchSeriesRequest) (*SearchSeriesResponse, int, error) {
	var limit int64 = 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	resp, err := h.seriesClient.SearchSeries(ctx, &seriespb.SearchSeriesRequest{
		Query:   req.Query,
		QueryBy: req.QueryBy,
		Limit:   limit,
		Offset:  req.Offset,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("seriesClient.SearchSeries: %w", err)
	}
	return &SearchSeriesResponse{
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
				Tags:        s.Tags,
			}
		}),
		Count:  resp.Count,
		Limit:  resp.Limit,
		Offset: resp.Offset,
	}, http.StatusOK, nil
}
