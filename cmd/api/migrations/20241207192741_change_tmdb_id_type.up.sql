ALTER TABLE movies
	ADD COLUMN tmdb_id_old INTEGER NOT NULL DEFAULT 0;

UPDATE
	movies
SET
	tmdb_id_old = CASE WHEN tmdb_id IS NULL
		OR tmdb_id = '' THEN
		0
	ELSE
		CAST(tmdb_id AS INTEGER)
	END;

ALTER TABLE movies DROP COLUMN tmdb_id;

ALTER TABLE movies RENAME COLUMN tmdb_id_old TO tmdb_id;