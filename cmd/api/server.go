package main

import (
	"context"
	"database/sql"
	"fmt"
	"mediaserver/config"
	movieapp "mediaserver/internal/domain/core/application/movie"
	movieservice "mediaserver/internal/domain/core/service/movie"
	"mediaserver/internal/infrastructure/movierepository"
	"mediaserver/internal/pkg/idgen"
	"mediaserver/internal/port"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg      *config.Config
	logger   port.Logger
	router   *fiber.App
	movieApp *movieapp.Movie
}

func NewServer(cfg *config.Config, logger port.Logger) *Server {
	router := fiber.New(fiber.Config{
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

	movieRepository, err := movierepository.New(db)
	if err != nil {
		return fmt.Errorf("movierepository.New: %w", err)
	}

	movieService, err := movieservice.New(movieRepository, idGen, s.logger)
	if err != nil {
		return fmt.Errorf("movieservice.New: %w", err)
	}

	s.movieApp, err = movieapp.New(movieService)
	if err != nil {
		return fmt.Errorf("movieapp.New: %w", err)
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
