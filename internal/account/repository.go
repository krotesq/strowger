package account

import (
	"context"

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

func (repository *repository) findByID(ctx context.Context, id string) (*account, error) {
	var account account
	const q = `SELECT * FROM account WHERE id = $1`
	err := pgxscan.Get(ctx, repository.pool, &account, q, id)
	return &account, err
}

func (repository *repository) incrementFailedLoginAttemptsByID(ctx context.Context, id string) (*account, error) {
	var account account
	const q = `UPDATE account SET failed_login_attempts = failed_login_attempts + 1 WHERE id = $1 RETURNING *`
	err := pgxscan.Get(ctx, repository.pool, &account, q, id)
	return &account, err
}

func (repository *repository) resetFailedLoginAttemptsByID(ctx context.Context, id string) (*account, error) {
	var account account
	const q = `UPDATE account SET failed_login_attempts = 0 WHERE id = $1 RETURNING *`
	err := pgxscan.Get(ctx, repository.pool, &account, q, id)
	return &account, err
}

func (repository *repository) findByUsername(ctx context.Context, username string) (*account, error) {
	var account account
	const q = `SELECT * FROM account WHERE username = $1`
	err := pgxscan.Get(ctx, repository.pool, &account, q, username)
	return &account, err
}

func (repository *repository) create(ctx context.Context, username, passwordHash string) (*account, error) {
	var account account
	const q = `INSERT INTO account (username, password_hash) VALUES ($1, $2) RETURNING *`
	err := pgxscan.Get(ctx, repository.pool, &account, q, username, passwordHash)
	return &account, err
}

func (repository *repository) deleteByID(ctx context.Context, id string) (*account, error) {
	var account account
	const q = `DELETE FROM account WHERE id = $1 RETURNING *`
	err := pgxscan.Get(ctx, repository.pool, &account, q, id)
	return &account, err
}

func (repository *repository) deactivateByID(ctx context.Context, id string) (*account, error) {
	var account account
	const q = `UPDATE account SET active = false WHERE id = $1 RETURNING *`
	err := pgxscan.Get(ctx, repository.pool, &account, q, id)
	return &account, err
}

func (repository *repository) activateByID(ctx context.Context, id string) (*account, error) {
	var account account
	const q = `UPDATE account SET active = true WHERE id = $1 RETURNING *`
	err := pgxscan.Get(ctx, repository.pool, &account, q, id)
	return &account, err
}
