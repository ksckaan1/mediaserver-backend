package media

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/mediapb"
	"strings"
)

type UpdateMediaByID struct {
	mediaClient mediapb.MediaServiceClient
}

func NewUpdateMediaByID(mediaClient mediapb.MediaServiceClient) *UpdateMediaByID {
	return &UpdateMediaByID{
		mediaClient: mediaClient,
	}
}

type UpdateMediaByIDRequest struct {
	MediaID string `params:"media_id"`
	Title   string `json:"title"`
}

func (h *UpdateMediaByID) Handle(ctx context.Context, req *UpdateMediaByIDRequest) (*any, int, error) {
	_, err := h.mediaClient.UpdateMediaByID(ctx, &mediapb.UpdateMediaByIDRequest{
		MediaId: req.MediaID,
		Title:   req.Title,
	})
	if err != nil {
		if strings.Contains(err.Error(), "media not found") {
			return nil, http.StatusNotFound, errors.New("media not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("mediaClient.UpdateMediaByID: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
