-- TMDB Service
CREATE SCOPE IF NOT EXISTS `media_server`.tmdb_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.tmdb_service.infos;

-- Media Service
CREATE SCOPE IF NOT EXISTS `media_server`.media_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.media_service.medias;

-- Movie Service
CREATE SCOPE IF NOT EXISTS `media_server`.movie_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.movie_service.movies;

-- Series Service
CREATE SCOPE IF NOT EXISTS `media_server`.series_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.series_service.series;

-- Episode Service
CREATE SCOPE IF NOT EXISTS `media_server`.episode_service;
CREATE COLLECTION IF NOT EXISTS `media_server`.episode_service.episodes;
