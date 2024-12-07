-- name: CreateMovie :exec
INSERT INTO movies (
  id,
  created_at,
  title,
  tmdb_id,
  description
) VALUES (
  ?, ?, ?, ?, ?
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

-- name: UpdateMovieByID :exec
UPDATE movies
SET title = ?,
    tmdb_id = ?,
    description = ?
WHERE id = ?;

-- name: DeleteMovieByID :exec
DELETE FROM movies
WHERE id = ?;