package episode

import (
	"bff-service/internal/core/models"
	"context"
	"fmt"
	"net/http"
	"shared/enums/mediatype"
	"shared/pb/episodepb"

	"github.com/samber/lo"
)

type ListEpisodes struct {
	episodeClient episodepb.EpisodeServiceClient
}

func NewListEpisodes(episodeClient episodepb.EpisodeServiceClient) *ListEpisodes {
	return &ListEpisodes{
		episodeClient: episodeClient,
	}
}

type ListEpisodesRequest struct {
	SeasonID string `query:"season_id"`
}

type ListEpisodesResponse struct {
	List []*models.Episode `json:"list"`
}

func (h *ListEpisodes) Handle(ctx context.Context, req *ListEpisodesRequest) (*ListEpisodesResponse, int, error) {
	episodes, err := h.episodeClient.ListEpisodesBySeasonID(ctx, &episodepb.ListEpisodesBySeasonIDRequest{
		SeasonId: req.SeasonID,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("episodeClient.ListEpisodesBySeasonID: %w", err)
	}
	return &ListEpisodesResponse{
		List: lo.Map(episodes.List, func(e *episodepb.Episode, _ int) *models.Episode {
			var mediaInfo *models.Media
			if e.MediaInfo != nil {
				mediaInfo = &models.Media{
					ID:        e.MediaInfo.Id,
					CreatedAt: e.MediaInfo.CreatedAt.AsTime(),
					UpdatedAt: e.MediaInfo.UpdatedAt.AsTime(),
					Title:     e.MediaInfo.Title,
					Path:      e.MediaInfo.Path,
					Type:      mediatype.FromString(e.MediaInfo.Type),
					MimeType:  e.MediaInfo.MimeType,
					Size:      e.MediaInfo.Size,
				}
			}
			return &models.Episode{
				ID:          e.Id,
				CreatedAt:   e.CreatedAt.AsTime(),
				UpdatedAt:   e.UpdatedAt.AsTime(),
				Title:       e.Title,
				Description: e.Description,
				Order:       e.Order,
				SeasonID:    e.SeasonId,
				MediaInfo:   mediaInfo,
			}
		}),
	}, http.StatusOK, nil
}
