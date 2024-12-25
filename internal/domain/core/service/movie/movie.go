package movie

import (
	"context"
	"fmt"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/port"
)

type Repository interface {
	// Movie
	CreateMovie(ctx context.Context, movie *model.Movie) error
	GetMovieByID(ctx context.Context, id string) (*model.Movie, error)
	ListMovies(ctx context.Context, limit, offset int64) (*model.MovieList, error)
	UpdateMovieByID(ctx context.Context, movie *model.Movie) error
	DeleteMovieByID(ctx context.Context, id string) error

	// TMDBInfo
	SetTMDBInfo(ctx context.Context, info *model.TMDBInfo) error
	GetTMDBInfoByID(ctx context.Context, id string) (*model.TMDBInfo, error)

	// Media
	GetMediaByID(ctx context.Context, id string) (*model.Media, error)
}

type TMDBClient interface {
	GetMovieDetail(ctx context.Context, id string) (*model.TMDBInfo, error)
}

type Movie struct {
	repo       Repository
	tmdbClient TMDBClient
	idgen      port.IDGenerator
	logger     port.Logger
}

func New(repo Repository, tmdbClient TMDBClient, idgen port.IDGenerator, lg port.Logger) (*Movie, error) {
	return &Movie{
		repo:       repo,
		tmdbClient: tmdbClient,
		idgen:      idgen,
		logger:     lg,
	}, nil
}

func (m *Movie) CreateMovie(ctx context.Context, movie *model.Movie) (string, error) {
	movie.ID = m.idgen.NewID()

	err := m.validateMedia(ctx, movie.MediaID)
	if err != nil {
		err = fmt.Errorf("validate media: %w", err)
		m.logger.Error(ctx, "error when creating movie",
			"error", err,
		)
		return "", err
	}

	var tmdbInfo *model.TMDBInfo

	if movie.TMDBID != "" {
		tmdbInfo, err = m.tmdbClient.GetMovieDetail(ctx, movie.TMDBID)
		if err != nil {
			err = fmt.Errorf("tmdbClient.GetMovieDetail: %w", err)
			m.logger.Error(ctx, "error when getting movie detail from tmdb",
				"error", err,
			)
			return "", err
		}
	}

	if movie.MediaID != "" {
		_, err := m.repo.GetMediaByID(ctx, movie.MediaID)
		if err != nil {
			err = fmt.Errorf("repo.GetMediaByID: %w", err)
			m.logger.Error(ctx, "error when getting media",
				"error", err,
			)
			return "", err
		}
	}

	err = m.repo.CreateMovie(ctx, movie)
	if err != nil {
		err = fmt.Errorf("repo.CreateMovie: %w", err)
		m.logger.Error(ctx, "error when creating movie",
			"error", err,
		)
		return "", err
	}

	m.logger.Info(ctx, "movie created",
		"id", movie.ID, "title", movie.Title,
	)

	if tmdbInfo != nil {
		err = m.repo.SetTMDBInfo(ctx, tmdbInfo)
		if err != nil {
			err = fmt.Errorf("repo.SetTMDBInfo: %w", err)
			m.logger.Error(ctx, "error when setting tmdb info",
				"error", err,
			)
			return "", fmt.Errorf("repo.SetTMDBInfo: %w", err)
		}

		m.logger.Info(ctx, "tmdb info set",
			"id", tmdbInfo.ID,
		)
	}

	return movie.ID, nil
}

