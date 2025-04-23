package setting

import (
	"bff-service/internal/core/models"
	"context"
	"fmt"
	"net/http"
	"shared/pb/settingpb"

	"github.com/samber/lo"
)

type ListSettings struct {
	settingClient settingpb.SettingServiceClient
}

func NewListSettings(settingClient settingpb.SettingServiceClient) *ListSettings {
	return &ListSettings{settingClient: settingClient}
}

type ListSettingsRequest struct {
	Limit  *int64 `query:"limit"`
	Offset int64  `query:"offset"`
}

type ListSettingsResponse struct {
	List   []*models.Setting `json:"list"`
	Count  int64             `json:"count"`
	Limit  int64             `json:"limit"`
	Offset int64             `json:"offset"`
}

func (h *ListSettings) Handle(ctx context.Context, req *ListSettingsRequest) (*ListSettingsResponse, int, error) {
	var limit int64 = 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	settings, err := h.settingClient.List(ctx, &settingpb.ListRequest{
		Limit:  limit,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("settingClient.List: %w", err)
	}
	return &ListSettingsResponse{
		List: lo.Map(settings.List, func(setting *settingpb.Setting, _ int) *models.Setting {
			return &models.Setting{
				Key:   setting.Key,
				Value: setting.Value,
			}
		}),
		Count:  settings.Count,
		Limit:  limit,
		Offset: req.Offset,
	}, http.StatusOK, nil
}
