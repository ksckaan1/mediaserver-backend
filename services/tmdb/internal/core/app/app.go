package app

import (
	"context"
	"errors"
	"fmt"
	"shared/pb/tmdbpb"
	"strings"
	"time"
	"tmdb_service/internal/core/customerrors"
	"tmdb_service/internal/core/models"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type App struct {
	tmdbpb.UnimplementedTMDBServiceServer
	tmdbClient TMDBClient
	repository Repository
}

func New(tmdbClient TMDBClient, repository Repository) *App {
	return &App{
		tmdbClient: tmdbClient,
		repository: repository,
	}
}

func (a *App) GetTMDBInfo(ctx context.Context, req *tmdbpb.GetTMDBInfoRequest) (*tmdbpb.TMDBInfo, error) {
	info, err := a.repository.GetTMDBInfo(ctx, req.Id)
	if err != nil && !errors.Is(err, customerrors.ErrRecordNotFound) {
		return nil, fmt.Errorf("repository.GetTMDBInfo: %w", err)
	}

	if info != nil {
		structData, err := structpb.NewStruct(info.Data)
		if err != nil {
			return nil, fmt.Errorf("structpb.NewStruct: %w", err)
		}
		return &tmdbpb.TMDBInfo{
			Id:        info.Id,
			Data:      structData,
			UpdatedAt: timestamppb.New(info.UpdatedAt),
		}, nil
	}

	var data map[string]any

	entityType, id := a.getEntityType(req.Id)

	switch entityType {
	case "movie":
		data, err = a.tmdbClient.GetMovieDetail(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("tmdbClient.GetMovieDetail: %w", err)
		}
	case "series":
		data, err = a.tmdbClient.GetSeriesDetail(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("tmdbClient.GetSeriesDetail: %w", err)
		}
	default:
		return nil, fmt.Errorf("invalid entity type: %s", entityType)
	}

	structData, err := structpb.NewStruct(data)
	if err != nil {
		return nil, fmt.Errorf("structpb.NewStruct: %w", err)
	}

	info = &models.TMDBInfo{
		Id:   req.Id,
		Data: data,
	}

	err = a.repository.SetTMDBInfo(ctx, info)
	if err != nil {
		return nil, fmt.Errorf("repository.SetTMDBInfo: %w", err)
	}

	return &tmdbpb.TMDBInfo{
		Id:        req.Id,
		Data:      structData,
		UpdatedAt: timestamppb.New(time.Now()),
	}, nil
}

func (a *App) getEntityType(id string) (string, string) {
	if id == "" {
		return "", ""
	}
	if strings.HasPrefix(id, "series:") {
		return "series", id[7:]
	}
	if strings.HasPrefix(id, "movie:") {
		return "movie", id[6:]
	}
	return "", ""
}
