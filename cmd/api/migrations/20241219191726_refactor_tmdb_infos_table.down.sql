DROP TABLE IF EXISTS tmdb_infos;

CREATE TABLE tmdb_infos (
    id VARCHAR(50) PRIMARY KEY,
    title TEXT NOT NULL,
    original_title TEXT NOT NULL,
    poster_path TEXT NOT NULL,
    backdrop_path TEXT NOT NULL,
    vote_average REAL NOT NULL,
    vote_count INTEGER NOT NULL,
    popularity REAL NOT NULL,
    release_date TEXT NOT NULL
);