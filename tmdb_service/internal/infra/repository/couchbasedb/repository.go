package couchbasedb

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tmdb_service/internal/core/customerrors"
	"tmdb_service/internal/core/models"

	"github.com/couchbase/gocb/v2"
)

type Repository struct {
	coll *gocb.Collection
}

func New(bucket *gocb.Bucket) *Repository {
	return &Repository{
		coll: bucket.Scope("tmdb_service").Collection("infos"),
	}
}

func (r *Repository) GetTMDBInfo(ctx context.Context, id string) (*models.TMDBInfo, error) {
	result, err := r.coll.Get(id, &gocb.GetOptions{})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, fmt.Errorf("coll.Get: %w", err)
	}
	var info models.TMDBInfo
	if err := result.Content(&info); err != nil {
		return nil, fmt.Errorf("result.Content: %w", err)
	}
	return &info, nil
}

func (r *Repository) SetTMDBInfo(ctx context.Context, info *models.TMDBInfo) error {
	info.UpdatedAt = time.Now()
	_, err := r.coll.Upsert(info.Id, info, &gocb.UpsertOptions{
		Expiry: 24 * time.Hour,
	})
	if err != nil {
		return fmt.Errorf("coll.Upsert: %w", err)
	}
	return nil
}
