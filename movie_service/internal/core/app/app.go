package app

import (
	"common/pb/mediapb"
	"common/pb/moviepb"
	"common/pb/tmdbpb"
	"common/ports"
	"context"
	"fmt"
	"movie_service/internal/core/models"

	"time"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type App struct {
	moviepb.UnimplementedMovieServiceServer
	idGenerator ports.IDGenerator
	mediaClient mediapb.MediaServiceClient
	repo        Repository
	tmdbClient  tmdbpb.TMDBServiceClient
}

func New(
	repo Repository,
	idGenerator ports.IDGenerator,
	mediaClient mediapb.MediaServiceClient,
	tmdbClient tmdbpb.TMDBServiceClient,
) *App {
	return &App{
		repo:        repo,
		idGenerator: idGenerator,
		mediaClient: mediaClient,
		tmdbClient:  tmdbClient,
	}
}

func (a *App) CreateMovie(ctx context.Context, request *moviepb.CreateMovieRequest) (*moviepb.CreateMovieResponse, error) {
	err := a.validateMedia(ctx, request.MediaId)
	if err != nil {
		return nil, fmt.Errorf("validateMedia: %w", err)
	}
	err = a.validateTMDBInfo(ctx, request.TmdbId)
	if err != nil {
		return nil, fmt.Errorf("validateTMDBInfo: %w", err)
	}
	id := a.idGenerator.NewID()
	err = a.repo.CreateMovie(ctx, &models.Movie{
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

func (a *App) GetMovieByID(ctx context.Context, request *moviepb.GetMovieByIDRequest) (*moviepb.Movie, error) {
	movie, err := a.repo.GetMovieByID(ctx, request.MovieId)
	if err != nil {
		return nil, fmt.Errorf("repo.GetMovieByID: %w", err)
	}
	var mediaInfo *moviepb.Media
	if movie.MediaID != "" {
		result, err := a.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{
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
		result, err := a.tmdbClient.GetTMDBInfo(ctx, &tmdbpb.GetTMDBInfoRequest{
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

func (a *App) ListMovies(ctx context.Context, request *moviepb.ListMoviesRequest) (*moviepb.MovieList, error) {
	movies, err := a.repo.ListMovies(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, fmt.Errorf("repo.ListMovies: %w", err)
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

func (a *App) UpdateMovieByID(ctx context.Context, request *moviepb.UpdateMovieByIDRequest) (*emptypb.Empty, error) {
	err := a.validateMedia(ctx, request.MediaId)
	if err != nil {
		return nil, fmt.Errorf("validateMedia: %w", err)
	}
	err = a.validateTMDBInfo(ctx, request.TmdbId)
	if err != nil {
		return nil, fmt.Errorf("validateTMDBInfo: %w", err)
	}
	err = a.repo.UpdateMovieByID(ctx, &models.Movie{
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

func (a *App) DeleteMovieByID(ctx context.Context, request *moviepb.DeleteMovieByIDRequest) (*emptypb.Empty, error) {
	err := a.repo.DeleteMovieByID(ctx, request.MovieId)
	if err != nil {
		return nil, fmt.Errorf("repo.DeleteMovieByID: %w", err)
	}
	return nil, nil
}
