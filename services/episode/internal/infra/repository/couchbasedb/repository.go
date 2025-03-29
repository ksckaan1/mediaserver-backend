package couchbasedb

import (
	"context"
	"episode_service/internal/core/app"
	"episode_service/internal/core/customerrors"
	"episode_service/internal/core/models"
	"errors"
	"fmt"
	"time"

	"github.com/couchbase/gocb/v2"
)

var _ app.Repository = (*Repository)(nil)

type Repository struct {
	coll  *gocb.Collection
	scope *gocb.Scope
}

func New(bucket *gocb.Bucket) *Repository {
	return &Repository{
		coll:  bucket.Scope("episode_service").Collection("episodes"),
		scope: bucket.Scope("episode_service"),
	}
}

func (r *Repository) CreateEpisode(ctx context.Context, episode *models.Episode) error {
	now := time.Now()
	episode.CreatedAt = now
	episode.UpdatedAt = now
	_, err := r.coll.Insert(episode.ID, episode, &gocb.InsertOptions{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("coll.Insert: %w", err)
	}
	return nil
}

func (r *Repository) DeleteAllEpisodesBySeasonID(ctx context.Context, seasonID string) error {
	query := `DELETE FROM episodes WHERE season_id = $1;`
	_, err := r.scope.Query(query, &gocb.QueryOptions{
		Context:              ctx,
		PositionalParameters: []any{seasonID},
	})
	if err != nil {
		return fmt.Errorf("scope.Query: %w", err)
	}
	return nil
}

func (r *Repository) DeleteEpisodeByID(ctx context.Context, episodeID string) error {
	_, err := r.coll.Remove(episodeID, &gocb.RemoveOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return customerrors.ErrEpisodeNotFound
		}
		return fmt.Errorf("coll.Remove: %w", err)
	}
	return nil
}

func (r *Repository) GetEpisodeByID(ctx context.Context, episodeID string) (*models.Episode, error) {
	result, err := r.coll.Get(episodeID, &gocb.GetOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, customerrors.ErrEpisodeNotFound
		}
		return nil, fmt.Errorf("coll.Get: %w", err)
	}
	var episode models.Episode
	err = result.Content(&episode)
	if err != nil {
		return nil, fmt.Errorf("result.Content: %w", err)
	}
	return &episode, nil
}

type listResult struct {
	Episodes models.Episode `json:"episodes"`
}

const listEpisodesQuery = "SELECT * FROM episodes WHERE season_id = $1 ORDER BY `order` ASC;"

func (r *Repository) ListEpisodesBySeasonID(ctx context.Context, seasonID string) ([]*models.Episode, error) {
	cursor, err := r.scope.Query(listEpisodesQuery, &gocb.QueryOptions{
		Context:              ctx,
		PositionalParameters: []any{seasonID},
	})
	if err != nil {
		return nil, fmt.Errorf("scope.Query: %w", err)
	}
	var episodes []*models.Episode
	for cursor.Next() {
		var result listResult
		err = cursor.Row(&result)
		if err != nil {
			return nil, fmt.Errorf("result.Row: %w", err)
		}
		episodes = append(episodes, &result.Episodes)
	}
	return episodes, nil
}

const updateQuery = `UPDATE episodes SET title = $title, media_id = $media_id, description = $description, updated_at = $updated_at WHERE id = $id RETURNING *;`

func (r *Repository) UpdateEpisodeByID(ctx context.Context, episode *models.Episode) error {
	result, err := r.scope.Query(updateQuery, &gocb.QueryOptions{
		Context: ctx,
		NamedParameters: map[string]any{
			"title":       episode.Title,
			"media_id":    episode.MediaID,
			"description": episode.Description,
			"updated_at":  time.Now(),
			"id":          episode.ID,
		},
	})
	if err != nil {
		return fmt.Errorf("scope.Query: %w", err)
	}
	if !result.Next() {
		return customerrors.ErrEpisodeNotFound
	}
	return nil
}

const updateOrder = "UPDATE episodes SET `order` = $order, updated_at = $updated_at WHERE id = $id RETURNING *;"

func (r *Repository) UpdateEpisodeOrder(ctx context.Context, episode *models.Episode) error {
	result, err := r.scope.Query(updateOrder, &gocb.QueryOptions{
		Context: ctx,
		NamedParameters: map[string]any{
			"order":      episode.Order,
			"updated_at": time.Now(),
			"id":         episode.ID,
		},
	})
	if err != nil {
		return fmt.Errorf("scope.Query: %w", err)
	}
	if !result.Next() {
		return customerrors.ErrEpisodeNotFound
	}
	return nil
}
