package app

import (
	"context"
	"fmt"
	"series_service/internal/core/models"
	"shared/pb/seriespb"
	"shared/pb/tmdbpb"
	"shared/ports"
	"time"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ seriespb.SeriesServiceServer = (*App)(nil)

type App struct {
	seriespb.UnimplementedSeriesServiceServer
	tmdbClient  tmdbpb.TMDBServiceClient
	repo        Repository
	idGenerator ports.IDGenerator
	searcher    Searcher
}

func New(repo Repository, tmdbClient tmdbpb.TMDBServiceClient, idGenerator ports.IDGenerator, searcher Searcher) *App {
	return &App{
		tmdbClient:  tmdbClient,
		repo:        repo,
		idGenerator: idGenerator,
		searcher:    searcher,
	}
}

func (h *App) CreateSeries(ctx context.Context, req *seriespb.CreateSeriesRequest) (*seriespb.CreateSeriesResponse, error) {
	tmdbInfo, err := h.validateTMDBInfo(ctx, req.TmdbId)
	if err != nil {
		return nil, fmt.Errorf("validateTMDBInfo: %w", err)
	}
	id := h.idGenerator.NewID()
	now := time.Now()
	series := &models.Series{
		ID:          id,
		CreatedAt:   now,
		UpdatedAt:   now,
		Title:       req.Title,
		Description: req.Description,
		TMDBID:      req.TmdbId,
		Tags:        req.Tags,
	}
	err = h.repo.CreateSeries(ctx, series)
	if err != nil {
		return nil, fmt.Errorf("repo.CreateSeries: %w", err)
	}

	var searchTMDBInfo *seriespb.TMDBInfo
	if tmdbInfo != nil {
		searchTMDBInfo = &seriespb.TMDBInfo{
			Id:   tmdbInfo.Id,
			Data: tmdbInfo.Data,
		}
	}
	err = h.searcher.AddDocument(ctx, "series", &seriespb.Series{
		Id:          id,
		CreatedAt:   timestamppb.New(series.CreatedAt),
		UpdatedAt:   timestamppb.New(series.UpdatedAt),
		Title:       series.Title,
		Description: series.Description,
		TmdbInfo:    searchTMDBInfo,
		Tags:        series.Tags,
	})
	if err != nil {
		return nil, fmt.Errorf("searcher.AddDocument: %w", err)
	}
	return &seriespb.CreateSeriesResponse{
		SeriesId: id,
	}, nil
}

func (h *App) GetSeriesByID(ctx context.Context, req *seriespb.GetSeriesByIDRequest) (*seriespb.Series, error) {
	series, err := h.repo.GetSeriesByID(ctx, req.SeriesId)
	if err != nil {
		return nil, fmt.Errorf("repo.GetSeriesByID: %w", err)
	}
	var tmdbInfo *seriespb.TMDBInfo
	if series.TMDBID != "" {
		info, err := h.tmdbClient.GetTMDBInfo(ctx, &tmdbpb.GetTMDBInfoRequest{
			Id: series.TMDBID,
		})
		if err != nil {
			return nil, fmt.Errorf("tmdbClient.GetTMDBInfo: %w", err)
		}
		tmdbInfo = &seriespb.TMDBInfo{
			Id:   series.TMDBID,
			Data: info.Data,
		}
	}
	return &seriespb.Series{
		Id:          series.ID,
		CreatedAt:   timestamppb.New(series.CreatedAt),
		UpdatedAt:   timestamppb.New(series.UpdatedAt),
		Title:       series.Title,
		Description: series.Description,
		TmdbInfo:    tmdbInfo,
		Tags:        series.Tags,
	}, nil
}

func (h *App) ListSeries(ctx context.Context, req *seriespb.ListSeriesRequest) (*seriespb.SeriesList, error) {
	series, err := h.repo.ListSeries(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("repo.ListSeries: %w", err)
	}
	return &seriespb.SeriesList{
		List: lo.Map(series.List, func(s *models.Series, _ int) *seriespb.Series {
			return &seriespb.Series{
				Id:          s.ID,
				CreatedAt:   timestamppb.New(s.CreatedAt),
				UpdatedAt:   timestamppb.New(s.UpdatedAt),
				Title:       s.Title,
				Description: s.Description,
				Tags:        s.Tags,
			}
		}),
		Count:  series.Count,
		Limit:  series.Limit,
		Offset: series.Offset,
	}, nil
}

func (h *App) SearchSeries(ctx context.Context, req *seriespb.SearchSeriesRequest) (*seriespb.SeriesList, error) {
	var series []*seriespb.Series
	count, err := h.searcher.Search(ctx,
		"series",
		req.Query,
		req.QueryBy,
		int(req.Limit),
		int(req.Offset),
		false,
		&series,
	)
	if err != nil {
		return nil, fmt.Errorf("searcher.Search: %w", err)
	}
	return &seriespb.SeriesList{
		List:   series,
		Count:  int64(count),
		Limit:  int64(req.Limit),
		Offset: int64(req.Offset),
	}, nil
}

func (h *App) UpdateSeriesByID(ctx context.Context, req *seriespb.UpdateSeriesByIDRequest) (*emptypb.Empty, error) {
	series, err := h.repo.GetSeriesByID(ctx, req.SeriesId)
	if err != nil {
		return nil, fmt.Errorf("repo.GetSeriesByID: %w", err)
	}
	tmdbInfo, err := h.validateTMDBInfo(ctx, req.TmdbId)
	if err != nil {
		return nil, fmt.Errorf("validateTMDBInfo: %w", err)
	}

	series.UpdatedAt = time.Now()
	series.Title = req.Title
	series.Description = req.Description
	series.TMDBID = req.TmdbId
	series.Tags = req.Tags

	err = h.repo.UpdateSeriesByID(ctx, series)
	if err != nil {
		return nil, fmt.Errorf("repo.UpdateSeriesByID: %w", err)
	}
	var searchTMDBInfo *seriespb.TMDBInfo
	if tmdbInfo != nil {
		searchTMDBInfo = &seriespb.TMDBInfo{
			Id:   tmdbInfo.Id,
			Data: tmdbInfo.Data,
		}
	}
	err = h.searcher.UpdateDocument(ctx, "series", req.SeriesId, &seriespb.Series{
		Id:          series.ID,
		CreatedAt:   timestamppb.New(series.CreatedAt),
		UpdatedAt:   timestamppb.New(series.UpdatedAt),
		Title:       series.Title,
		Description: series.Description,
		TmdbInfo:    searchTMDBInfo,
		Tags:        series.Tags,
	})
	if err != nil {
		return nil, fmt.Errorf("searcher.Update: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (h *App) DeleteSeriesByID(ctx context.Context, req *seriespb.DeleteSeriesByIDRequest) (*emptypb.Empty, error) {
	err := h.repo.DeleteSeriesByID(ctx, req.SeriesId)
	if err != nil {
		return nil, fmt.Errorf("repo.DeleteSeriesByID: %w", err)
	}
	err = h.searcher.DeleteDocument(ctx, "series", req.SeriesId)
	if err != nil {
		return nil, fmt.Errorf("searcher.Delete: %w", err)
	}
	return &emptypb.Empty{}, nil
}
