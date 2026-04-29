package source

import (
	"context"
	"database/sql"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

// handles all database interactions and returns models

type repository struct {
	pool *pgxpool.Pool
}

func newRepository(pool *pgxpool.Pool) *repository {
	return &repository{pool: pool}
}

func (repository *repository) findWithRtmpByID(ctx context.Context, id string) (*source, *rtmp, error) {
	var _source source
	const sourceQuery = `SELECT uuid, name, description, source_type_id, account_id, created_at FROM source WHERE uuid = $1`
	if err := pgxscan.Get(ctx, repository.pool, &_source, sourceQuery, id); err != nil {
		return &source{}, &rtmp{}, err
	}

	var _rtmp rtmp
	const rtmpQuery = `SELECT source_id, url, stream_key, created_at FROM source_rtmp WHERE source_id = $1`
	if err := pgxscan.Get(ctx, repository.pool, &_rtmp, rtmpQuery, _source.ID); err != nil {
		return &source{}, &rtmp{}, err
	}

	return &_source, &_rtmp, nil
}

func (repository *repository) createWithRtmp(ctx context.Context, name string, description string, accountID string, url string, stream_key string) (*source, *rtmp, error) {
	tx, err := repository.pool.Begin(ctx)
	if err != nil {
		return &source{}, &rtmp{}, err
	}
	defer tx.Rollback(ctx)

	var sourceType _type
	const sourceTypeQuery = `SELECT id, name FROM source_type WHERE name = 'rtmp'`
	if err := pgxscan.Get(ctx, tx, &sourceType, sourceTypeQuery); err != nil {
		return &source{}, &rtmp{}, err
	}

	const sourceQuery = `
			INSERT INTO source (name, description, source_type_id, account_id)
			VALUES (
				$1,
				$2,
				$3,
				$4
			)
			RETURNING uuid, name, description, source_type_id, account_id, created_at
		`
	var _source source
	if err := pgxscan.Get(ctx, tx, &_source, sourceQuery, name, description, sourceType.ID, accountID); err != nil {
		return &source{}, &rtmp{}, err
	}

	const rtmpQuery = `
			INSERT INTO source_rtmp (source_id, url, stream_key)
			VALUES ($1, $2, $3)
			RETURNING source_id, url, stream_key, created_at
		`
	var _rtmp rtmp
	if err := pgxscan.Get(ctx, tx, &_rtmp, rtmpQuery, _source.ID, url, stream_key); err != nil {
		return &source{}, &rtmp{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return &source{}, &rtmp{}, err
	}

	return &_source, &_rtmp, err
}

func (repository *repository) deleteByID(ctx context.Context, id string) error {
	const query = `
		DELETE FROM source
		WHERE uuid = $1
	`
	cmdTag, err := repository.pool.Exec(ctx, query, id)

	if cmdTag.RowsAffected() == 0 {
		return sql.ErrNoRows
	}

	return err
}
