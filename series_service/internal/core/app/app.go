package app

import (
	"common/pb/seriespb"
	"common/pb/tmdbpb"
	"common/ports"
	"context"
	"fmt"
	"series_service/internal/core/models"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type App struct {
	seriespb.UnimplementedSeriesServiceServer
	tmdbClient  tmdbpb.TMDBServiceClient
	repo        Repository
	idGenerator ports.IDGenerator
}

func New(repo Repository, tmdbClient tmdbpb.TMDBServiceClient, idGenerator ports.IDGenerator) *App {
	return &App{
		tmdbClient:  tmdbClient,
		repo:        repo,
		idGenerator: idGenerator,
	}
}

func (h *App) CreateSeries(ctx context.Context, req *seriespb.CreateSeriesRequest) (*seriespb.CreateSeriesResponse, error) {
	err := h.validateTMDBInfo(ctx, req.TmdbId)
	if err != nil {
		return nil, fmt.Errorf("validateTMDBInfo: %w", err)
	}
	id := h.idGenerator.NewID()
	err = h.repo.CreateSeries(ctx, &models.Series{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		TMDBID:      req.TmdbId,
	})
	if err != nil {
		return nil, fmt.Errorf("repo.CreateSeries: %w", err)
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
			}
		}),
		Count:  series.Count,
		Limit:  series.Limit,
		Offset: series.Offset,
	}, nil
}

func (h *App) UpdateSeriesByID(ctx context.Context, req *seriespb.UpdateSeriesByIDRequest) (*emptypb.Empty, error) {
	err := h.validateTMDBInfo(ctx, req.TmdbId)
	if err != nil {
		return nil, fmt.Errorf("validateTMDBInfo: %w", err)
	}
	err = h.repo.UpdateSeriesByID(ctx, &models.Series{
		ID:          req.SeriesId,
		Title:       req.Title,
		Description: req.Description,
		TMDBID:      req.TmdbId,
	})
	if err != nil {
		return nil, fmt.Errorf("repo.UpdateSeriesByID: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (h *App) DeleteSeriesByID(ctx context.Context, req *seriespb.DeleteSeriesByIDRequest) (*emptypb.Empty, error) {
	err := h.repo.DeleteSeriesByID(ctx, req.SeriesId)
	if err != nil {
		return nil, fmt.Errorf("repo.DeleteSeriesByID: %w", err)
	}
	return &emptypb.Empty{}, nil
}
