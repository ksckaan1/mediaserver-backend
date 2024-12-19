-- name: GetTMDBInfo :one
SELECT * FROM tmdb_infos WHERE id = ?;


-- name: SetTMDBInfo :exec
INSERT INTO tmdb_infos (id, data)
		VALUES(?1, ?2) ON CONFLICT (id)
		DO
		UPDATE
		SET
			data = ?2;