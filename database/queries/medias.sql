-- name: CreateMedia :exec
INSERT INTO medias (id, path, type, storage_type, mime_type, size, created_at)
		VALUES(?, ?, ?, ?, ?, ?, (datetime (CURRENT_TIMESTAMP, 'localtime')));

-- name: GetMediaByID :one
SELECT *
FROM medias
WHERE id = ?;

-- name: DeleteMediaByID :one
DELETE FROM medias
WHERE id = ?
RETURNING id;