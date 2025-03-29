package episode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/episodepb"
	"strings"
)

type CreateEpisode struct {
	episodeClient episodepb.EpisodeServiceClient
}

func NewCreateEpisode(episodeClient episodepb.EpisodeServiceClient) *CreateEpisode {
	return &CreateEpisode{
		episodeClient: episodeClient,
	}
}

type CreateEpisodeRequest struct {
	SeasonID    string `json:"season_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	MediaID     string `json:"media_id"`
}

type CreateEpisodeResponse struct {
	EpisodeID string `json:"episode_id"`
}

func (h *CreateEpisode) Handle(ctx context.Context, req *CreateEpisodeRequest) (*CreateEpisodeResponse, int, error) {
	resp, err := h.episodeClient.CreateEpisode(ctx, &episodepb.CreateEpisodeRequest{
		SeasonId:    req.SeasonID,
		Title:       req.Title,
		Description: req.Description,
		MediaId:     req.MediaID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "season not found") {
			return nil, http.StatusNotFound, errors.New("season not found")
		}
		if strings.Contains(err.Error(), "media not found") {
			return nil, http.StatusNotFound, errors.New("media not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("episodeClient.CreateEpisode: %w", err)
	}
	return &CreateEpisodeResponse{
		EpisodeID: resp.EpisodeId,
	}, http.StatusCreated, nil
}
