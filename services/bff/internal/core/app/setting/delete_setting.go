package setting

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shared/pb/settingpb"
	"strings"
)

type DeleteSetting struct {
	settingClient settingpb.SettingServiceClient
}

func NewDeleteSetting(settingClient settingpb.SettingServiceClient) *DeleteSetting {
	return &DeleteSetting{settingClient: settingClient}
}

type DeleteSettingRequest struct {
	Key string `params:"key"`
}

func (h *DeleteSetting) Handle(ctx context.Context, req *DeleteSettingRequest) (*any, int, error) {
	_, err := h.settingClient.Delete(ctx, &settingpb.DeleteRequest{
		Key: req.Key,
	})
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, http.StatusNotFound, errors.New("setting not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("settingClient.Delete: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
