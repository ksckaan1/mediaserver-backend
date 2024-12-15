// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: series.sql

package sqlcgen

import (
	"context"
)

const countEpisodesBySeasonID = `-- name: CountEpisodesBySeasonID :one
SELECT COUNT(*)
FROM episodes
WHERE season_id = ?
`

func (q *Queries) CountEpisodesBySeasonID(ctx context.Context, seasonID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countEpisodesBySeasonID, seasonID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countSeasonsBySeriesID = `-- name: CountSeasonsBySeriesID :one
SELECT COUNT(*)
FROM seasons
WHERE series_id = ?
`

func (q *Queries) CountSeasonsBySeriesID(ctx context.Context, seriesID string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countSeasonsBySeriesID, seriesID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countSeries = `-- name: CountSeries :one
SELECT COUNT(*)
FROM series
`

func (q *Queries) CountSeries(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countSeries)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createEpisode = `-- name: CreateEpisode :exec

INSERT INTO episodes (id, created_at, updated_at, name, description, season_id, ` + "`" + `order` + "`" + `, media_id)
    VALUES(?, (datetime (CURRENT_TIMESTAMP, 'localtime')), (datetime (CURRENT_TIMESTAMP, 'localtime')), ?, ?, ?, ?, ?)
`

type CreateEpisodeParams struct {
	ID          string
	Name        string
	Description string
	SeasonID    string
	Order       int64
	MediaID     string
}

// ------------
// EPISODES --
// ------------
func (q *Queries) CreateEpisode(ctx context.Context, arg CreateEpisodeParams) error {
	_, err := q.db.ExecContext(ctx, createEpisode,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.SeasonID,
		arg.Order,
		arg.MediaID,
	)
	return err
}

const createSeason = `-- name: CreateSeason :exec

INSERT INTO seasons (id, created_at, updated_at, name, description, series_id, ` + "`" + `order` + "`" + `)
    VALUES(?, (datetime (CURRENT_TIMESTAMP, 'localtime')), (datetime (CURRENT_TIMESTAMP, 'localtime')), ?, ?, ?, ?)
`

type CreateSeasonParams struct {
	ID          string
	Name        string
	Description string
	SeriesID    string
	Order       int64
}

// -----------
// SEASONS --
// -----------
func (q *Queries) CreateSeason(ctx context.Context, arg CreateSeasonParams) error {
	_, err := q.db.ExecContext(ctx, createSeason,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.SeriesID,
		arg.Order,
	)
	return err
}

const createSeries = `-- name: CreateSeries :exec

INSERT INTO series (id, created_at, updated_at, name, description, tmdb_id)
		VALUES(?, (datetime (CURRENT_TIMESTAMP, 'localtime')), (datetime (CURRENT_TIMESTAMP, 'localtime')), ?, ?, ?)
`

type CreateSeriesParams struct {
	ID          string
	Name        string
	Description string
	TmdbID      int64
}

// ----------
// SERIES --
// ----------
func (q *Queries) CreateSeries(ctx context.Context, arg CreateSeriesParams) error {
	_, err := q.db.ExecContext(ctx, createSeries,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.TmdbID,
	)
	return err
}

const deleteEpisodeByID = `-- name: DeleteEpisodeByID :one
DELETE FROM episodes
WHERE id = ?
RETURNING id
`

func (q *Queries) DeleteEpisodeByID(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, deleteEpisodeByID, id)
	err := row.Scan(&id)
	return id, err
}

const deleteSeasonByID = `-- name: DeleteSeasonByID :one
DELETE FROM seasons
WHERE id = ?
RETURNING id
`

func (q *Queries) DeleteSeasonByID(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, deleteSeasonByID, id)
	err := row.Scan(&id)
	return id, err
}

const deleteSeriesByID = `-- name: DeleteSeriesByID :one
DELETE FROM series
WHERE id = ?
RETURNING id
`

func (q *Queries) DeleteSeriesByID(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, deleteSeriesByID, id)
	err := row.Scan(&id)
	return id, err
}

const getEpisodeByID = `-- name: GetEpisodeByID :one
SELECT id, created_at, updated_at, name, description, season_id, ` + "`" + `order` + "`" + `, media_id
FROM episodes
WHERE id = ?
`

func (q *Queries) GetEpisodeByID(ctx context.Context, id string) (Episode, error) {
	row := q.db.QueryRowContext(ctx, getEpisodeByID, id)
	var i Episode
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.SeasonID,
		&i.Order,
		&i.MediaID,
	)
	return i, err
}

const getSeasonByID = `-- name: GetSeasonByID :one
SELECT id, created_at, updated_at, name, description, series_id, ` + "`" + `order` + "`" + `
FROM seasons
WHERE id = ?
`

