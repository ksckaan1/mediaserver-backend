package app

import (
	"context"
	"episode_service/internal/core/models"
	"fmt"
	"shared/pb/episodepb"
	"shared/pb/mediapb"
	"shared/pb/seasonpb"
	"shared/ports"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type App struct {
	episodepb.UnimplementedEpisodeServiceServer
	repository   Repository
	idGenerator  ports.IDGenerator
	mediaClient  mediapb.MediaServiceClient
	seasonClient seasonpb.SeasonServiceClient
}

func New(
	repository Repository,
	idGenerator ports.IDGenerator,
	mediaClient mediapb.MediaServiceClient,
	seasonClient seasonpb.SeasonServiceClient,
) *App {
	return &App{
		repository:   repository,
		idGenerator:  idGenerator,
		mediaClient:  mediaClient,
		seasonClient: seasonClient,
	}
}

func (a *App) CreateEpisode(ctx context.Context, req *episodepb.CreateEpisodeRequest) (*episodepb.CreateEpisodeResponse, error) {
	_, err := a.seasonClient.GetSeasonByID(ctx, &seasonpb.GetSeasonByIDRequest{SeasonId: req.SeasonId})
	if err != nil {
		return nil, fmt.Errorf("seasonClient.GetSeasonByID: %w", err)
	}
	err = a.validateMediaID(ctx, req.MediaId)
	if err != nil {
		return nil, fmt.Errorf("validateMediaID: %w", err)
	}
	order, err := a.generateOrderNumber(ctx, req.SeasonId)
	if err != nil {
		return nil, fmt.Errorf("generateOrderNumber: %w", err)
	}
	id := a.idGenerator.NewID()
	err = a.repository.CreateEpisode(ctx, &models.Episode{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		SeasonID:    req.SeasonId,
		MediaID:     req.MediaId,
		Order:       order,
	})
	if err != nil {
		return nil, fmt.Errorf("repository.CreateEpisode: %w", err)
	}
	return &episodepb.CreateEpisodeResponse{
		EpisodeId: id,
	}, nil
}

func (a *App) DeleteEpisodeByID(ctx context.Context, req *episodepb.DeleteEpisodeByIDRequest) (*emptypb.Empty, error) {
	episode, err := a.repository.GetEpisodeByID(ctx, req.EpisodeId)
	if err != nil {
		return nil, fmt.Errorf("repository.GetEpisodeByID: %w", err)
	}
	err = a.repository.DeleteEpisodeByID(ctx, req.EpisodeId)
	if err != nil {
		return nil, fmt.Errorf("repository.DeleteEpisodeByID: %w", err)
	}
	episodes, err := a.repository.ListEpisodesBySeasonID(ctx, episode.SeasonID)
	if err != nil {
		return nil, fmt.Errorf("repository.ListEpisodesBySeasonID: %w", err)
	}
	if len(episodes) == 0 {
		return &emptypb.Empty{}, nil
	}
	_, err = a.ReorderEpisodesBySeasonID(ctx, &episodepb.ReorderEpisodesBySeasonIDRequest{
		SeasonId:   episode.SeasonID,
		EpisodeIds: lo.Map(episodes, func(e *models.Episode, _ int) string { return e.ID }),
	})
	if err != nil {
		return nil, fmt.Errorf("ReorderEpisodesBySeasonID: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (a *App) GetEpisodeByID(ctx context.Context, req *episodepb.GetEpisodeByIDRequest) (*episodepb.Episode, error) {
	episode, err := a.repository.GetEpisodeByID(ctx, req.EpisodeId)
	if err != nil {
		return nil, fmt.Errorf("repository.GetEpisodeByID: %w", err)
	}
	var mediaInfo *episodepb.Media
	if episode.MediaID != "" {
		resp, err := a.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{MediaId: episode.MediaID})
		if err != nil {
			return nil, fmt.Errorf("mediaClient.GetMediaByID: %w", err)
		}
		mediaInfo = &episodepb.Media{
			Id:        resp.Id,
			CreatedAt: resp.CreatedAt,
			UpdatedAt: resp.UpdatedAt,
			Title:     resp.Title,
			Path:      resp.Path,
			Type:      episodepb.MediaType(resp.Type),
			MimeType:  resp.MimeType,
			Size:      resp.Size,
		}
	}
	return &episodepb.Episode{
		Id:          episode.ID,
		CreatedAt:   timestamppb.New(episode.CreatedAt),
		UpdatedAt:   timestamppb.New(episode.UpdatedAt),
		Title:       episode.Title,
		Description: episode.Description,
		SeasonId:    episode.SeasonID,
		Order:       int32(episode.Order),
		MediaInfo:   mediaInfo,
	}, nil
}

func (a *App) ListEpisodesBySeasonID(ctx context.Context, req *episodepb.ListEpisodesBySeasonIDRequest) (*episodepb.EpisodeList, error) {
	episodes, err := a.repository.ListEpisodesBySeasonID(ctx, req.SeasonId)
	if err != nil {
		return nil, fmt.Errorf("repository.ListEpisodesBySeasonID: %w", err)
	}

	resultList := make([]*episodepb.Episode, 0, len(episodes))
	for _, episode := range episodes {
		var mediaInfo *episodepb.Media
		if episode.MediaID != "" {
			resp, err := a.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{MediaId: episode.MediaID})
			if err != nil {
				return nil, fmt.Errorf("mediaClient.GetMediaByID: %w", err)
			}
			mediaInfo = &episodepb.Media{
				Id:        resp.Id,
				CreatedAt: resp.CreatedAt,
				UpdatedAt: resp.UpdatedAt,
				Title:     resp.Title,
				Path:      resp.Path,
				Type:      episodepb.MediaType(resp.Type),
				MimeType:  resp.MimeType,
				Size:      resp.Size,
			}
		}
		resultList = append(resultList, &episodepb.Episode{
			Id:          episode.ID,
			CreatedAt:   timestamppb.New(episode.CreatedAt),
			UpdatedAt:   timestamppb.New(episode.UpdatedAt),
			Title:       episode.Title,
			Description: episode.Description,
			SeasonId:    episode.SeasonID,
			Order:       int32(episode.Order),
			MediaInfo:   mediaInfo,
		})
	}
	return &episodepb.EpisodeList{
		List: resultList,
	}, nil
}

func (a *App) ReorderEpisodesBySeasonID(ctx context.Context, req *episodepb.ReorderEpisodesBySeasonIDRequest) (*emptypb.Empty, error) {
	episodes, err := a.repository.ListEpisodesBySeasonID(ctx, req.SeasonId)
	if err != nil {
		return nil, fmt.Errorf("repository.ListEpisodesBySeasonID: %w", err)
	}
	if len(lo.Intersect(req.EpisodeIds, lo.Map(episodes, func(e *models.Episode, _ int) string {
		return e.ID
	}))) != len(episodes) {
		return nil, fmt.Errorf("incorrect episode ids")
	}

	for i, episode := range episodes {
		if episode.ID == req.EpisodeIds[i] {
			continue
		}
		index := lo.IndexOf(req.EpisodeIds, episode.ID)
		episode.Order = int32(index + 1)
		err = a.repository.UpdateEpisodeOrder(ctx, episode)
		if err != nil {
			return nil, fmt.Errorf("repository.UpdateEpisodeOrder: %w", err)
		}
	}
	return &emptypb.Empty{}, nil
}

func (a *App) UpdateEpisodeByID(ctx context.Context, req *episodepb.UpdateEpisodeByIDRequest) (*emptypb.Empty, error) {
	err := a.validateMediaID(ctx, req.MediaId)
	if err != nil {
		return nil, fmt.Errorf("validateMediaID: %w", err)
	}
	err = a.repository.UpdateEpisodeByID(ctx, &models.Episode{
		ID:          req.EpisodeId,
		Title:       req.Title,
		Description: req.Description,
		MediaID:     req.MediaId,
	})
	if err != nil {
		return nil, fmt.Errorf("repository.UpdateEpisodeByID: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (a *App) DeleteAllEpisodesBySeasonID(ctx context.Context, req *episodepb.DeleteAllEpisodesBySeasonIDRequest) (*emptypb.Empty, error) {
	err := a.repository.DeleteAllEpisodesBySeasonID(ctx, req.SeasonId)
	if err != nil {
		return nil, fmt.Errorf("repository.DeleteAllEpisodesBySeasonID: %w", err)
	}
	return &emptypb.Empty{}, nil
}
