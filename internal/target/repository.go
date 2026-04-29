package target

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

func (repository *repository) findWithRtmpByID(ctx context.Context, id string) (*target, *rtmp, error) {
	var _target target
	const targetQuery = `SELECT uuid, name, description, target_type_id, account_id, created_at FROM target WHERE uuid = $1`
	if err := pgxscan.Get(ctx, repository.pool, &_target, targetQuery, id); err != nil {
		return &target{}, &rtmp{}, err
	}

	var _rtmp rtmp
	const rtmpQuery = `SELECT target_id, url, stream_key, created_at FROM target_rtmp WHERE target_id = $1`
	if err := pgxscan.Get(ctx, repository.pool, &_rtmp, rtmpQuery, _target.ID); err != nil {
		return &target{}, &rtmp{}, err
	}

	return &_target, &_rtmp, nil
}

func (repository *repository) createWithRtmp(ctx context.Context, name string, description string, accountID string, url string, stream_key string) (*target, *rtmp, error) {
	tx, err := repository.pool.Begin(ctx)
	if err != nil {
		return &target{}, &rtmp{}, err
	}
	defer tx.Rollback(ctx)

	var targetType _type
	const targetTypeQuery = `SELECT id, name FROM target_type WHERE name = 'rtmp'`
	if err := pgxscan.Get(ctx, tx, &targetType, targetTypeQuery); err != nil {
		return &target{}, &rtmp{}, err
	}

	const targetQuery = `
			INSERT INTO target (name, description, target_type_id, account_id)
			VALUES (
				$1,
				$2,
				$3,
				$4
			)
			RETURNING uuid, name, description, target_type_id, account_id, created_at
		`
	var _target target
	if err := pgxscan.Get(ctx, tx, &_target, targetQuery, name, description, targetType.ID, accountID); err != nil {
		return &target{}, &rtmp{}, err
	}

	const rtmpQuery = `
			INSERT INTO target_rtmp (target_id, url, stream_key)
			VALUES ($1, $2, $3)
			RETURNING target_id, url, stream_key, created_at
		`
	var _rtmp rtmp
	if err := pgxscan.Get(ctx, tx, &_rtmp, rtmpQuery, _target.ID, url, stream_key); err != nil {
		return &target{}, &rtmp{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return &target{}, &rtmp{}, err
	}

	return &_target, &_rtmp, err
}

func (repository *repository) delete(ctx context.Context, id string) error {
	const query = `
		DELETE FROM target
		WHERE uuid = $1
	`
	cmdTag, err := repository.pool.Exec(ctx, query, id)

	if cmdTag.RowsAffected() == 0 {
		return sql.ErrNoRows
	}

	return err
}
