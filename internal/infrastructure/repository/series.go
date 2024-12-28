package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/infrastructure/repository/sqlcgen"

	"github.com/samber/lo"
)

func (m *Repository) CreateSeries(ctx context.Context, series *model.Series) error {
	err := m.queries.CreateSeries(ctx, sqlcgen.CreateSeriesParams{
		ID:          series.ID,
		Title:       series.Title,
		Description: series.Description,
		TmdbID:      series.TMDBID,
	})
	if err != nil {
		return fmt.Errorf("queries.CreateSeries: %w", err)
	}
	return nil
}

func (m *Repository) ListSeries(ctx context.Context, limit, offset int64) (*model.SeriesList, error) {
	count, err := m.queries.CountSeries(ctx)
	if err != nil {
		return nil, fmt.Errorf("queries.CountSeries: %w", err)
	}
	if count == 0 {
		return &model.SeriesList{
			List:   make([]*model.Series, 0),
			Count:  count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	list, err := m.queries.ListSeries(ctx, sqlcgen.ListSeriesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("queries.ListSeries: %w", err)
	}
	return &model.SeriesList{
		List: lo.Map(list, func(s sqlcgen.Series, _ int) *model.Series {
			return &model.Series{
				ID:          s.ID,
				CreatedAt:   s.CreatedAt,
				UpdatedAt:   s.UpdatedAt,
				Title:       s.Title,
				Description: s.Description,
				TMDBID:      s.TmdbID,
			}
		}),
		Count:  count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (m *Repository) GetSeriesByID(ctx context.Context, id string) (*model.Series, error) {
	series, err := m.queries.GetSeriesByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("queries.GetSeriesByID: %w", customerrors.ErrSeriesNotFound)
		}
		return nil, fmt.Errorf("queries.GetSeriesByID: %w", err)
	}
	return &model.Series{
		ID:          series.ID,
		CreatedAt:   series.CreatedAt,
		UpdatedAt:   series.UpdatedAt,
		Title:       series.Title,
		Description: series.Description,
		TMDBID:      series.TmdbID,
	}, nil
}

func (m *Repository) UpdateSeriesByID(ctx context.Context, series *model.Series) error {
	_, err := m.queries.UpdateSeriesByID(ctx, sqlcgen.UpdateSeriesByIDParams{
		ID:          series.ID,
		Title:       series.Title,
		Description: series.Description,
		TmdbID:      series.TMDBID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.UpdateSeriesByID: %w", customerrors.ErrSeriesNotFound)
		}
		return fmt.Errorf("queries.UpdateSeriesByID: %w", err)
	}
	return nil
}

func (m *Repository) DeleteSeriesByID(ctx context.Context, id string) error {
	_, err := m.queries.DeleteSeriesByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.DeleteSeriesByID: %w", customerrors.ErrSeriesNotFound)
		}
		return fmt.Errorf("queries.DeleteSeriesByID: %w", err)
	}
	return nil
}

func (m *Repository) CreateSeason(ctx context.Context, season *model.Season) error {
	err := m.queries.CreateSeason(ctx, sqlcgen.CreateSeasonParams{
		ID:          season.ID,
		Title:       season.Title,
		Description: season.Description,
		SeriesID:    season.SeriesID,
		Order:       season.Order,
	})
	if err != nil {
		return fmt.Errorf("queries.CreateSeason: %w", err)
	}
	return nil
}

func (m *Repository) ListSeasonsBySeriesID(ctx context.Context, seriesID string, limit, offset int64) (*model.SeasonList, error) {
	count, err := m.queries.CountSeasonsBySeriesID(ctx, seriesID)
	if err != nil {
		return nil, fmt.Errorf("queries.CountSeasonsBySeriesID: %w", err)
	}
	if count == 0 {
		return &model.SeasonList{
			List:   make([]*model.Season, 0),
			Count:  count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	list, err := m.queries.ListSeasonsBySeriesID(ctx, sqlcgen.ListSeasonsBySeriesIDParams{
		SeriesID: seriesID,
		Limit:    limit,
		Offset:   offset,
	})
	if err != nil {
		return nil, fmt.Errorf("queries.ListSeasonsBySeriesID: %w", err)
	}
	return &model.SeasonList{
		List: lo.Map(list, func(s sqlcgen.Season, _ int) *model.Season {
			return &model.Season{
				ID:          s.ID,
				CreatedAt:   s.CreatedAt,
				UpdatedAt:   s.UpdatedAt,
				Title:       s.Title,
				Description: s.Description,
				SeriesID:    s.SeriesID,
				Order:       s.Order,
			}
		}),
		Count:  count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (m *Repository) GetSeasonByID(ctx context.Context, id string) (*model.Season, error) {
	season, err := m.queries.GetSeasonByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("queries.GetSeasonByID: %w", customerrors.ErrSeasonNotFound)
		}
		return nil, fmt.Errorf("queries.GetSeasonByID: %w", err)
	}
	return &model.Season{
		ID:          season.ID,
		CreatedAt:   season.CreatedAt,
		UpdatedAt:   season.UpdatedAt,
		Title:       season.Title,
		Description: season.Description,
		SeriesID:    season.SeriesID,
		Order:       season.Order,
	}, nil
}

func (m *Repository) UpdateSeasonByID(ctx context.Context, season *model.Season) error {
	_, err := m.queries.UpdateSeasonByID(ctx, sqlcgen.UpdateSeasonByIDParams{
		ID:          season.ID,
		Title:       season.Title,
		Description: season.Description,
		Order:       season.Order,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.UpdateSeasonByID: %w", customerrors.ErrSeasonNotFound)
		}
		return fmt.Errorf("queries.UpdateSeasonByID: %w", err)
	}
	return nil
}

func (m *Repository) DeleteSeasonByID(ctx context.Context, id string) error {
	_, err := m.queries.DeleteSeasonByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.DeleteSeasonByID: %w", customerrors.ErrSeasonNotFound)
		}
		return fmt.Errorf("queries.DeleteSeasonByID: %w", err)
	}
	return nil
}

