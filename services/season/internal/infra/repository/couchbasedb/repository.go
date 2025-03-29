package couchbasedb

import (
	"context"
	"errors"
	"fmt"
	"season_service/internal/core/customerrors"
	"season_service/internal/core/models"
	"time"

	"github.com/couchbase/gocb/v2"
)

type Repository struct {
	coll  *gocb.Collection
	scope *gocb.Scope
}

func New(bucket *gocb.Bucket) *Repository {
	return &Repository{
		coll:  bucket.Scope("season_service").Collection("seasons"),
		scope: bucket.Scope("season_service"),
	}
}

func (r *Repository) CreateSeason(ctx context.Context, season *models.Season) error {
	now := time.Now()
	season.CreatedAt = now
	season.UpdatedAt = now
	_, err := r.coll.Insert(season.ID, season, &gocb.InsertOptions{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("coll.Insert: %w", err)
	}
	return nil
}

func (r *Repository) DeleteSeasonByID(ctx context.Context, id string) error {
	_, err := r.coll.Remove(id, &gocb.RemoveOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return customerrors.ErrSeasonNotFound
		}
		return fmt.Errorf("coll.Remove: %w", err)
	}
	return nil
}

func (r *Repository) GetSeasonByID(ctx context.Context, id string) (*models.Season, error) {
	result, err := r.coll.Get(id, &gocb.GetOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, customerrors.ErrSeasonNotFound
		}
		return nil, fmt.Errorf("coll.Get: %w", err)
	}
	var season models.Season
	err = result.Content(&season)
	if err != nil {
		return nil, fmt.Errorf("result.Content: %w", err)
	}
	return &season, nil
}

const (
	listSeasonsQuery = `SELECT * FROM seasons WHERE series_id = $1 ORDER BY 'order' ASC;`
)

type listResult struct {
	Seasons models.Season `json:"seasons"`
}

func (r *Repository) ListSeasonsBySeriesID(ctx context.Context, seriesID string) ([]*models.Season, error) {
	result, err := r.scope.Query(listSeasonsQuery, &gocb.QueryOptions{
		PositionalParameters: []any{seriesID},
	})
	if err != nil {
		return nil, fmt.Errorf("scope.Query: %w", err)
	}
	var seasons []*models.Season
	for result.Next() {
		var row listResult
		err := result.Row(&row)
		if err != nil {
			return nil, fmt.Errorf("result.Row: %w", err)
		}
		seasons = append(seasons, &row.Seasons)
	}
	return seasons, nil
}

const updateQuery = `UPDATE seasons SET title = $title, description = $description, updated_at = $updated_at WHERE id = $id RETURNING *;`

func (r *Repository) UpdateSeasonByID(ctx context.Context, season *models.Season) error {
	result, err := r.scope.Query(updateQuery, &gocb.QueryOptions{
		Context: ctx,
		NamedParameters: map[string]any{
			"title":       season.Title,
			"description": season.Description,
			"updated_at":  time.Now(),
			"id":          season.ID,
		},
	})
	if err != nil {
		return fmt.Errorf("scope.Query: %w", err)
	}
	if !result.Next() {
		return customerrors.ErrSeasonNotFound
	}
	return nil
}

const updateOrderQuery = `UPDATE seasons SET order = $order, updated_at = $updated_at WHERE id = $id RETURNING *;`

func (r *Repository) UpdateSeasonOrderByID(ctx context.Context, season *models.Season) error {
	result, err := r.scope.Query(updateOrderQuery, &gocb.QueryOptions{
		Context: ctx,
		NamedParameters: map[string]any{
			"order":      season.Order,
			"updated_at": time.Now(),
			"id":         season.ID,
		},
	})
	if err != nil {
		return fmt.Errorf("scope.Query: %w", err)
	}
	if !result.Next() {
		return customerrors.ErrSeasonNotFound
	}
	return nil
}
