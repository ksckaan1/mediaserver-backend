package app

import (
	"common/pb/tmdbpb"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"tmdb_service/internal/domain/core/customerrors"
	"tmdb_service/internal/domain/core/models"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ tmdbpb.TMDBServiceServer = (*TMDB)(nil)

type TMDB struct {
	tmdbClient TMDBClient
	repository Repository
	tmdbpb.UnimplementedTMDBServiceServer
}

func New(tmdbClient TMDBClient, repository Repository) *TMDB {
	return &TMDB{
		tmdbClient: tmdbClient,
		repository: repository,
	}
}

func (t *TMDB) GetTMDBInfo(ctx context.Context, req *tmdbpb.GetTMDBInfoRequest) (*tmdbpb.TMDBInfo, error) {
	info, err := t.repository.GetTMDBInfo(ctx, req.Id)
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

	entityType, id := t.getEntityType(req.Id)

	switch entityType {
	case "movie":
		data, err = t.tmdbClient.GetMovieDetail(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("tmdbClient.GetMovieDetail: %w", err)
		}
	case "series":
		data, err = t.tmdbClient.GetSeriesDetail(ctx, id)
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

	err = t.repository.SetTMDBInfo(ctx, info)
	if err != nil {
		return nil, fmt.Errorf("repository.SetTMDBInfo: %w", err)
	}

	return &tmdbpb.TMDBInfo{
		Id:        req.Id,
		Data:      structData,
		UpdatedAt: timestamppb.New(time.Now()),
	}, nil
}

func (t *TMDB) getEntityType(id string) (string, string) {
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
