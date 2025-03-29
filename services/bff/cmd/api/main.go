package main

import (
	"context"
	"shared/configer"
	"shared/logger"
	"shared/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()

	cfg := configer.New[service.ServiceConfig]()
	err := cfg.Load()
	if err != nil {
		panic(err)
	}

	lg, err := logger.New()
	if err != nil {
		panic(err)
	}

	app := fiber.New(fiber.Config{
		BodyLimit:             2 * 1024 * 1024 * 1024, // 2GB
		DisableStartupMessage: true,
	})

	v1 := app.Group("/api/v1")

	mediaApp, err := initMediaRoutes(cfg.Data)
	if err != nil {
		lg.Fatal(ctx, "initMediaRoutes",
			"error", err,
		)
	}
	v1.Mount("/media", mediaApp)

	movieApp, err := initMovieRoutes(cfg.Data)
	if err != nil {
		lg.Fatal(ctx, "initMovieRoutes",
			"error", err,
		)
	}
	v1.Mount("/movie", movieApp)

	seriesApp, err := initSeriesRoutes(cfg.Data)
	if err != nil {
		lg.Fatal(ctx, "initSeriesRoutes",
			"error", err,
		)
	}
	v1.Mount("/series", seriesApp)

	seasonApp, err := initSeasonRoutes(cfg.Data)
	if err != nil {
		lg.Fatal(ctx, "initSeasonRoutes",
			"error", err,
		)
	}
	v1.Mount("/season", seasonApp)

	episodeApp, err := initEpisodeRoutes(cfg.Data)
	if err != nil {
		lg.Fatal(ctx, "initEpisodeRoutes",
			"error", err,
		)
	}
	v1.Mount("/episode", episodeApp)

	app.Hooks().OnListen(func(data fiber.ListenData) error {
		lg.Info(ctx, "server is running",
			"host", data.Host,
			"port", data.Port,
			"tls", data.TLS,
		)
		return nil
	})

	err = app.Listen(cfg.Data.Addr)
	if err != nil {
		lg.Fatal(ctx, "Listen",
			"error", err,
		)
	}
}
