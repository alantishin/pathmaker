package repository

import (
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
	// sql, _, err := w.pg.Builder.
	// 	Select("source, destination, original, translation").
	// 	From("history").
	// 	ToSql()
	// if err != nil {
	// 	return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Builder: %w", err)
	// }

	// rows, err := w.pg.Pool.Query(ctx, sql)
	// if err != nil {
	// 	return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Pool.Query: %w", err)
	// }

	// defer rows.Close()
}

// SELECT
// 	d.*,
// 	w.the_geom

// FROM pgr_dijkstra(
// 	'
// SELECT gid as id, source, target, cost as cost, cost as reverse_cost
// FROM public.ways
// 	',
// 	167884,
// 	167911,
// 	true
// ) as d

// INNER JOIN public.ways as w
// ON w.gid = d.edge
