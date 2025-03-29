package media

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/mediapb"
	"strings"
)

type DeleteMediaByID struct {
	mediaClient mediapb.MediaServiceClient
}

func NewDeleteMediaByID(mediaClient mediapb.MediaServiceClient) *DeleteMediaByID {
	return &DeleteMediaByID{
		mediaClient: mediaClient,
	}
}

type DeleteMediaByIDRequest struct {
	MediaID string `params:"media_id"`
}

func (h *DeleteMediaByID) Handle(ctx context.Context, req *DeleteMediaByIDRequest) (*any, int, error) {
	_, err := h.mediaClient.DeleteMediaByID(ctx, &mediapb.DeleteMediaByIDRequest{
		MediaId: req.MediaID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "media not found") {
			return nil, http.StatusNotFound, errors.New("media not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("mediaClient.DeleteMediaByID: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
