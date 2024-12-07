ALTER TABLE movies
  ADD COLUMN tmdb_id_old TEXT NOT NULL DEFAULT '';

UPDATE
  movies
SET
  tmdb_id_old = CAST(tmdb_id AS TEXT)
WHERE
  tmdb_id IS NOT 0;

ALTER TABLE movies DROP COLUMN tmdb_id;

ALTER TABLE movies
	RENAME COLUMN tmdb_id_old TO tmdb_id;