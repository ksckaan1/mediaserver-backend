package app

import (
	"common/pb/tmdbpb"
	"context"
	"fmt"
	"series_service/internal/domain/core/models"

	"github.com/samber/lo"
)

func (s *SeriesService) validateTMDBInfo(ctx context.Context, tmdbId string) error {
	if tmdbId == "" {
		return nil
	}
	_, err := s.tmdbClient.GetTMDBInfo(ctx, &tmdbpb.GetTMDBInfoRequest{
		Id: tmdbId,
	})
	if err != nil {
		return fmt.Errorf("tmdbClient.GetTMDBInfo: %w", err)
	}
	return nil
}

func (s *SeriesService) list(ctx context.Context, limit, offset int64, search string) (*models.SeriesList, error) {
	if search == "" {
		series, err := s.repo.ListSeries(ctx, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("repo.ListSeries: %w", err)
		}
		return series, nil
	}
	seriesSearchList, err := s.repo.ListSeriesSearch(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo.ListSeriesSearch: %w", err)
	}
	filteredSeriesSearchList := s.fuzzySearcher.Search(seriesSearchList, search)
	if len(filteredSeriesSearchList) == 0 {
		return &models.SeriesList{
			List:   []*models.Series{},
			Count:  0,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	ids := lo.Map(filteredSeriesSearchList, func(s *models.SeriesSearch, _ int) string {
		return s.ID
	})
	series, err := s.repo.ListSeriesWithIDs(ctx, ids, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("repo.ListSeriesWithIDs: %w", err)
	}
	return series, nil
}