func (m *Movie) GetMovieByID(ctx context.Context, id string) (*model.GetMovieByIDResponse, error) {
	movie, err := m.repo.GetMovieByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("repo.GetMovieByID: %w", err)
		m.logger.Error(ctx, "error when getting movie",
			"error", err,
			"id", id,
		)
		return nil, err
	}

	m.logger.Info(ctx, "movie retrieved",
		"id", movie.ID,
		"title", movie.Title,
	)

	var ti *model.TMDBInfo

	if movie.TMDBID != "" {
		tmdbInfo, err := m.repo.GetTMDBInfoByID(ctx, fmt.Sprintf("movie-%s", movie.TMDBID))
		if err != nil {
			err = fmt.Errorf("repo.GetTMDBInfoByID: %w", err)
			m.logger.Error(ctx, "error when getting tmdb info",
				"error", err,
				"movie_id", movie.ID,
				"tmdb_id", movie.TMDBID,
			)
		}
		ti = tmdbInfo
	}

	var mi *model.Media

	if movie.MediaID != "" {
		media, err := m.repo.GetMediaByID(ctx, movie.MediaID)
		if err != nil {
			err = fmt.Errorf("repo.GetMediaByID: %w", err)
			m.logger.Error(ctx, "error when getting media",
				"error", err,
				"movie_id", movie.ID,
				"media_id", movie.MediaID,
			)
		}
		mi = media
	}

	return &model.GetMovieByIDResponse{
		ID:          id,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
		Title:       movie.Title,
		Description: movie.Description,
		TMDBInfo:    ti,
		MediaInfo:   mi,
	}, nil
}

func (m *Movie) ListMovies(ctx context.Context, limit, offset int64) (*model.MovieList, error) {
	movies, err := m.repo.ListMovies(ctx, limit, offset)
	if err != nil {
		err = fmt.Errorf("repo.ListMovies: %w", err)
		m.logger.Error(ctx, "error when listing movies",
			"error", err,
		)
		return nil, err
	}
	m.logger.Info(ctx, "movies listed",
		"count", movies.Count,
		"limit", movies.Limit,
		"offset", movies.Offset,
	)
	return movies, nil
}

func (m *Movie) UpdateMovieByID(ctx context.Context, movie *model.Movie) error {
	_, err := m.repo.GetMovieByID(ctx, movie.ID)
	if err != nil {
		err = fmt.Errorf("repo.GetMovieByID: %w", err)
		m.logger.Error(ctx, "error when getting movie",
			"error", err,
			"id", movie.ID,
		)
		return err
	}

	var tmdbInfo *model.TMDBInfo

	if movie.TMDBID != "" {
		tmdbInfo, err = m.tmdbClient.GetMovieDetail(ctx, movie.TMDBID)
		if err != nil {
			err = fmt.Errorf("tmdbClient.GetMovieDetail: %w", err)
			m.logger.Error(ctx, "error when getting tmdb info",
				"error", err,
				"movie_id", movie.ID,
				"tmdb_id", movie.TMDBID,
			)
			return err
		}
	}

	if movie.MediaID != "" {
		_, err := m.repo.GetMediaByID(ctx, movie.MediaID)
		if err != nil {
			err = fmt.Errorf("repo.GetMediaByID: %w", err)
			m.logger.Error(ctx, "error when getting media",
				"error", err,
			)
			return err
		}
	}

	err = m.repo.UpdateMovieByID(ctx, movie)
	if err != nil {
		err = fmt.Errorf("repo.UpdateMovieByID: %w", err)
		m.logger.Error(ctx, "error when updating movie",
			"error", err,
			"id", movie.ID,
		)
		return err
	}
	m.logger.Info(ctx, "movie updated",
		"id", movie.ID,
		"title", movie.Title,
	)

	if tmdbInfo != nil {
		err = m.repo.SetTMDBInfo(ctx, tmdbInfo)
		if err != nil {
			err = fmt.Errorf("repo.SetTMDBInfo: %w", err)
			m.logger.Error(ctx, "error when setting tmdb info",
				"error", err,
			)
			return fmt.Errorf("repo.SetTMDBInfo: %w", err)
		}

		m.logger.Info(ctx, "tmdb info set",
			"id", tmdbInfo.ID,
		)
	}

	return nil
}

func (m *Movie) DeleteMovieByID(ctx context.Context, id string) error {
	err := m.repo.DeleteMovieByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("repo.DeleteMovieByID: %w", err)
		m.logger.Error(ctx, "error when deleting movie",
			"error", err,
			"id", id,
		)
		return err
	}
	m.logger.Info(ctx, "movie deleted",
		"id", id,
	)
	return nil
}
