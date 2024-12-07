CREATE TABLE movies (
    id VARCHAR(50) PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME,
    title TEXT NOT NULL,
    tmdb_id TEXT NOT NULL,
    description TEXT NOT NULL
);