package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"shared/configer"
	"shared/logger"
	"shared/ports"
	"syscall"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
)

type RESTService[CFG any] struct {
	Cfg        *CFG
	ServiceCfg *ServiceConfig
	Logger     ports.Logger

	ServiceClients *ServiceClients

	Router *fiber.App

	// protected fields
	initializer func(context.Context, *RESTService[CFG]) error
}

func NewREST[CFG any](initializer func(context.Context, *RESTService[CFG]) error) *RESTService[CFG] {
	return &RESTService[CFG]{
		Logger:         logger.New(),
		initializer:    initializer,
		ServiceClients: &ServiceClients{},
	}
}

func (s *RESTService[CFG]) Run(ctx context.Context) error {
	serviceCfg := configer.New[ServiceConfig]()
	err := serviceCfg.Load()
	if err != nil {
		return fmt.Errorf("configer[ServiceConfig].Load: %w", err)
	}
	s.ServiceCfg = serviceCfg.Data

	cfg := configer.New[CFG]()
	err = cfg.Load()
	if err != nil {
		return fmt.Errorf("configer[%T].Load: %w", cfg, err)
	}
	s.Cfg = cfg.Data

	err = s.startTracing()
	if err != nil {
		return fmt.Errorf("startTracing: %w", err)
	}

	err = s.ServiceClients.initServiceClients(s.ServiceCfg)
	if err != nil {
		return fmt.Errorf("initServiceClients: %w", err)
	}

	s.Router = fiber.New(
		fiber.Config{
			BodyLimit:             2 * 1024 * 1024 * 1024, // 2GB
			DisableStartupMessage: true,
		},
	)

	if s.ServiceCfg.TracerEndpoint != "" {
		s.Router.Use(otelfiber.Middleware())
	}

	s.Logger.Info(ctx, "service initializing")

	err = s.initializer(ctx, s)
	if err != nil {
		return fmt.Errorf("initializer: %w", err)
	}
	s.Logger.Info(ctx, "service initialized")

	s.Router.Hooks().OnListen(func(data fiber.ListenData) error {
		s.Logger.Info(ctx, "server is running",
			"host", data.Host,
			"port", data.Port,
			"tls", data.TLS,
		)
		return nil
	})

	go s.handleGracefulShutdown()

	err = s.Router.Listen(s.ServiceCfg.Addr)
	if err != nil {
		return fmt.Errorf("Router.Listen: %w", err)
	}

	s.Logger.Info(ctx, "service stopped")

	return nil
}

func (s *RESTService[CFG]) handleGracefulShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	s.Logger.Info(context.Background(), "shutting down gracefully")
	err := s.Router.Shutdown()
	if err != nil {
		s.Logger.Error(context.Background(), "server shutdown failed", "error", err)
	}
	s.Logger.Info(context.Background(), "server gracefully stopped")
}
