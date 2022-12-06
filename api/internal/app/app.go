// Package app configures and runs application.
package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"ways/config"
	"ways/internal/controller"
	"ways/pkg/httpserver"
	"ways/pkg/logger"
	"ways/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	lg := logger.NewLogger(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		lg.Fatal().Msgf("app - Run - postgres.New: %w", err)
	}
	defer pg.Close()

	// HTTP Server
	handler := gin.New()
	controller.NewRouter(handler, lg)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		lg.Info().Msg("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		lg.Error().Msgf("app - Run - httpServer.Notify: %w", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		lg.Error().Msgf("app - Run - httpServer.Shutdown: %w", err)
	}
}
