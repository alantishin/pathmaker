package di

import (
	"ways/config"
	"ways/internal/repository"
	"ways/internal/service"
	"ways/pkg/logger"
	"ways/pkg/postgres"

	"github.com/rs/zerolog"
)

func NewDI(cfg *config.Config) *DI {
	lg := logger.NewLogger(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		lg.Fatal().Msgf("app - Run - postgres.New: %w", err)
	}

	verticesRepo := repository.NewWaysVerticesRepo(lg, pg)

	return &DI{
		Logger:      lg,
		DB:          pg,
		WaysService: *service.NewWaysService(lg, verticesRepo),
	}
}

type DI struct {
	Logger      *zerolog.Logger
	DB          *postgres.Postgres
	WaysService service.WaysService
}
