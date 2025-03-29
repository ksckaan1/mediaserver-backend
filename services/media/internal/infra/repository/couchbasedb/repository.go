package couchbasedb

import (
	"context"
	"errors"
	"fmt"
	"media_service/internal/core/customerrors"
	"media_service/internal/core/models"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/samber/lo"
)

type Repository struct {
	scope *gocb.Scope
	coll  *gocb.Collection
}

func New(bucket *gocb.Bucket) *Repository {
	return &Repository{
		scope: bucket.Scope("media_service"),
		coll:  bucket.Scope("media_service").Collection("medias"),
	}
}

func (r *Repository) CreateMedia(ctx context.Context, media *models.Media) error {
	now := time.Now()
	media.CreatedAt = now
	media.UpdatedAt = now
	_, err := r.coll.Insert(media.ID, media, &gocb.InsertOptions{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("coll.Insert: %w", err)
	}
	return nil
}

func (r *Repository) DeleteMediaByID(ctx context.Context, id string) error {
	_, err := r.coll.Remove(id, &gocb.RemoveOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return customerrors.ErrRecordNotFound
		}
		return fmt.Errorf("coll.Remove: %w", err)
	}
	return nil
}

func (r *Repository) GetMediaByID(ctx context.Context, id string) (*models.Media, error) {
	result, err := r.coll.Get(id, &gocb.GetOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, fmt.Errorf("coll.Get: %w", err)
	}
	var media models.Media
	err = result.Content(&media)
	if err != nil {
		return nil, fmt.Errorf("result.Content: %w", err)
	}
	return &media, nil
}

const countQuery = "SELECT COUNT(*) as count FROM medias;"
const listQueryWithLimit = "SELECT * FROM medias ORDER BY id ASC LIMIT $limit OFFSET $offset;"
const listQuery = "SELECT * FROM medias ORDER BY id ASC OFFSET $offset;"

type listResult struct {
	Medias models.Media `json:"medias"`
}

func (r *Repository) ListMedias(ctx context.Context, limit int64, offset int64) (*models.MediaList, error) {
	var result struct {
		Count int64 `json:"count"`
	}
	countCursor, err := r.scope.Query(countQuery, &gocb.QueryOptions{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("scope.Query: %w", err)
	}
	defer countCursor.Close()
	err = countCursor.One(&result)
	if err != nil {
		return nil, fmt.Errorf("cursor.One: %w", err)
	}
	if result.Count == 0 || limit == 0 {
		return &models.MediaList{
			List:   make([]*models.Media, 0),
			Count:  result.Count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	cursor, err := r.scope.Query(lo.Ternary(limit < 0, listQuery, listQueryWithLimit), &gocb.QueryOptions{
		NamedParameters: map[string]any{
			"limit":  limit,
			"offset": offset,
		},
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("scope.Query: %w", err)
	}
	defer cursor.Close()
	var medias []*models.Media
	for cursor.Next() {
		var result listResult
		err = cursor.Row(&result)
		if err != nil {
			return nil, fmt.Errorf("cursor.Row: %w", err)
		}
		medias = append(medias, &result.Medias)
	}
	return &models.MediaList{
		List:   medias,
		Count:  result.Count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

const updateQuery = "UPDATE medias SET title = $title, updated_at = $updated_at WHERE id = $id RETURNING *;"

func (r *Repository) UpdateMediaByID(ctx context.Context, media *models.Media) error {
	media.UpdatedAt = time.Now()
	result, err := r.scope.Query(updateQuery, &gocb.QueryOptions{
		Context: ctx,
		NamedParameters: map[string]any{
			"title":      media.Title,
			"updated_at": time.Now(),
			"id":         media.ID,
		},
	})
	if err != nil {
		return fmt.Errorf("scope.Query: %w", err)
	}
	if !result.Next() {
		return customerrors.ErrRecordNotFound
	}
	return nil
}
