package repository

import (
	"fmt"
	"ways/config"
	"ways/internal/entity"
	"ways/pkg/postgres"
)

func NewWaysRepo(cfg config.Config, pg *postgres.Postgres) *WaysRepo {
	return &WaysRepo{
		pg: pg,
	}
}

type WaysRepo struct {
	pg *postgres.Postgres
}

func (w *WaysRepo) QueryNearestVerticy(point entity.GeoPoint) {
	sql, _, err := w.pg.Builder.
		Select("source, destination, original, translation").
		From("history").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Builder: %w", err)
	}

	rows, err := w.pg.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Pool.Query: %w", err)
	}

	defer rows.Close()
}