func (q *Queries) GetSeasonByID(ctx context.Context, id string) (Season, error) {
	row := q.db.QueryRowContext(ctx, getSeasonByID, id)
	var i Season
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.SeriesID,
		&i.Order,
	)
	return i, err
}

const getSeriesByID = `-- name: GetSeriesByID :one
SELECT id, created_at, updated_at, name, description, tmdb_id
FROM series
WHERE id = ?
`

func (q *Queries) GetSeriesByID(ctx context.Context, id string) (Series, error) {
	row := q.db.QueryRowContext(ctx, getSeriesByID, id)
	var i Series
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.TmdbID,
	)
	return i, err
}

const listEpisodesBySeasonID = `-- name: ListEpisodesBySeasonID :many
SELECT id, created_at, updated_at, name, description, season_id, ` + "`" + `order` + "`" + `, media_id
FROM episodes
WHERE season_id = ?
LIMIT ? OFFSET ?
`

type ListEpisodesBySeasonIDParams struct {
	SeasonID string
	Limit    int64
	Offset   int64
}

func (q *Queries) ListEpisodesBySeasonID(ctx context.Context, arg ListEpisodesBySeasonIDParams) ([]Episode, error) {
	rows, err := q.db.QueryContext(ctx, listEpisodesBySeasonID, arg.SeasonID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Episode
	for rows.Next() {
		var i Episode
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
			&i.SeasonID,
			&i.Order,
			&i.MediaID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSeasonsBySeriesID = `-- name: ListSeasonsBySeriesID :many
SELECT id, created_at, updated_at, name, description, series_id, ` + "`" + `order` + "`" + `
FROM seasons
WHERE series_id = ?
LIMIT ? OFFSET ?
`

type ListSeasonsBySeriesIDParams struct {
	SeriesID string
	Limit    int64
	Offset   int64
}

func (q *Queries) ListSeasonsBySeriesID(ctx context.Context, arg ListSeasonsBySeriesIDParams) ([]Season, error) {
	rows, err := q.db.QueryContext(ctx, listSeasonsBySeriesID, arg.SeriesID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Season
	for rows.Next() {
		var i Season
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
			&i.SeriesID,
			&i.Order,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSeries = `-- name: ListSeries :many
SELECT id, created_at, updated_at, name, description, tmdb_id
FROM series
LIMIT ? OFFSET ?
`

type ListSeriesParams struct {
	Limit  int64
	Offset int64
}

func (q *Queries) ListSeries(ctx context.Context, arg ListSeriesParams) ([]Series, error) {
	rows, err := q.db.QueryContext(ctx, listSeries, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Series
	for rows.Next() {
		var i Series
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
			&i.TmdbID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEpisodeByID = `-- name: UpdateEpisodeByID :one
UPDATE episodes
SET name = ?, description = ?, ` + "`" + `order` + "`" + ` = ?, media_id = ?, updated_at = (datetime (CURRENT_TIMESTAMP, 'localtime'))
WHERE id = ?
RETURNING id
`

type UpdateEpisodeByIDParams struct {
	Name        string
	Description string
	Order       int64
	MediaID     string
	ID          string
}

func (q *Queries) UpdateEpisodeByID(ctx context.Context, arg UpdateEpisodeByIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, updateEpisodeByID,
		arg.Name,
		arg.Description,
		arg.Order,
		arg.MediaID,
		arg.ID,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const updateSeasonByID = `-- name: UpdateSeasonByID :one
UPDATE seasons
SET name = ?, description = ?, ` + "`" + `order` + "`" + ` = ?, updated_at = (datetime (CURRENT_TIMESTAMP, 'localtime'))
WHERE id = ?
RETURNING id
`

type UpdateSeasonByIDParams struct {
	Name        string
	Description string
	Order       int64
	ID          string
}

func (q *Queries) UpdateSeasonByID(ctx context.Context, arg UpdateSeasonByIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, updateSeasonByID,
		arg.Name,
		arg.Description,
		arg.Order,
		arg.ID,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const updateSeriesByID = `-- name: UpdateSeriesByID :one
UPDATE series
SET name = ?, description = ?, tmdb_id = ?, updated_at = (datetime (CURRENT_TIMESTAMP, 'localtime'))
WHERE id = ?
RETURNING id
`

type UpdateSeriesByIDParams struct {
	Name        string
	Description string
	TmdbID      int64
	ID          string
}

func (q *Queries) UpdateSeriesByID(ctx context.Context, arg UpdateSeriesByIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, updateSeriesByID,
		arg.Name,
		arg.Description,
		arg.TmdbID,
		arg.ID,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}
