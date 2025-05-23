package couchbasedb

import (
	"context"
	"errors"
	"fmt"
	"movie_service/internal/core/customerrors"
	"movie_service/internal/core/models"

	"github.com/couchbase/gocb/v2"
	"github.com/samber/lo"
)

type Repository struct {
	scope *gocb.Scope
	coll  *gocb.Collection
}

func New(bucket *gocb.Bucket) *Repository {
	return &Repository{
		scope: bucket.Scope("movie_service"),
		coll:  bucket.Scope("movie_service").Collection("movies"),
	}
}

func (r *Repository) CreateMovie(ctx context.Context, movie *models.Movie) error {
	_, err := r.coll.Insert(movie.ID, movie, &gocb.InsertOptions{
		Context: ctx,
	})
	if err != nil {
		return fmt.Errorf("coll.Insert: %w", err)
	}
	return nil
}

func (r *Repository) DeleteMovieByID(ctx context.Context, id string) error {
	_, err := r.coll.Remove(id, &gocb.RemoveOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return customerrors.ErrMovieNotFound
		}
		return fmt.Errorf("coll.Remove: %w", err)
	}
	return nil
}

func (r *Repository) GetMovieByID(ctx context.Context, id string) (*models.Movie, error) {
	result, err := r.coll.Get(id, &gocb.GetOptions{
		Context: ctx,
	})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, customerrors.ErrMovieNotFound
		}
		return nil, fmt.Errorf("coll.Get: %w", err)
	}
	var movie models.Movie
	err = result.Content(&movie)
	if err != nil {
		return nil, fmt.Errorf("result.Content: %w", err)
	}
	return &movie, nil
}

const (
	countQuery         = `SELECT COUNT(*) AS count FROM movies;`
	listQueryWithLimit = `SELECT * FROM movies WHERE id != "" ORDER BY id ASC LIMIT $limit OFFSET $offset;`
	listQuery          = `SELECT * FROM movies WHERE id != "" ORDER BY id ASC OFFSET $offset;`
)

type countResult struct {
	Count int64 `json:"count"`
}

type listResult struct {
	Movies models.Movie `json:"movies"`
}

func (r *Repository) ListMovies(ctx context.Context, limit int64, offset int64) (*models.MovieList, error) {
	var countResult countResult
	countCursor, err := r.scope.Query(countQuery, &gocb.QueryOptions{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("scope.Query: %w", err)
	}
	err = countCursor.One(&countResult)
	if err != nil {
		return nil, fmt.Errorf("countCursor.One: %w", err)
	}
	if countResult.Count == 0 || limit == 0 {
		return &models.MovieList{
			List:   []*models.Movie{},
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
		return nil, fmt.Errorf("scope.Query: %w", err)
	}
	var movies []*models.Movie
	for cursor.Next() {
		var result listResult
		err = cursor.Row(&result)
		if err != nil {
			return nil, fmt.Errorf("cursor.Row: %w", err)
		}
		movies = append(movies, &result.Movies)
	}
	return &models.MovieList{
		List:   movies,
		Count:  countResult.Count,
		Limit:  limit,
		Offset: offset,
	}, nil
}

const updateQuery = `UPDATE movies SET title = $title, description = $description, media_id = $media_id, tmdb_id = $tmdb_id, tags = $tags, updated_at = $updated_at WHERE id = $id RETURNING *;`

func (r *Repository) UpdateMovieByID(ctx context.Context, movie *models.Movie) error {
	result, err := r.scope.Query(updateQuery, &gocb.QueryOptions{
		Context: ctx,
		NamedParameters: map[string]any{
			"title":       movie.Title,
			"description": movie.Description,
			"media_id":    movie.MediaID,
			"tmdb_id":     movie.TMDBID,
			"tags":        movie.Tags,
			"updated_at":  movie.UpdatedAt,
			"id":          movie.ID,
		},
	})
	if err != nil {
		return fmt.Errorf("scope.Query: %w", err)
	}
	if !result.Next() {
		return customerrors.ErrMovieNotFound
	}
	return nil
}
