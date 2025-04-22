package media

import (
	"bff-service/internal/core/models"
	"context"
	"fmt"
	"net/http"
	"shared/enums/mediatype"
	"shared/pb/mediapb"

	"github.com/samber/lo"
)

type ListMedias struct {
	mediaClient mediapb.MediaServiceClient
}

func NewListMedias(mediaClient mediapb.MediaServiceClient) *ListMedias {
	return &ListMedias{
		mediaClient: mediaClient,
	}
}

type ListMediasRequest struct {
	Limit  *int64 `query:"limit"`
	Offset int64  `query:"offset"`
}

type ListMediasResponse struct {
	List   []*models.Media `json:"list"`
	Count  int64           `json:"count"`
	Limit  int64           `json:"limit"`
	Offset int64           `json:"offset"`
}

func (h *ListMedias) Handle(ctx context.Context, req *ListMediasRequest) (*ListMediasResponse, int, error) {
	var limit int64 = 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	resp, err := h.mediaClient.ListMedias(ctx, &mediapb.ListMediasRequest{
		Limit:  limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("mediaClient.ListMedias: %w", err)
	}
	return &ListMediasResponse{
		List: lo.Map(resp.List, func(m *mediapb.Media, _ int) *models.Media {
			return &models.Media{
				ID:        m.Id,
				CreatedAt: m.CreatedAt.AsTime(),
				UpdatedAt: m.UpdatedAt.AsTime(),
				Title:     m.Title,
				Path:      m.Path,
				Type:      mediatype.FromString(m.Type),
				MimeType:  m.MimeType,
				Size:      m.Size,
			}
		}),
		Count:  resp.Count,
		Limit:  resp.Limit,
		Offset: resp.Offset,
	}, http.StatusOK, nil
}
