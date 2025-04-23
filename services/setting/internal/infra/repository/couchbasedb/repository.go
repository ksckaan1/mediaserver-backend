package couchbasedb

import (
	"context"
	"errors"
	"fmt"
	"setting_service/internal/core/customerrors"
	"setting_service/internal/core/models"

	"github.com/couchbase/gocb/v2"
	"github.com/samber/lo"
)

type Repository struct {
	coll  *gocb.Collection
	scope *gocb.Scope
}

func New(bucket *gocb.Bucket) *Repository {
	return &Repository{
		coll:  bucket.Scope("setting_service").Collection("settings"),
		scope: bucket.Scope("setting_service"),
	}
}

func (r *Repository) Set(ctx context.Context, setting *models.Setting) error {
	_, err := r.coll.Upsert(setting.Key, setting, &gocb.UpsertOptions{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("repository.Set: %w", err)
	}
	return nil
}

func (r *Repository) Get(ctx context.Context, key string) (*models.Setting, error) {
	result, err := r.coll.Get(key, &gocb.GetOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, customerrors.ErrSettingNotFound
		}
		return nil, fmt.Errorf("r.coll.Get: %w", err)
	}
	var setting models.Setting
	err = result.Content(&setting)
	if err != nil {
		return nil, fmt.Errorf("result.Content: %w", err)
	}
	return &setting, nil
}

const (
	queryListCount     = "SELECT COUNT(*) AS count FROM settings"
	queryList          = "SELECT * FROM settings OFFSET $offset"
	queryListWithLimit = "SELECT * FROM settings LIMIT $limit OFFSET $offset"
)

type countResponse struct {
	Count int64 `json:"count"`
}

type listResponse struct {
	Settings *models.Setting `json:"settings"`
}

func (r *Repository) List(ctx context.Context, limit int64, offset int64) (*models.SettingList, error) {
	countCur, err := r.scope.Query(queryListCount, &gocb.QueryOptions{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("r.scope.Query: %w", err)
	}
	var countResponse countResponse
	err = countCur.One(&countResponse)
	if err != nil {
		return nil, fmt.Errorf("countCur.One: %w", err)
	}
	if countResponse.Count == 0 || limit == 0 {
		return &models.SettingList{
			List:   []*models.Setting{},
			Count:  countResponse.Count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	cursor, err := r.scope.Query(lo.Ternary(limit < 0, queryList, queryListWithLimit), &gocb.QueryOptions{
		Context: ctx,
		NamedParameters: map[string]any{
			"limit":  limit,
			"offset": offset,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("r.scope.Query: %w", err)
	}
	var settings []*models.Setting
	for cursor.Next() {
		var result listResponse
		err = cursor.Row(&result)
		if err != nil {
			return nil, fmt.Errorf("cursor.Row: %w", err)
		}
		settings = append(settings, result.Settings)
	}
	return &models.SettingList{
		List:   settings,
		Count:  countResponse.Count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (r *Repository) Delete(ctx context.Context, key string) error {
	_, err := r.coll.Remove(key, &gocb.RemoveOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return customerrors.ErrSettingNotFound
		}
		return fmt.Errorf("r.coll.Remove: %w", err)
	}
	return nil
}
