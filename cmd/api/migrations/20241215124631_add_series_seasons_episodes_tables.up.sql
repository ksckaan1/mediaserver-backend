CREATE TABLE IF NOT EXISTS series (
  id VARCHAR(50) PRIMARY KEY,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  tmdb_id INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS seasons (
  id VARCHAR(50) PRIMARY KEY,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  series_id VARCHAR(50) NOT NULL,
  `order` INTEGER NOT NULL,
  FOREIGN KEY (series_id) REFERENCES series(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS episodes (
  id VARCHAR(50) PRIMARY KEY,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  season_id VARCHAR(50) NOT NULL,
  `order` INTEGER NOT NULL,
  media_id VARCHAR(50) NOT NULL DEFAULT '',
  FOREIGN KEY (season_id) REFERENCES seasons(id) ON DELETE CASCADE
);