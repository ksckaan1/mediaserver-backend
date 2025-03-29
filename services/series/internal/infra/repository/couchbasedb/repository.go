package couchbasedb

import (
	"context"
	"errors"
	"fmt"
	"series_service/internal/core/customerrors"
	"series_service/internal/core/models"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/samber/lo"
)

type Repository struct {
	coll  *gocb.Collection
	scope *gocb.Scope
}

func New(bucket *gocb.Bucket) *Repository {
	return &Repository{
		coll:  bucket.Scope("series_service").Collection("series"),
		scope: bucket.Scope("series_service"),
	}
}

func (r *Repository) CreateSeries(ctx context.Context, series *models.Series) error {
	now := time.Now()
	series.CreatedAt = now
	series.UpdatedAt = now
	_, err := r.coll.Insert(series.ID, series, &gocb.InsertOptions{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("r.coll.Insert: %w", err)
	}
	return nil
}

func (r *Repository) DeleteSeriesByID(ctx context.Context, id string) error {
	_, err := r.coll.Remove(id, &gocb.RemoveOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return customerrors.ErrRecordNotFound
		}
		return fmt.Errorf("r.coll.Remove: %w", err)
	}
	return nil
}

func (r *Repository) GetSeriesByID(ctx context.Context, id string) (*models.Series, error) {
	result, err := r.coll.Get(id, &gocb.GetOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, customerrors.ErrRecordNotFound
		}
		return nil, fmt.Errorf("r.coll.Get: %w", err)
	}
	var series models.Series
	err = result.Content(&series)
	if err != nil {
		return nil, fmt.Errorf("result.Content: %w", err)
	}
	return &series, nil
}

const (
	countQuery         = `SELECT COUNT(*) AS count FROM series;`
	listQueryWithLimit = `SELECT * FROM series ORDER BY id ASC LIMIT $limit OFFSET $offset;`
	listQuery          = `SELECT * FROM series ORDER BY id ASC OFFSET $offset;`
)

type countResult struct {
	Count int64 `json:"count"`
}

type listResult struct {
	Series *models.Series `json:"series"`
}

func (r *Repository) ListSeries(ctx context.Context, limit int64, offset int64) (*models.SeriesList, error) {
	countCursor, err := r.scope.Query(countQuery, &gocb.QueryOptions{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("r.scope.Query: %w", err)
	}
	var countResult countResult
	err = countCursor.One(&countResult)
	if err != nil {
		return nil, fmt.Errorf("countCursor.One: %w", err)
	}
	if countResult.Count == 0 || limit == 0 {
		return &models.SeriesList{
			List:   []*models.Series{},
			Count:  countResult.Count,
			Limit:  limit,
			Offset: offset,
		}, nil
	}
	cursor, err := r.scope.Query(lo.Ternary(limit < 0, listQuery, listQueryWithLimit), &gocb.QueryOptions{
		Context: ctx,
		NamedParameters: map[string]any{
			"limit":  limit,
			"offset": offset,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("r.scope.Query: %w", err)
	}
	var seriesList []*models.Series
	for cursor.Next() {
		var seriesResult listResult
		err = cursor.Row(&seriesResult)
		if err != nil {
			return nil, fmt.Errorf("cursor.Row: %w", err)
		}
		seriesList = append(seriesList, seriesResult.Series)
	}
	return &models.SeriesList{
		List:   seriesList,
		Count:  countResult.Count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

func (r *Repository) UpdateSeriesByID(ctx context.Context, series *models.Series) error {
	_, err := r.coll.Replace(series.ID, map[string]any{
		"series":      series.Title,
		"description": series.Description,
		"releaseDate": series.TMDBID,
		"updatedAt":   time.Now(),
	}, &gocb.ReplaceOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return customerrors.ErrRecordNotFound
		}
		return fmt.Errorf("r.coll.Replace: %w", err)
	}
	return nil
}
