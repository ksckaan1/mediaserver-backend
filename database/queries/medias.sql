-- name: CreateMedia :exec
INSERT INTO medias (id, path, type, storage_type, mime_type, size, created_at)
		VALUES(?, ?, ?, ?, ?, ?, (datetime (CURRENT_TIMESTAMP, 'localtime')));

-- name: GetMediaByID :one
SELECT *
FROM medias
WHERE id = ?;

-- name: ListMedias :many
SELECT *
FROM medias
LIMIT ? OFFSET ?;

-- name: CountMedias :one
SELECT COUNT(*)
FROM medias;

-- name: DeleteMediaByID :one
DELETE FROM medias
WHERE id = ?
RETURNING id;