package media

import (
	"bff-service/internal/core/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/enums/mediatype"
	"shared/pb/mediapb"
	"strings"
)

type GetMediaByID struct {
	mediaClient mediapb.MediaServiceClient
}

func NewGetMediaByID(mediaClient mediapb.MediaServiceClient) *GetMediaByID {
	return &GetMediaByID{
		mediaClient: mediaClient,
	}
}

type GetMediaByIDRequest struct {
	MediaID string `params:"media_id"`
}

func (h *GetMediaByID) Handle(ctx context.Context, req *GetMediaByIDRequest) (*models.Media, int, error) {
	resp, err := h.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{
		MediaId: req.MediaID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "media not found") {
			return nil, http.StatusNotFound, errors.New("media not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("mediaClient.GetMediaByID: %w", err)
	}
	return &models.Media{
		ID:        resp.Id,
		CreatedAt: resp.CreatedAt.AsTime(),
		UpdatedAt: resp.UpdatedAt.AsTime(),
		Title:     resp.Title,
		Path:      resp.Path,
		Type:      mediatype.FromNumber(int32(resp.Type)),
		MimeType:  resp.MimeType,
		Size:      resp.Size,
	}, http.StatusOK, nil
}
