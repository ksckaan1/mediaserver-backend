-- name: CreateMovie :exec
INSERT INTO movies (
  id,
  created_at,
  updated_at,
  title,
  tmdb_id,
  description,
  media_id
) VALUES (
  ?,
  (datetime(CURRENT_TIMESTAMP, 'localtime')),
  (datetime(CURRENT_TIMESTAMP, 'localtime')),
  ?, ?, ?, ?
);

-- name: ListMovies :many
SELECT *
FROM movies
LIMIT ? OFFSET ?;

-- name: CountMovies :one
SELECT COUNT(*)
FROM movies;

-- name: GetMovieByID :one
SELECT *
FROM movies
WHERE id = ?;

-- name: UpdateMovieByID :one
UPDATE movies
SET updated_at = (datetime(CURRENT_TIMESTAMP, 'localtime')),
    title = ?,
    tmdb_id = ?,
    media_id = ?,
    description = ?
WHERE id = ?
RETURNING id;

-- name: DeleteMovieByID :one
DELETE FROM movies
WHERE id = ?
RETURNING id;