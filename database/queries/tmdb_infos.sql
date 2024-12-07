-- name: GetTMDBInfo :one
SELECT * FROM tmdb_infos WHERE id = ?;


-- name: SetTMDBInfo :exec
INSERT INTO tmdb_infos (id, original_title, poster_path, backdrop_path, vote_average, vote_count, popularity, release_date)
		VALUES(?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8) ON CONFLICT (id)
		DO
		UPDATE
		SET
			original_title = ?2,
			poster_path = ?3,
			backdrop_path = ?4,
			vote_average = ?5,
			vote_count = ?6,
			popularity = ?7,
			release_date = ?8;