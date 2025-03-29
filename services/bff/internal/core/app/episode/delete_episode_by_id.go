package episode

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/episodepb"
	"strings"
)

type DeleteEpisodeByID struct {
	episodeClient episodepb.EpisodeServiceClient
}

func NewDeleteEpisodeByID(episodeClient episodepb.EpisodeServiceClient) *DeleteEpisodeByID {
	return &DeleteEpisodeByID{
		episodeClient: episodeClient,
	}
}

type DeleteEpisodeByIDRequest struct {
	EpisodeID string `params:"episode_id"`
}

func (h *DeleteEpisodeByID) Handle(ctx context.Context, req *DeleteEpisodeByIDRequest) (*any, int, error) {
	_, err := h.episodeClient.DeleteEpisodeByID(ctx, &episodepb.DeleteEpisodeByIDRequest{
		EpisodeId: req.EpisodeID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "episode not found") {
			return nil, http.StatusNotFound, errors.New("episode not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("episodeClient.DeleteEpisodeByID: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
