-- name: CreateUser :exec
INSERT INTO users (id, created_at, updated_at, email, display_name)
VALUES (
  ?, 
  (datetime(CURRENT_TIMESTAMP, 'localtime')),
  (datetime(CURRENT_TIMESTAMP, 'localtime')),
  ?, ?);

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = ?;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = ?;

-- name: ListUsers :many
SELECT *
FROM users
LIMIT ? OFFSET ?;

-- name: CountUsers :one
SELECT COUNT(*)
FROM users;

-- name: UpdateUserByID :one
UPDATE users
SET email = ?, display_name = ?, updated_at = (datetime(CURRENT_TIMESTAMP, 'localtime'))
WHERE id = ?
RETURNING id;

-- name: DeleteUserByID :one
DELETE
FROM users
WHERE id = ?
RETURNING id;