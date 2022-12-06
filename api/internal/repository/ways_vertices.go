package repository

import (
	"ways/pkg/postgres"

	"github.com/rs/zerolog"
)

func NewWaysVerticesRepo(lg *zerolog.Logger, pg *postgres.Postgres) *WaysVerticesRepo {
	return &WaysVerticesRepo{
		logger: lg,
		pg:     pg,
	}
}

type WaysVerticesRepo struct {
	logger *zerolog.Logger
	pg     *postgres.Postgres
}
