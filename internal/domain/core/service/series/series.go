package series

import (
	"context"
	"fmt"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/port"
)

type Repository interface {
	// Series
	CreateSeries(ctx context.Context, series *model.Series) error
	GetSeriesByID(ctx context.Context, id string) (*model.Series, error)
	ListSeries(ctx context.Context, limit, offset int64) (*model.SeriesList, error)
	UpdateSeriesByID(ctx context.Context, series *model.Series) error
	DeleteSeriesByID(ctx context.Context, id string) error

	// TMDBInfo
	SetTMDBInfo(ctx context.Context, info *model.TMDBInfo) error
	GetTMDBInfoByID(ctx context.Context, id string) (*model.TMDBInfo, error)
}

type TMDBClient interface {
	GetMovieDetail(ctx context.Context, id string) (*model.TMDBInfo, error)
	GetSeriesDetail(ctx context.Context, id string) (*model.TMDBInfo, error)
}

type Series struct {
	repo       Repository
	tmdbClient TMDBClient
	idgen      port.IDGenerator
	logger     port.Logger
}

func New(repo Repository, tmdbClient TMDBClient, idgen port.IDGenerator, lg port.Logger) (*Series, error) {
	return &Series{
		repo:       repo,
		tmdbClient: tmdbClient,
		idgen:      idgen,
		logger:     lg,
	}, nil
}

func (s *Series) CreateSeries(ctx context.Context, series *model.Series) (string, error) {
	var (
		tmdbInfo *model.TMDBInfo
		err      error
	)

	if series.TMDBID != "" {
		tmdbInfo, err = s.tmdbClient.GetSeriesDetail(ctx, series.TMDBID)
		if err != nil {
			err = fmt.Errorf("tmdbClient.GetSeriesDetail: %w", err)
			s.logger.Error(ctx, "error when getting series detail from tmdb",
				"error", err,
			)
			return "", err
		}
	}

	series.ID = s.idgen.NewID()
	err = s.repo.CreateSeries(ctx, series)
	if err != nil {
		err = fmt.Errorf("repo.CreateSeries: %w", err)
		s.logger.Error(ctx, "error when creating series",
			"error", err,
		)
		return "", err
	}
	s.logger.Info(ctx, "series created",
		"series_id", series.ID,
		"title", series.Title,
	)

	if tmdbInfo != nil {
		err = s.repo.SetTMDBInfo(ctx, tmdbInfo)
		if err != nil {
			err = fmt.Errorf("repo.SetTMDBInfo: %w", err)
			s.logger.Error(ctx, "error when setting tmdb info",
				"error", err,
			)
			return "", fmt.Errorf("repo.SetTMDBInfo: %w", err)
		}

		s.logger.Info(ctx, "tmdb info set",
			"id", tmdbInfo.ID,
		)
	}

	return series.ID, nil
}

func (s *Series) GetSeriesByID(ctx context.Context, id string) (*model.GetSeriesResponse, error) {
	series, err := s.repo.GetSeriesByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("repo.GetSeriesByID: %w", err)
		s.logger.Error(ctx, "error when getting series",
			"error", err,
		)
		return nil, err
	}
	s.logger.Info(ctx, "series retrieved",
		"series_id", series.ID,
		"title", series.Title,
	)

	var ti *model.TMDBInfo

	if series.TMDBID != "" {
		tmdbInfo, err := s.repo.GetTMDBInfoByID(ctx, fmt.Sprintf("series-%s", series.TMDBID))
		if err != nil {
			err = fmt.Errorf("repo.GetTMDBInfoByID: %w", err)
			s.logger.Error(ctx, "error when getting tmdb info",
				"error", err,
				"series_id", series.ID,
				"tmdb_id", series.TMDBID,
			)
		}
		ti = tmdbInfo
	}

	return &model.GetSeriesResponse{
		ID:          id,
		CreatedAt:   series.CreatedAt,
		UpdatedAt:   series.UpdatedAt,
		Title:       series.Title,
		Description: series.Description,
		TMDBInfo:    ti,
	}, nil
}

func (s *Series) ListSeries(ctx context.Context, limit, offset int64) (*model.SeriesList, error) {
	list, err := s.repo.ListSeries(ctx, limit, offset)
	if err != nil {
		err = fmt.Errorf("repo.ListSeries: %w", err)
		s.logger.Error(ctx, "error when listing series",
			"error", err,
		)
		return nil, err
	}
	s.logger.Info(ctx, "series listed",
		"count", list.Count,
		"limit", list.Limit,
		"offset", list.Offset,
	)
	return list, nil
}

func (s *Series) UpdateSeriesByID(ctx context.Context, series *model.Series) error {
	err := s.repo.UpdateSeriesByID(ctx, series)
	if err != nil {
		err = fmt.Errorf("repo.UpdateSeriesByID: %w", err)
		s.logger.Error(ctx, "error when updating series",
			"error", err,
		)
		return err
	}
	s.logger.Info(ctx, "series updated",
		"series_id", series.ID,
		"title", series.Title,
	)
	return nil
}

func (s *Series) DeleteSeriesByID(ctx context.Context, id string) error {
	err := s.repo.DeleteSeriesByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("repo.DeleteSeriesByID: %w", err)
		s.logger.Error(ctx, "error when deleting series",
			"error", err,
		)
		return err
	}
	s.logger.Info(ctx, "series deleted",
		"series_id", id,
	)
	return nil
}
