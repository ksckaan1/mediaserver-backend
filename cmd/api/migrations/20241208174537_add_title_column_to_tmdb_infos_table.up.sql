ALTER TABLE tmdb_infos ADD COLUMN title TEXT NOT NULL DEFAULT '';

UPDATE
  tmdb_infos
SET
  title = original_title;