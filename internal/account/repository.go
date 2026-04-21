package account

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/georgysavva/scany/v2/pgxscan"
)

// handles all database interactions and returns models

type repository struct {
	pool *pgxpool.Pool
}

func newRepository(pool *pgxpool.Pool) *repository {
	return &repository{pool: pool}
}

func (r *repository) findByUsername(ctx context.Context, username string) (*account, error) {
	var account account
	const q = `SELECT uuid, username, password_hash, failed_login_attempts, active, created_at FROM account WHERE username = $1`
	err := pgxscan.Get(ctx, r.pool, &account, q, username)
	return &account, err
}