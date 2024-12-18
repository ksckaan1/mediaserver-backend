package main

import (
	"context"
	"database/sql"
	"fmt"
	"mediaserver/config"
	mediaapp "mediaserver/internal/domain/core/application/media"
	mediaservice "mediaserver/internal/domain/core/service/media"

	movieapp "mediaserver/internal/domain/core/application/movie"
	"mediaserver/internal/domain/core/service/localstorage"
	movieservice "mediaserver/internal/domain/core/service/movie"
	"mediaserver/internal/infrastructure/repository"
	"mediaserver/internal/infrastructure/tmdbclient"
	"mediaserver/internal/pkg/idgen"
	"mediaserver/internal/port"

	_ "modernc.org/sqlite"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg      *config.Config
	logger   port.Logger
	router   *fiber.App
	movieApp *movieapp.Movie
	mediaApp *mediaapp.Media
}

func NewServer(cfg *config.Config, logger port.Logger) *Server {
	router := fiber.New(fiber.Config{
		BodyLimit:             1 << 30, // 1GB
		DisableStartupMessage: true,
	})

	return &Server{
		cfg:    cfg,
		logger: logger,
		router: router,
	}
}

func (s *Server) Init(ctx context.Context) error {
	idGen, err := idgen.New()
	if err != nil {
		return fmt.Errorf("idgen.New: %w", err)
	}

	db, err := sql.Open("sqlite3", s.cfg.DBPath)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	if s.cfg.DBAutoMigrate {
		err = s.migrate(ctx)
		if err != nil {
			return fmt.Errorf("migrate: %w", err)
		}
	}

	repo, err := repository.New(db)
	if err != nil {
		return fmt.Errorf("movierepository.New: %w", err)
	}

	tmdbClient, err := tmdbclient.New(s.cfg.TMDBAPIKey)
	if err != nil {
		return fmt.Errorf("tmdbclient.New: %w", err)
	}

	movieService, err := movieservice.New(repo, tmdbClient, idGen, s.logger)
	if err != nil {
		return fmt.Errorf("movieservice.New: %w", err)
	}

	s.movieApp, err = movieapp.New(movieService)
	if err != nil {
		return fmt.Errorf("movieapp.New: %w", err)
	}

	lss, err := localstorage.New(repo, s.cfg, idGen, s.logger)
	if err != nil {
		return fmt.Errorf("localstorage.New: %w", err)
	}

	mediaService, err := mediaservice.New(repo, idGen, s.logger)
	if err != nil {
		return fmt.Errorf("mediaservice.New: %w", err)
	}

	s.mediaApp, err = mediaapp.New(lss, mediaService, s.logger)
	if err != nil {
		return fmt.Errorf("mediaapp.New: %w", err)
	}

	s.linkRoutes()

	return nil
}

func (s *Server) Run(ctx context.Context) error {
	s.logger.Info(ctx, "api starting",
		"port", s.cfg.Port,
	)
	err := s.router.Listen(fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return fmt.Errorf("router.Listen: %w", err)
	}

	return nil
}
