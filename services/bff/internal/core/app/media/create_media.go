package media

import (
	"context"
	"fmt"
	"net/http"
	"shared/pb/mediapb"
)

type CreateMedia struct {
	mediaClient mediapb.MediaServiceClient
}

func NewCreateMedia(mediaClient mediapb.MediaServiceClient) *CreateMedia {
	return &CreateMedia{
		mediaClient: mediaClient,
	}
}

type CreateMediaRequest struct {
}

type CreateMediaResponse struct {
	MediaID      string            `json:"media_id"`
	PresignedURL string            `json:"presigned_url"`
	FormData     map[string]string `json:"form_data"`
}

func (h *CreateMedia) Handle(ctx context.Context, req *CreateMediaRequest) (*CreateMediaResponse, int, error) {
	resp, err := h.mediaClient.CreateMedia(ctx, &mediapb.CreateMediaRequest{})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("mediaClient.CreateMedia: %w", err)
	}
	return &CreateMediaResponse{
		MediaID:      resp.MediaId,
		PresignedURL: resp.PresignedUrl,
		FormData:     resp.FormData,
	}, http.StatusCreated, nil
}
