package app

import (
	"common/pb/mediapb"
	"context"
	"fmt"
	"movie_service/internal/domain/core/models"

	"github.com/samber/lo"
)

func (m *Movie) validateMedia(ctx context.Context, mediaID string) error {
	if mediaID == "" {
		return nil
	}
	_, err := m.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{
		MediaId: mediaID,
	})
	if err != nil {
		return fmt.Errorf("mediaClient.GetMediaByID: %w", err)
	}
	return nil
}

func (m *Movie) list(ctx context.Context, limit, offset int64, search string) (*models.MovieList, error) {
	if search == "" {
		movies, err := m.repo.ListMovies(ctx, limit, offset)
		if err != nil {
			return nil, fmt.Errorf("repo.ListMovies: %w", err)
		}
		return movies, nil
	}
	moviesSearchList, err := m.repo.ListMoviesSearch(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo.ListMoviesSearch: %w", err)
	}
	filteredMoviesSearchList := m.fuzzySearcher.Search(moviesSearchList, search)
	if len(filteredMoviesSearchList) == 0 {
		return &models.MovieList{
			List:   []*models.Movie{},
			Count:  0,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	ids := lo.Map(filteredMoviesSearchList, func(m *models.MovieSearch, _ int) string {
		return m.ID
	})
	movies, err := m.repo.ListMoviesWithIDs(ctx, ids, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("repo.ListMoviesWithIDs: %w", err)
	}
	return movies, nil
}
