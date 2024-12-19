------------
-- SERIES --
------------

-- name: CreateSeries :exec
INSERT INTO series (id, created_at, updated_at, title, description, tmdb_id)
		VALUES(?, (datetime (CURRENT_TIMESTAMP, 'localtime')), (datetime (CURRENT_TIMESTAMP, 'localtime')), ?, ?, ?);

-- name: ListSeries :many
SELECT *
FROM series
LIMIT ? OFFSET ?;

-- name: CountSeries :one
SELECT COUNT(*)
FROM series;

-- name: GetSeriesByID :one
SELECT *
FROM series
WHERE id = ?;

-- name: UpdateSeriesByID :one
UPDATE series
SET title = ?, description = ?, tmdb_id = ?, updated_at = (datetime (CURRENT_TIMESTAMP, 'localtime'))
WHERE id = ?
RETURNING id;

-- name: DeleteSeriesByID :one
DELETE FROM series
WHERE id = ?
RETURNING id;

-------------
-- SEASONS --
-------------

-- name: CreateSeason :exec
INSERT INTO seasons (id, created_at, updated_at, title, description, series_id, `order`)
    VALUES(?, (datetime (CURRENT_TIMESTAMP, 'localtime')), (datetime (CURRENT_TIMESTAMP, 'localtime')), ?, ?, ?, ?);

-- name: ListSeasonsBySeriesID :many
SELECT *
FROM seasons
WHERE series_id = ?
LIMIT ? OFFSET ?;

-- name: CountSeasonsBySeriesID :one
SELECT COUNT(*)
FROM seasons
WHERE series_id = ?;

-- name: GetSeasonByID :one
SELECT *
FROM seasons
WHERE id = ?;

-- name: UpdateSeasonByID :one
UPDATE seasons
SET title = ?, description = ?, `order` = ?, updated_at = (datetime (CURRENT_TIMESTAMP, 'localtime'))
WHERE id = ?
RETURNING id;

-- name: DeleteSeasonByID :one
DELETE FROM seasons
WHERE id = ?
RETURNING id;

--------------
-- EPISODES --
--------------

-- name: CreateEpisode :exec
INSERT INTO episodes (id, created_at, updated_at, title, description, season_id, `order`, media_id)
    VALUES(?, (datetime (CURRENT_TIMESTAMP, 'localtime')), (datetime (CURRENT_TIMESTAMP, 'localtime')), ?, ?, ?, ?, ?);

-- name: ListEpisodesBySeasonID :many
SELECT *
FROM episodes
WHERE season_id = ?
LIMIT ? OFFSET ?;

-- name: CountEpisodesBySeasonID :one
SELECT COUNT(*)
FROM episodes
WHERE season_id = ?;

-- name: GetEpisodeByID :one
SELECT *
FROM episodes
WHERE id = ?;

-- name: UpdateEpisodeByID :one
UPDATE episodes
SET title = ?, description = ?, `order` = ?, media_id = ?, updated_at = (datetime (CURRENT_TIMESTAMP, 'localtime'))
WHERE id = ?
RETURNING id; 

-- name: DeleteEpisodeByID :one
DELETE FROM episodes
WHERE id = ?
RETURNING id;
