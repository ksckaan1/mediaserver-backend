package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"mediaserver/internal/customerrors"
	"mediaserver/internal/domain/core/model"
	"mediaserver/internal/infrastructure/repository/sqlcgen"
)

func (m *Repository) GetTMDBInfoByID(ctx context.Context, id string) (*model.TMDBInfo, error) {
	info, err := m.queries.GetTMDBInfo(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("queries.GetTMDBInfo: %w", customerrors.ErrRecordNotFound)
		}
		return nil, fmt.Errorf("queries.GetTMDBInfo: %w", err)
	}
	var data model.TMDBData
	if len(info.Data) > 0 {
		err = json.Unmarshal(info.Data, &data)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %w", err)
		}
	}
	return &model.TMDBInfo{
		ID:   id,
		Data: data,
	}, nil
}

func (m *Repository) SetTMDBInfo(ctx context.Context, info *model.TMDBInfo) error {
	data, err := json.Marshal(info.Data)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}
	err = m.queries.SetTMDBInfo(ctx, sqlcgen.SetTMDBInfoParams{
		ID:   info.ID,
		Data: data,
	})
	if err != nil {
		return fmt.Errorf("queries.SetTMDBInfo: %w", err)
	}
	return nil
}
