package episode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/episodepb"
	"strings"
)

type UpdateEpisodeByID struct {
	episodeClient episodepb.EpisodeServiceClient
}

func NewUpdateEpisodeByID(episodeClient episodepb.EpisodeServiceClient) *UpdateEpisodeByID {
	return &UpdateEpisodeByID{
		episodeClient: episodeClient,
	}
}

type UpdateEpisodeByIDRequest struct {
	EpisodeID   string `params:"episode_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	MediaID     string `json:"media_id"`
}

func (h *UpdateEpisodeByID) Handle(ctx context.Context, req *UpdateEpisodeByIDRequest) (*any, int, error) {
	_, err := h.episodeClient.UpdateEpisodeByID(ctx, &episodepb.UpdateEpisodeByIDRequest{
		EpisodeId:   req.EpisodeID,
		Title:       req.Title,
		Description: req.Description,
		MediaId:     req.MediaID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "episode not found") {
			return nil, http.StatusNotFound, errors.New("episode not found")
		}
		if strings.Contains(err.Error(), "media not found") {
			return nil, http.StatusNotFound, errors.New("media not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("episodeClient.UpdateEpisodeByID: %w", err)
	}
	return nil, http.StatusOK, nil
}
