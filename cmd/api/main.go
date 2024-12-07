package main

import (
	"context"
	"fmt"
	"mediaserver/config"

	"mediaserver/internal/pkg/logger"
)

func main() {
	ctx := context.Background()

	cfg := config.New()
	err := cfg.Load()
	if err != nil {
		panic(fmt.Errorf("cfg.Load: %w", err))
	}

	lg, err := logger.New()
	if err != nil {
		panic(fmt.Errorf("logger.New: %w", err))
	}

	s := NewServer(cfg, lg)

	err = s.Init(ctx)
	if err != nil {
		lg.Fatal(ctx, "error when init",
			"error", err,
		)
	}

	err = s.Run(ctx)
	if err != nil {
		lg.Fatal(ctx, "failed to run",
			"error", err,
		)
	}
}
