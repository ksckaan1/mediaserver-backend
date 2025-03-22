package app

import (
	"common/pb/mediapb"
	"common/pb/moviepb"
	"common/pb/tmdbpb"
	"context"
	"fmt"
	"movie_service/internal/domain/core/models"
	"movie_service/internal/port"
	"time"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ moviepb.MovieServiceServer = (*Movie)(nil)

type Movie struct {
	moviepb.UnimplementedMovieServiceServer
	idGenerator   port.IDGenerator
	mediaClient   mediapb.MediaServiceClient
	repo          Repository
	fuzzySearcher *fuzzySearch
	tmdbClient    tmdbpb.TMDBServiceClient
}

func New(
	repo Repository,
	idGenerator port.IDGenerator,
	mediaClient mediapb.MediaServiceClient,
	tmdbClient tmdbpb.TMDBServiceClient,
) (*Movie, error) {
	return &Movie{
		repo:          repo,
		idGenerator:   idGenerator,
		mediaClient:   mediaClient,
		fuzzySearcher: &fuzzySearch{},
		tmdbClient:    tmdbClient,
	}, nil
}

func (m *Movie) CreateMovie(ctx context.Context, request *moviepb.CreateMovieRequest) (*moviepb.CreateMovieResponse, error) {
	err := m.validateMedia(ctx, request.MediaId)
	if err != nil {
		return nil, fmt.Errorf("validateMedia: %w", err)
	}
	err = m.validateTMDBInfo(ctx, request.TmdbId)
	if err != nil {
		return nil, fmt.Errorf("validateTMDBInfo: %w", err)
	}
	id := m.idGenerator.NewID()
	err = m.repo.CreateMovie(ctx, &models.Movie{
		ID:          id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       request.Title,
		Description: request.Description,
		MediaID:     request.MediaId,
		TMDBID:      request.TmdbId,
	})
	if err != nil {
		return nil, fmt.Errorf("repo.CreateMovie: %w", err)
	}
	return &moviepb.CreateMovieResponse{
		MovieId: id,
	}, nil
}

func (m *Movie) GetMovieByID(ctx context.Context, request *moviepb.GetMovieByIDRequest) (*moviepb.Movie, error) {
	movie, err := m.repo.GetMovieByID(ctx, request.MovieId)
	if err != nil {
		return nil, fmt.Errorf("repo.GetMovieByID: %w", err)
	}
	var mediaInfo *moviepb.Media
	if movie.MediaID != "" {
		result, err := m.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{
			MediaId: movie.MediaID,
		})
		if err != nil {
			return nil, fmt.Errorf("mediaClient.GetMediaByID: %w", err)
		}
		mediaInfo = &moviepb.Media{
			Id:        result.Id,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
			Title:     result.Title,
			Path:      result.Path,
			Type:      moviepb.MediaType(result.Type),
			MimeType:  result.MimeType,
			Size:      result.Size,
		}
	}
	var tmdbInfo *moviepb.TMDBInfo
	if movie.TMDBID != "" {
		result, err := m.tmdbClient.GetTMDBInfo(ctx, &tmdbpb.GetTMDBInfoRequest{
			Id: movie.TMDBID,
		})
		if err != nil {
			return nil, fmt.Errorf("tmdbClient.GetTMDBInfo: %w", err)
		}
		tmdbInfo = &moviepb.TMDBInfo{
			Id:   result.Id,
			Data: result.Data,
		}
	}
	return &moviepb.Movie{
		Id:          movie.ID,
		CreatedAt:   timestamppb.New(movie.CreatedAt),
		UpdatedAt:   timestamppb.New(movie.UpdatedAt),
		Title:       movie.Title,
		Description: movie.Description,
		MediaInfo:   mediaInfo,
		TmdbInfo:    tmdbInfo,
	}, nil
}

func (m *Movie) ListMovies(ctx context.Context, request *moviepb.ListMoviesRequest) (*moviepb.MovieList, error) {
	movies, err := m.list(ctx, request.Limit, request.Offset, request.Search)
	if err != nil {
		return nil, fmt.Errorf("list: %w", err)
	}
	return &moviepb.MovieList{
		List: lo.Map(movies.List, func(m *models.Movie, _ int) *moviepb.Movie {
			return &moviepb.Movie{
				Id:          m.ID,
				CreatedAt:   timestamppb.New(m.CreatedAt),
				UpdatedAt:   timestamppb.New(m.UpdatedAt),
				Title:       m.Title,
				Description: m.Description,
			}
		}),
		Count:  movies.Count,
		Limit:  movies.Limit,
		Offset: movies.Offset,
	}, nil
}

func (m *Movie) UpdateMovieByID(ctx context.Context, request *moviepb.UpdateMovieByIDRequest) (*emptypb.Empty, error) {
	err := m.validateMedia(ctx, request.MediaId)
	if err != nil {
		return nil, fmt.Errorf("validateMedia: %w", err)
	}
	err = m.validateTMDBInfo(ctx, request.TmdbId)
	if err != nil {
		return nil, fmt.Errorf("validateTMDBInfo: %w", err)
	}
	err = m.repo.UpdateMovieByID(ctx, &models.Movie{
		ID:          request.MovieId,
		Title:       request.Title,
		Description: request.Description,
		MediaID:     request.MediaId,
		TMDBID:      request.TmdbId,
	})
	if err != nil {
		return nil, fmt.Errorf("repo.UpdateMovieByID: %w", err)
	}
	return nil, nil
}

func (m *Movie) DeleteMovieByID(ctx context.Context, request *moviepb.DeleteMovieByIDRequest) (*emptypb.Empty, error) {
	err := m.repo.DeleteMovieByID(ctx, request.MovieId)
	if err != nil {
		return nil, fmt.Errorf("repo.DeleteMovieByID: %w", err)
	}
	return nil, nil
}
