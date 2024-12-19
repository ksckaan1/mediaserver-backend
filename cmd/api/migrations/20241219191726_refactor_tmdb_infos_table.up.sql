DROP TABLE IF EXISTS tmdb_infos;

CREATE TABLE IF NOT EXISTS tmdb_infos (
    id VARCHAR(50) PRIMARY KEY,
    data jsonb NOT NULL DEFAULT '{}'
);