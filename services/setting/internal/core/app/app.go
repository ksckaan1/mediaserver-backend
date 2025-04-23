package app

import (
	"context"
	"fmt"
	"setting_service/internal/core/models"
	"shared/pb/settingpb"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ settingpb.SettingServiceServer = (*App)(nil)

type App struct {
	settingpb.UnimplementedSettingServiceServer
	repository Repository
}

func New(repository Repository) *App {
	return &App{
		repository: repository,
	}
}

func (a *App) Get(ctx context.Context, req *settingpb.GetRequest) (*settingpb.Setting, error) {
	setting, err := a.repository.Get(ctx, req.Key)
	if err != nil {
		return nil, fmt.Errorf("repository.Get: %w", err)
	}
	return &settingpb.Setting{
		Key:   setting.Key,
		Value: setting.Value,
	}, nil
}

func (a *App) List(ctx context.Context, req *settingpb.ListRequest) (*settingpb.ListResponse, error) {
	settings, err := a.repository.List(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, fmt.Errorf("repository.List: %w", err)
	}
	return &settingpb.ListResponse{
		List: lo.Map(settings.List, func(setting *models.Setting, _ int) *settingpb.Setting {
			return &settingpb.Setting{
				Key:   setting.Key,
				Value: setting.Value,
			}
		}),
		Count:  settings.Count,
		Limit:  settings.Limit,
		Offset: settings.Offset,
	}, nil
}

func (a *App) Set(ctx context.Context, req *settingpb.SetRequest) (*emptypb.Empty, error) {
	setting := &models.Setting{
		Key:   req.Key,
		Value: req.Value,
	}
	err := a.repository.Set(ctx, setting)
	if err != nil {
		return nil, fmt.Errorf("repository.Set: %w", err)
	}
	return &emptypb.Empty{}, nil
}

func (a *App) Delete(ctx context.Context, req *settingpb.DeleteRequest) (*emptypb.Empty, error) {
	err := a.repository.Delete(ctx, req.Key)
	if err != nil {
		return nil, fmt.Errorf("repository.Delete: %w", err)
	}
	return &emptypb.Empty{}, nil
}
