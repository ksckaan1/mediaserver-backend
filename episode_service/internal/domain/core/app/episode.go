package app

import (
	"common/pb/episodepb"
	"common/pb/mediapb"
	"context"
	"episode_service/internal/domain/core/models"
	"fmt"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Episode struct {
	episodepb.UnimplementedEpisodeServiceServer
	repository  Repository
	idGenerator IDGenerator
	mediaClient mediapb.MediaServiceClient
}

func New(repository Repository, idGenerator IDGenerator, mediaClient mediapb.MediaServiceClient) *Episode {
	return &Episode{
		repository:  repository,
		idGenerator: idGenerator,
		mediaClient: mediaClient,
	}
}

func (h *Episode) CreateEpisode(ctx context.Context, req *episodepb.CreateEpisodeRequest) (*episodepb.CreateEpisodeResponse, error) {
	err := h.validateMediaID(ctx, req.MediaId)
	if err != nil {
		return nil, fmt.Errorf("validateMediaID: %w", err)
	}
	order, err := h.generateOrderNumber(ctx, req.SeasonId)
	if err != nil {
		return nil, fmt.Errorf("generateOrderNumber: %w", err)
	}
	id := h.idGenerator.NewID()
	err = h.repository.CreateEpisode(ctx, &models.Episode{
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

func (h *Episode) DeleteEpisodeByID(ctx context.Context, req *episodepb.DeleteEpisodeByIDRequest) (*emptypb.Empty, error) {
	episode, err := h.repository.GetEpisodeByID(ctx, req.EpisodeId)
	if err != nil {
		return nil, fmt.Errorf("repository.GetEpisodeByID: %w", err)
	}
	err = h.repository.DeleteEpisodeByID(ctx, req.EpisodeId)
	if err != nil {
		return nil, fmt.Errorf("repository.DeleteEpisodeByID: %w", err)
	}
	episodes, err := h.repository.ListEpisodesBySeasonID(ctx, episode.SeasonID)
	if err != nil {
		return nil, fmt.Errorf("repository.ListEpisodesBySeasonID: %w", err)
	}
	if len(episodes) == 0 {
		return &emptypb.Empty{}, nil
	}
	_, err = h.ReorderEpisodesBySeasonID(ctx, &episodepb.ReorderEpisodesBySeasonIDRequest{
		SeasonId:   episode.SeasonID,
		EpisodeIds: lo.Map(episodes, func(e *models.Episode, _ int) string { return e.ID }),
	})
	if err != nil {
		return nil, fmt.Errorf("ReorderEpisodesBySeasonID: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (h *Episode) GetEpisodeByID(ctx context.Context, req *episodepb.GetEpisodeByIDRequest) (*episodepb.Episode, error) {
	episode, err := h.repository.GetEpisodeByID(ctx, req.EpisodeId)
	if err != nil {
		return nil, fmt.Errorf("repository.GetEpisodeByID: %w", err)
	}
	return &episodepb.Episode{
		Id:          episode.ID,
		CreatedAt:   timestamppb.New(episode.CreatedAt),
		UpdatedAt:   timestamppb.New(episode.UpdatedAt),
		Title:       episode.Title,
		Description: episode.Description,
		SeasonId:    episode.SeasonID,
		MediaId:     episode.MediaID,
		Order:       int32(episode.Order),
	}, nil
}

func (h *Episode) ListEpisodesBySeasonID(ctx context.Context, req *episodepb.ListEpisodesBySeasonIDRequest) (*episodepb.EpisodeList, error) {
	episodes, err := h.repository.ListEpisodesBySeasonID(ctx, req.SeasonId)
	if err != nil {
		return nil, fmt.Errorf("repository.ListEpisodesBySeasonID: %w", err)
	}
	return &episodepb.EpisodeList{
		List: lo.Map(episodes, func(e *models.Episode, _ int) *episodepb.Episode {
			return &episodepb.Episode{
				Id:          e.ID,
				CreatedAt:   timestamppb.New(e.CreatedAt),
				UpdatedAt:   timestamppb.New(e.UpdatedAt),
				Title:       e.Title,
				Description: e.Description,
				SeasonId:    e.SeasonID,
				MediaId:     e.MediaID,
				Order:       int32(e.Order),
			}
		}),
	}, nil
}

func (h *Episode) ReorderEpisodesBySeasonID(ctx context.Context, req *episodepb.ReorderEpisodesBySeasonIDRequest) (*emptypb.Empty, error) {
	episodes, err := h.repository.ListEpisodesBySeasonID(ctx, req.SeasonId)
	if err != nil {
		return nil, fmt.Errorf("repository.ListEpisodesBySeasonID: %w", err)
	}
	if len(req.EpisodeIds) != len(episodes) {
		return nil, fmt.Errorf("invalid episode ids length")
	}

	for i, episode := range episodes {
		if episode.ID == req.EpisodeIds[i] {
			continue
		}
		index := lo.IndexOf(req.EpisodeIds, episode.ID)
		episode.Order = int32(index + 1)
		err = h.repository.UpdateEpisodeOrder(ctx, episode)
		if err != nil {
			return nil, fmt.Errorf("repository.UpdateEpisodeOrder: %w", err)
		}
	}
	return &emptypb.Empty{}, nil
}

func (h *Episode) UpdateEpisodeByID(ctx context.Context, req *episodepb.UpdateEpisodeByIDRequest) (*emptypb.Empty, error) {
	err := h.validateMediaID(ctx, req.MediaId)
	if err != nil {
		return nil, fmt.Errorf("validateMediaID: %w", err)
	}
	err = h.repository.UpdateEpisodeByID(ctx, &models.Episode{
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

func (h *Episode) DeleteAllEpisodesBySeasonID(ctx context.Context, req *episodepb.DeleteAllEpisodesBySeasonIDRequest) (*emptypb.Empty, error) {
	err := h.repository.DeleteAllEpisodesBySeasonID(ctx, req.SeasonId)
	if err != nil {
		return nil, fmt.Errorf("repository.DeleteAllEpisodesBySeasonID: %w", err)
	}
	return &emptypb.Empty{}, nil
}