func (m *Repository) CreateEpisode(ctx context.Context, episode *model.Episode) error {
	err := m.queries.CreateEpisode(ctx, sqlcgen.CreateEpisodeParams{
		ID:          episode.ID,
		Title:       episode.Title,
		Description: episode.Description,
		SeasonID:    episode.SeasonID,
		Order:       episode.Order,
		MediaID:     episode.MediaID,
	})
	if err != nil {
		return fmt.Errorf("queries.CreateEpisode: %w", err)
	}
	return nil
}

func (m *Repository) ListEpisodesBySeasonID(ctx context.Context, seasonID string, limit, offset int64) (*model.EpisodeList, error) {
	count, err := m.queries.CountEpisodesBySeasonID(ctx, seasonID)
	if err != nil {
		return nil, fmt.Errorf("queries.CountEpisodesBySeasonID: %w", err)
	}
	if count == 0 {
		return &model.EpisodeList{
			List:   make([]*model.Episode, 0),
			Count:  count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	list, err := m.queries.ListEpisodesBySeasonID(ctx, sqlcgen.ListEpisodesBySeasonIDParams{
		SeasonID: seasonID,
		Limit:    limit,
		Offset:   offset,
	})
	if err != nil {
		return nil, fmt.Errorf("queries.ListEpisodesBySeasonID: %w", err)
	}
	return &model.EpisodeList{
		List: lo.Map(list, func(e sqlcgen.Episode, _ int) *model.Episode {
			return &model.Episode{
				ID:          e.ID,
				CreatedAt:   e.CreatedAt,
				UpdatedAt:   e.UpdatedAt,
				Title:       e.Title,
				Description: e.Description,
				SeasonID:    e.SeasonID,
				Order:       e.Order,
				MediaID:     e.MediaID,
			}
		}),
		Count:  count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (m *Repository) GetEpisodeByID(ctx context.Context, id string) (*model.Episode, error) {
	episode, err := m.queries.GetEpisodeByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("queries.GetEpisodeByID: %w", customerrors.ErrEpisodeNotFound)
		}
		return nil, fmt.Errorf("queries.GetEpisodeByID: %w", err)
	}
	return &model.Episode{
		ID:          episode.ID,
		CreatedAt:   episode.CreatedAt,
		UpdatedAt:   episode.UpdatedAt,
		Title:       episode.Title,
		Description: episode.Description,
		SeasonID:    episode.SeasonID,
		Order:       episode.Order,
		MediaID:     episode.MediaID,
	}, nil
}

func (m *Repository) UpdateEpisodeByID(ctx context.Context, episode *model.Episode) error {
	_, err := m.queries.UpdateEpisodeByID(ctx, sqlcgen.UpdateEpisodeByIDParams{
		ID:          episode.ID,
		Title:       episode.Title,
		Description: episode.Description,
		Order:       episode.Order,
		MediaID:     episode.MediaID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.UpdateEpisodeByID: %w", customerrors.ErrEpisodeNotFound)
		}
		return fmt.Errorf("queries.UpdateEpisodeByID: %w", err)
	}
	return nil
}

func (m *Repository) DeleteEpisodeByID(ctx context.Context, id string) error {
	_, err := m.queries.DeleteEpisodeByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("queries.DeleteEpisodeByID: %w", customerrors.ErrEpisodeNotFound)
		}
		return fmt.Errorf("queries.DeleteEpisodeByID: %w", err)
	}
	return nil
}
