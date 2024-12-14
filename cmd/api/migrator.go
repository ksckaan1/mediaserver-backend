package main

import (
	"context"
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func (s *Server) migrate(ctx context.Context) error {
	migrationSource, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("iofs.New: %w", err)
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		migrationSource,
		fmt.Sprintf("sqlite3://%s", s.cfg.DBPath),
	)
	if err != nil {
		return fmt.Errorf("migrate.NewWithSourceInstance: %w", err)
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			s.logger.Info(ctx, "no migration needed")
			return nil
		}

		return fmt.Errorf("m.Up: %w", err)
	}

	s.logger.Info(ctx, "migrations applied")

	return nil
}
