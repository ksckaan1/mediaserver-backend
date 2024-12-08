-- name: GetTMDBInfo :one
SELECT * FROM tmdb_infos WHERE id = ?;


-- name: SetTMDBInfo :exec
INSERT INTO tmdb_infos (id,title, original_title, poster_path, backdrop_path, vote_average, vote_count, popularity, release_date)
		VALUES(?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8, ?9) ON CONFLICT (id)
		DO
		UPDATE
		SET
			title = ?2,
			original_title = ?3,
			poster_path = ?4,
			backdrop_path = ?5,
			vote_average = ?6,
			vote_count = ?7,
			popularity = ?8,
			release_date = ?9;