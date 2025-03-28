package media

import (
	"context"
	"fmt"
	"shared/pb/mediapb"
	"time"
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
	MediaID string `param:"media_id"`
}

type GetMediaByIDResponse struct {
	ID        string            `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Title     string            `json:"title"`
	Path      string            `json:"path"`
	Type      mediapb.MediaType `json:"type"`
	MimeType  string            `json:"mime_type"`
	Size      int64             `json:"size"`
}

func (h *GetMediaByID) Handle(ctx context.Context, req *GetMediaByIDRequest) (*GetMediaByIDResponse, error) {
	resp, err := h.mediaClient.GetMediaByID(ctx, &mediapb.GetMediaByIDRequest{
		MediaId: req.MediaID,
	})
	if err != nil {
		return nil, fmt.Errorf("mediaClient.GetMediaByID: %w", err)
	}
	return &GetMediaByIDResponse{
		ID:        resp.Id,
		CreatedAt: resp.CreatedAt.AsTime(),
		UpdatedAt: resp.UpdatedAt.AsTime(),
		Title:     resp.Title,
		Path:      resp.Path,
		Type:      resp.Type,
		MimeType:  resp.MimeType,
		Size:      resp.Size,
	}, nil
}
