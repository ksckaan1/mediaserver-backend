-- SERIES
ALTER TABLE series RENAME COLUMN tmdb_id TO tmdb_id_old;
ALTER TABLE series ADD COLUMN tmdb_id INTEGER NOT NULL DEFAULT '';
UPDATE series SET tmdb_id = CASE
  WHEN tmdb_id_old == '' THEN 0
  ELSE CAST(tmdb_id_old AS INTEGER)
END;
ALTER TABLE series DROP COLUMN tmdb_id_old;

-- MOVIES
ALTER TABLE movies RENAME COLUMN tmdb_id TO tmdb_id_old;
ALTER TABLE movies ADD COLUMN tmdb_id INTEGER NOT NULL DEFAULT '';
UPDATE movies SET tmdb_id = CASE
  WHEN tmdb_id_old == '' THEN 0
  ELSE CAST(tmdb_id_old AS INTEGER)
END;
ALTER TABLE movies DROP COLUMN tmdb_id_old;