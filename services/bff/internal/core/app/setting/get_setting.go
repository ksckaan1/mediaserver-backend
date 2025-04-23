package setting

import (
	"bff-service/internal/core/models"
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/settingpb"
	"strings"
)

type GetSetting struct {
	settingClient settingpb.SettingServiceClient
}

func NewGetSetting(settingClient settingpb.SettingServiceClient) *GetSetting {
	return &GetSetting{settingClient: settingClient}
}

type GetSettingRequest struct {
	Key string `params:"key"`
}

func (h *GetSetting) Handle(ctx context.Context, req *GetSettingRequest) (*models.Setting, int, error) {
	setting, err := h.settingClient.Get(ctx, &settingpb.GetRequest{
		Key: req.Key,
	})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, http.StatusNotFound, errors.New("setting not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("settingClient.Get: %w", err)
	}
	return &models.Setting{
		Key:   setting.Key,
		Value: setting.Value,
	}, http.StatusOK, nil
}
