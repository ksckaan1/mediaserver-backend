package episode

import (
	"bff-service/internal/core/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/enums/mediatype"
	"shared/pb/episodepb"
	"strings"
)

type GetEpisodeByID struct {
	episodeClient episodepb.EpisodeServiceClient
}

func NewGetEpisodeByID(episodeClient episodepb.EpisodeServiceClient) *GetEpisodeByID {
	return &GetEpisodeByID{
		episodeClient: episodeClient,
	}
}

type GetEpisodeByIDRequest struct {
	EpisodeID string `params:"episode_id"`
}

func (h *GetEpisodeByID) Handle(ctx context.Context, req *GetEpisodeByIDRequest) (*models.Episode, int, error) {
	resp, err := h.episodeClient.GetEpisodeByID(ctx, &episodepb.GetEpisodeByIDRequest{
		EpisodeId: req.EpisodeID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "episode not found") {
			return nil, http.StatusNotFound, errors.New("episode not found")
		}
		if strings.Contains(err.Error(), "media not found") {
			return nil, http.StatusNotFound, errors.New("media not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("episodeClient.GetEpisodeByID: %w", err)
	}

	var mediaInfo *models.Media
	if resp.MediaInfo != nil {
		mediaInfo = &models.Media{
			ID:        resp.MediaInfo.Id,
			CreatedAt: resp.MediaInfo.CreatedAt.AsTime(),
			UpdatedAt: resp.MediaInfo.UpdatedAt.AsTime(),
			Title:     resp.MediaInfo.Title,
			Path:      resp.MediaInfo.Path,
			Type:      mediatype.FromString(resp.MediaInfo.Type),
			MimeType:  resp.MediaInfo.MimeType,
			Size:      resp.MediaInfo.Size,
		}
	}

	return &models.Episode{
		ID:          resp.Id,
		CreatedAt:   resp.CreatedAt.AsTime(),
		UpdatedAt:   resp.UpdatedAt.AsTime(),
		Title:       resp.Title,
		Description: resp.Description,
		Order:       resp.Order,
		SeasonID:    resp.SeasonId,
		MediaInfo:   mediaInfo,
	}, http.StatusOK, nil
}
