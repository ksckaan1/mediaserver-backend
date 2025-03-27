package app

import (
	"context"
	"fmt"
	"season_service/internal/core/models"
	"shared/pb/episodepb"
	"shared/pb/seasonpb"
	"shared/pb/seriespb"
	"shared/ports"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type App struct {
	seasonpb.UnimplementedSeasonServiceServer
	repository    Repository
	seriesClient  seriespb.SeriesServiceClient
	episodeClient episodepb.EpisodeServiceClient
	idGenerator   ports.IDGenerator
}

func New(repository Repository, seriesClient seriespb.SeriesServiceClient, episodeClient episodepb.EpisodeServiceClient, idGenerator ports.IDGenerator) *App {
	return &App{
		repository:    repository,
		seriesClient:  seriesClient,
		episodeClient: episodeClient,
		idGenerator:   idGenerator,
	}
}

func (a *App) CreateSeason(ctx context.Context, req *seasonpb.CreateSeasonRequest) (*seasonpb.CreateSeasonResponse, error) {
	_, err := a.seriesClient.GetSeriesByID(ctx, &seriespb.GetSeriesByIDRequest{SeriesId: req.SeriesId})
	if err != nil {
		return nil, fmt.Errorf("seriesClient.GetSeriesByID: %w", err)
	}

	order, err := a.generateOrderNumber(ctx, req.SeriesId)
	if err != nil {
		return nil, fmt.Errorf("generateOrderNumber: %w", err)
	}

	id := a.idGenerator.NewID()

	err = a.repository.CreateSeason(ctx, &models.Season{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Order:       order,
		SeriesID:    req.SeriesId,
	})
	if err != nil {
		return nil, fmt.Errorf("repository.CreateSeason: %w", err)
	}

	return &seasonpb.CreateSeasonResponse{
		SeasonId: id,
	}, nil
}

func (a *App) DeleteSeasonByID(ctx context.Context, req *seasonpb.DeleteSeasonByIDRequest) (*emptypb.Empty, error) {
	err := a.repository.DeleteSeasonByID(ctx, req.SeasonId)
	if err != nil {
		return nil, fmt.Errorf("repository.DeleteSeasonByID: %w", err)
	}

	_, err = a.episodeClient.DeleteAllEpisodesBySeasonID(ctx, &episodepb.DeleteAllEpisodesBySeasonIDRequest{
		SeasonId: req.SeasonId,
	})
	if err != nil {
		return nil, fmt.Errorf("episodeClient.DeleteAllEpisodesBySeasonID: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func (a *App) DeleteAllSeasonsBySeriesID(ctx context.Context, req *seasonpb.DeleteAllSeasonsBySeriesIDRequest) (*emptypb.Empty, error) {
	seasons, err := a.repository.ListSeasonsBySeriesID(ctx, req.SeriesId)
	if err != nil {
		return nil, fmt.Errorf("repository.ListSeasonsBySeriesID: %w", err)
	}

	if len(seasons) == 0 {
		return &emptypb.Empty{}, nil
	}

	for _, season := range seasons {
		_, err := a.DeleteSeasonByID(ctx, &seasonpb.DeleteSeasonByIDRequest{SeasonId: season.ID})
		if err != nil {
			return nil, fmt.Errorf("DeleteSeasonByID: %w", err)
		}
	}

	return &emptypb.Empty{}, nil
}

func (a *App) GetSeasonByID(ctx context.Context, req *seasonpb.GetSeasonByIDRequest) (*seasonpb.Season, error) {
	season, err := a.repository.GetSeasonByID(ctx, req.SeasonId)
	if err != nil {
		return nil, fmt.Errorf("repository.GetSeasonByID: %w", err)
	}

	return &seasonpb.Season{
		Id:          season.ID,
		CreatedAt:   timestamppb.New(season.CreatedAt),
		UpdatedAt:   timestamppb.New(season.UpdatedAt),
		Title:       season.Title,
		Description: season.Description,
		Order:       season.Order,
		SeriesId:    season.SeriesID,
	}, nil
}

func (a *App) ListSeasonsBySeriesID(ctx context.Context, req *seasonpb.ListSeasonsBySeriesIDRequest) (*seasonpb.SeasonList, error) {
	seasons, err := a.repository.ListSeasonsBySeriesID(ctx, req.SeriesId)
	if err != nil {
		return nil, fmt.Errorf("repository.ListSeasonsBySeriesID: %w", err)
	}

	return &seasonpb.SeasonList{
		List: lo.Map(seasons, func(season *models.Season, _ int) *seasonpb.Season {
			return &seasonpb.Season{
				Id:          season.ID,
				CreatedAt:   timestamppb.New(season.CreatedAt),
				UpdatedAt:   timestamppb.New(season.UpdatedAt),
				Title:       season.Title,
				Description: season.Description,
				Order:       season.Order,
				SeriesId:    season.SeriesID,
			}
		}),
	}, nil
}

func (a *App) ReorderSeasonsBySeriesID(ctx context.Context, req *seasonpb.ReorderSeasonsBySeriesIDRequest) (*emptypb.Empty, error) {
	seasons, err := a.repository.ListSeasonsBySeriesID(ctx, req.SeriesId)
	if err != nil {
		return nil, fmt.Errorf("repository.ListSeasonsBySeriesID: %w", err)
	}

	if len(lo.Intersect(req.SeasonIds, lo.Map(seasons, func(season *models.Season, _ int) string {
		return season.ID
	}))) != len(seasons) {
		return nil, fmt.Errorf("incorrect season ids")
	}

	for i, season := range seasons {
		if season.ID == req.SeasonIds[i] {
			continue
		}
		index := lo.IndexOf(req.SeasonIds, season.ID)
		season.Order = int32(index + 1)
		err = a.repository.UpdateSeasonOrderByID(ctx, season)
		if err != nil {
			return nil, fmt.Errorf("repository.UpdateSeasonOrderByID: %w", err)
		}
	}
	return &emptypb.Empty{}, nil
}

func (a *App) UpdateSeasonByID(ctx context.Context, req *seasonpb.UpdateSeasonByIDRequest) (*emptypb.Empty, error) {
	err := a.repository.UpdateSeasonByID(ctx, &models.Season{
		ID:          req.SeasonId,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("repository.UpdateSeasonByID: %w", err)
	}

	return &emptypb.Empty{}, nil
}
