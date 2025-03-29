package main

import (
	"bff-service/internal/core/app/media"
	"bff-service/internal/core/app/movie"
	"bff-service/internal/core/app/series"
	"fmt"
	"shared/service"

	"github.com/gofiber/fiber/v2"
)

func initMediaRoutes(cfg *service.ServiceConfig) (*fiber.App, error) {
	mediaClient, err := NewMediaServiceClient(cfg.MediaServiceAddr)
	if err != nil {
		return nil, fmt.Errorf("NewMediaServiceClient: %w", err)
	}
	app := fiber.New()
	app.Post("/", h(media.NewUploadMedia(mediaClient)))
	app.Get("/:media_id", h(media.NewGetMediaByID(mediaClient)))
	app.Get("/", h(media.NewListMedias(mediaClient)))
	app.Put("/:media_id", h(media.NewUpdateMediaByID(mediaClient)))
	app.Delete("/:media_id", h(media.NewDeleteMediaByID(mediaClient)))
	return app, nil
}

func initMovieRoutes(cfg *service.ServiceConfig) (*fiber.App, error) {
	movieClient, err := NewMovieServiceClient(cfg.MovieServiceAddr)
	if err != nil {
		return nil, fmt.Errorf("NewMovieServiceClient: %w", err)
	}
	app := fiber.New()
	app.Post("/", h(movie.NewCreateMovie(movieClient)))
	app.Get("/:movie_id", h(movie.NewGetMovieByID(movieClient)))
	app.Get("/", h(movie.NewListMovies(movieClient)))
	app.Put("/:movie_id", h(movie.NewUpdateMovieByID(movieClient)))
	app.Delete("/:movie_id", h(movie.NewDeleteMovieByID(movieClient)))
	return app, nil
}

func initSeriesRoutes(cfg *service.ServiceConfig) (*fiber.App, error) {
	seriesClient, err := NewSeriesServiceClient(cfg.SeriesServiceAddr)
	if err != nil {
		return nil, fmt.Errorf("NewSeriesServiceClient: %w", err)
	}
	app := fiber.New()
	app.Post("/", h(series.NewCreateSeries(seriesClient)))
	app.Get("/:series_id", h(series.NewGetSeriesByID(seriesClient)))
	app.Get("/", h(series.NewListSeries(seriesClient)))
	app.Put("/:series_id", h(series.NewUpdateSeriesByID(seriesClient)))
	app.Delete("/:series_id", h(series.NewDeleteSeriesByID(seriesClient)))
	return app, nil
}
