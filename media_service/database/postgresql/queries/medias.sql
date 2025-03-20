-- name: CreateMedia :exec
INSERT INTO medias
(id, path, type, title, mime_type, size, created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6, NOW(), NOW());

-- name: GetMediaByID :one
SELECT *
FROM medias
WHERE id = $1;

-- name: ListMedias :many
SELECT *
FROM medias
LIMIT $1 OFFSET $2;

-- name: CountMedias :one
SELECT COUNT(*)
FROM medias;

-- name: UpdateMediaByID :one
UPDATE medias
SET title = $2, updated_at = NOW()
WHERE id = $1
RETURNING id;

-- name: DeleteMediaByID :one
DELETE FROM medias
WHERE id = $1
RETURNING id;