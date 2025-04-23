package setting

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"shared/pb/settingpb"
)

type SetSetting struct {
	settingClient settingpb.SettingServiceClient
}

func NewSetSetting(settingClient settingpb.SettingServiceClient) *SetSetting {
	return &SetSetting{settingClient: settingClient}
}

type SetSettingRequest struct {
	Key   string          `params:"key"`
	Value json.RawMessage `json:"value"`
}

func (h *SetSetting) Handle(ctx context.Context, req *SetSettingRequest) (*any, int, error) {
	_, err := h.settingClient.Set(ctx, &settingpb.SetRequest{
		Key:   req.Key,
		Value: req.Value,
	})
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("settingClient.Set: %w", err)
	}
	return nil, http.StatusNoContent, nil
}
