package main

import (
	"bff-service/internal/core/app/episode"
	"bff-service/internal/core/app/media"
	"bff-service/internal/core/app/movie"
	"bff-service/internal/core/app/season"
	"bff-service/internal/core/app/series"
	"shared/pb/episodepb"
	"shared/pb/mediapb"
	"shared/pb/moviepb"
	"shared/pb/seasonpb"
	"shared/pb/seriespb"

	"github.com/gofiber/fiber/v2"
)

func initMediaRoutes(mediaClient mediapb.MediaServiceClient) *fiber.App {
	app := fiber.New()
	app.Post("/", h(media.NewUploadMedia(mediaClient)))
	app.Get("/:media_id", h(media.NewGetMediaByID(mediaClient)))
	app.Get("/", h(media.NewListMedias(mediaClient)))
	app.Put("/:media_id", h(media.NewUpdateMediaByID(mediaClient)))
	app.Delete("/:media_id", h(media.NewDeleteMediaByID(mediaClient)))
	return app
}

func initMovieRoutes(movieClient moviepb.MovieServiceClient) *fiber.App {
	app := fiber.New()
	app.Post("/", h(movie.NewCreateMovie(movieClient)))
	app.Get("/:movie_id", h(movie.NewGetMovieByID(movieClient)))
	app.Get("/", h(movie.NewListMovies(movieClient)))
	app.Put("/:movie_id", h(movie.NewUpdateMovieByID(movieClient)))
	app.Delete("/:movie_id", h(movie.NewDeleteMovieByID(movieClient)))
	return app
}

func initSeriesRoutes(seriesClient seriespb.SeriesServiceClient) *fiber.App {
	app := fiber.New()
	app.Post("/", h(series.NewCreateSeries(seriesClient)))
	app.Get("/:series_id", h(series.NewGetSeriesByID(seriesClient)))
	app.Get("/", h(series.NewListSeries(seriesClient)))
	app.Put("/:series_id", h(series.NewUpdateSeriesByID(seriesClient)))
	app.Delete("/:series_id", h(series.NewDeleteSeriesByID(seriesClient)))
	return app
}

func initSeasonRoutes(seasonClient seasonpb.SeasonServiceClient) *fiber.App {
	app := fiber.New()
	app.Post("/", h(season.NewCreateSeason(seasonClient)))
	app.Get("/:season_id", h(season.NewGetSeasonByID(seasonClient)))
	app.Get("/", h(season.NewListSeasons(seasonClient)))
	app.Put("/:season_id", h(season.NewUpdateSeasonByID(seasonClient)))
	app.Delete("/:season_id", h(season.NewDeleteSeasonByID(seasonClient)))
	return app
}

func initEpisodeRoutes(episodeClient episodepb.EpisodeServiceClient) *fiber.App {
	app := fiber.New()
	app.Post("/", h(episode.NewCreateEpisode(episodeClient)))
	app.Get("/:episode_id", h(episode.NewGetEpisodeByID(episodeClient)))
	app.Get("/", h(episode.NewListEpisodes(episodeClient)))
	app.Put("/:episode_id", h(episode.NewUpdateEpisodeByID(episodeClient)))
	app.Delete("/:episode_id", h(episode.NewDeleteEpisodeByID(episodeClient)))
	return app
}
