package app

import (
	"context"
	"fmt"
	"movie_service/internal/core/models"
	"shared/pb/mediapb"
	"shared/pb/moviepb"
	"shared/pb/tmdbpb"
	"shared/ports"

	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ moviepb.MovieServiceServer = (*App)(nil)

type App struct {
	moviepb.UnimplementedMovieServiceServer
	idGenerator ports.IDGenerator
	mediaClient mediapb.MediaServiceClient
	repo        Repository
	tmdbClient  tmdbpb.TMDBServiceClient
	searcher    Searcher
}

func New(
	repo Repository,
	idGenerator ports.IDGenerator,
	mediaClient mediapb.MediaServiceClient,
	tmdbClient tmdbpb.TMDBServiceClient,
	searcher Searcher,
) *App {
	return &App{
		repo:        repo,
		idGenerator: idGenerator,
		mediaClient: mediaClient,
		tmdbClient:  tmdbClient,
		searcher:    searcher,
	}
}

func (a *App) CreateMovie(ctx context.Context, request *moviepb.CreateMovieRequest) (*moviepb.CreateMovieResponse, error) {
	mediaInfo, err := a.validateMedia(ctx, request.MediaId)
	if err != nil {
		return nil, fmt.Errorf("validateMedia: %w", err)
	}
	tmdbInfo, err := a.validateTMDBInfo(ctx, request.TmdbId)
	if err != nil {
		return nil, fmt.Errorf("validateTMDBInfo: %w", err)
	}
	id := a.idGenerator.NewID()
	now := time.Now()
	movie := &models.Movie{
		ID:          id,
		CreatedAt:   now,
		UpdatedAt:   now,
		Title:       request.Title,
		Description: request.Description,
		MediaID:     request.MediaId,
		TMDBID:      request.TmdbId,
		Tags:        request.Tags,
	}
	err = a.repo.CreateMovie(ctx, movie)
	if err != nil {
		return nil, fmt.Errorf("repo.CreateMovie: %w", err)
	}
	var searchMedia *moviepb.Media
	if mediaInfo != nil {
		searchMedia = &moviepb.Media{
			Id:        mediaInfo.Id,
			CreatedAt: mediaInfo.CreatedAt,
			UpdatedAt: mediaInfo.UpdatedAt,
			Title:     mediaInfo.Title,
			Path:      mediaInfo.Path,
			Type:      mediaInfo.Type,
			MimeType:  mediaInfo.MimeType,
			Size:      mediaInfo.Size,
		}
	}
	var searchTMDBInfo *moviepb.TMDBInfo
	if tmdbInfo != nil {
		searchTMDBInfo = &moviepb.TMDBInfo{
			Id:   tmdbInfo.Id,
			Data: tmdbInfo.Data,
		}
	}
	err = a.searcher.AddDocument(ctx, "movies", &moviepb.Movie{
		Id:          id,
		CreatedAt:   timestamppb.New(movie.CreatedAt),
		UpdatedAt:   timestamppb.New(movie.UpdatedAt),
		Title:       request.Title,
		Description: request.Description,
		MediaInfo:   searchMedia,
		TmdbInfo:    searchTMDBInfo,
		Tags:        request.Tags,
	})
	if err != nil {
		return nil, fmt.Errorf("searcher.AddDocument: %w", err)
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
			Type:      result.Type,
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
		Tags:        movie.Tags,
	}, nil
}

func (a *App) ListMovies(ctx context.Context, request *moviepb.ListMoviesRequest) (*moviepb.MovieList, error) {
	movies, err := a.repo.ListMovies(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, fmt.Errorf("repo.ListMovies: %w", err)
	}

	resultList := make([]*moviepb.Movie, 0, len(movies.List))
	for _, m := range movies.List {
		var mediaInfo *moviepb.Media
		if m.MediaID != "" {
			result, err := a.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{
				MediaId: m.MediaID,
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
				Type:      result.Type,
				MimeType:  result.MimeType,
				Size:      result.Size,
			}
		}
		var tmdbInfo *moviepb.TMDBInfo
		if m.TMDBID != "" {
			result, err := a.tmdbClient.GetTMDBInfo(ctx, &tmdbpb.GetTMDBInfoRequest{
				Id: m.TMDBID,
			})
			if err != nil {
				return nil, fmt.Errorf("tmdbClient.GetTMDBInfo: %w", err)
			}
			tmdbInfo = &moviepb.TMDBInfo{
				Id:   result.Id,
				Data: result.Data,
			}
		}
		resultList = append(resultList, &moviepb.Movie{
			Id:          m.ID,
			CreatedAt:   timestamppb.New(m.CreatedAt),
			UpdatedAt:   timestamppb.New(m.UpdatedAt),
			Title:       m.Title,
			Description: m.Description,
			MediaInfo:   mediaInfo,
			TmdbInfo:    tmdbInfo,
			Tags:        m.Tags,
		})
	}

	return &moviepb.MovieList{
		List:   resultList,
		Count:  movies.Count,
		Limit:  movies.Limit,
		Offset: movies.Offset,
	}, nil
}

func (a *App) SearchMovie(ctx context.Context, request *moviepb.SearchMovieRequest) (*moviepb.MovieList, error) {
	var movies []*moviepb.Movie
	count, err := a.searcher.Search(ctx,
		"movies",
		request.Query,
		request.QueryBy,
		int(request.Limit),
		int(request.Offset),
		false,
		&movies,
	)
	if err != nil {
		return nil, fmt.Errorf("searcher.Search: %w", err)
	}
	return &moviepb.MovieList{
		List:   movies,
		Count:  int64(count),
		Limit:  int64(request.Limit),
		Offset: int64(request.Offset),
	}, nil
}

func (a *App) UpdateMovieByID(ctx context.Context, request *moviepb.UpdateMovieByIDRequest) (*emptypb.Empty, error) {
	movie, err := a.repo.GetMovieByID(ctx, request.MovieId)
	if err != nil {
		return nil, fmt.Errorf("repo.GetMovieByID: %w", err)
	}
	mediaInfo, err := a.validateMedia(ctx, request.MediaId)
	if err != nil {
		return nil, fmt.Errorf("validateMedia: %w", err)
	}
	tmdbInfo, err := a.validateTMDBInfo(ctx, request.TmdbId)
	if err != nil {
		return nil, fmt.Errorf("validateTMDBInfo: %w", err)
	}

	movie.MediaID = request.MediaId
	movie.TMDBID = request.TmdbId
	movie.Title = request.Title
	movie.Description = request.Description
	movie.Tags = request.Tags
	movie.UpdatedAt = time.Now()

	err = a.repo.UpdateMovieByID(ctx, movie)
	if err != nil {
		return nil, fmt.Errorf("repo.UpdateMovieByID: %w", err)
	}
	var searchMedia *moviepb.Media
	if mediaInfo != nil {
		searchMedia = &moviepb.Media{
			Id:        mediaInfo.Id,
			CreatedAt: mediaInfo.CreatedAt,
			UpdatedAt: mediaInfo.UpdatedAt,
			Title:     mediaInfo.Title,
			Path:      mediaInfo.Path,
			Type:      mediaInfo.Type,
			MimeType:  mediaInfo.MimeType,
			Size:      mediaInfo.Size,
		}
	}
	var searchTMDBInfo *moviepb.TMDBInfo
	if tmdbInfo != nil {
		searchTMDBInfo = &moviepb.TMDBInfo{
			Id:   tmdbInfo.Id,
			Data: tmdbInfo.Data,
		}
	}
	err = a.searcher.UpdateDocument(ctx, "movies", request.MovieId, &moviepb.Movie{
		Id:          request.MovieId,
		CreatedAt:   timestamppb.New(movie.CreatedAt),
		UpdatedAt:   timestamppb.New(movie.UpdatedAt),
		Title:       movie.Title,
		Description: movie.Description,
		MediaInfo:   searchMedia,
		TmdbInfo:    searchTMDBInfo,
		Tags:        movie.Tags,
	})
	if err != nil {
		return nil, fmt.Errorf("searcher.Update: %w", err)
	}
	return nil, nil
}

func (a *App) DeleteMovieByID(ctx context.Context, request *moviepb.DeleteMovieByIDRequest) (*emptypb.Empty, error) {
	err := a.repo.DeleteMovieByID(ctx, request.MovieId)
	if err != nil {
		return nil, fmt.Errorf("repo.DeleteMovieByID: %w", err)
	}
	err = a.searcher.DeleteDocument(ctx, "movies", request.MovieId)
	if err != nil {
		return nil, fmt.Errorf("searcher.DeleteDocument: %w", err)
	}
	return nil, nil
}
