-- TMDB Service
CREATE SCOPE IF NOT EXISTS `media_server`.tmdb_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.tmdb_service.infos;
CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.tmdb_service.infos;
CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.tmdb_service.infos(id);

-- Media Service
CREATE SCOPE IF NOT EXISTS `media_server`.media_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.media_service.medias;
CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.media_service.medias;
CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.media_service.medias(id);

-- Movie Service
CREATE SCOPE IF NOT EXISTS `media_server`.movie_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.movie_service.movies;
CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.movie_service.movies;
CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.movie_service.movies(id);

-- Series Service
CREATE SCOPE IF NOT EXISTS `media_server`.series_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.series_service.series;
CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.series_service.series;
CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.series_service.series(id);

-- Season Service
CREATE SCOPE IF NOT EXISTS `media_server`.season_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.season_service.seasons;
CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.season_service.seasons;
CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.season_service.seasons(id);
CREATE INDEX IF NOT EXISTS idx_series_id ON `media_server`.season_service.seasons(series_id);

-- Episode Service
CREATE SCOPE IF NOT EXISTS `media_server`.episode_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.episode_service.episodes;
CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.episode_service.episodes;
CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.episode_service.episodes(id);
CREATE INDEX IF NOT EXISTS idx_season_id ON `media_server`.episode_service.episodes(season_id);

-- User Service
CREATE SCOPE IF NOT EXISTS `media_server`.user_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.user_service.users;
CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.user_service.users;
CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.user_service.users(id);
CREATE INDEX IF NOT EXISTS idx_username ON `media_server`.user_service.users(username);

-- Auth Service
CREATE SCOPE IF NOT EXISTS `media_server`.auth_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.auth_service.sessions;
CREATE PRIMARY INDEX IF NOT EXISTS ON `media_server`.auth_service.sessions;
CREATE INDEX IF NOT EXISTS idx_id ON `media_server`.auth_service.sessions(session_id);
CREATE INDEX IF NOT EXISTS idx_user_id ON `media_server`.auth_service.sessions(user_id);

