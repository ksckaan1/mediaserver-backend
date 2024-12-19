CREATE TABLE tmdb_infos (
    id VARCHAR(50) PRIMARY KEY,
    data jsonb NOT NULL DEFAULT '{}'
);