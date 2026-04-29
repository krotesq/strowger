package account

import (
	"context"
	"net/netip"
	"time"

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

func (repository *repository) createRefreshToken(ctx context.Context, accountID string, tokenHash []byte, expiresAt time.Time, userAgent string, ipAddress netip.Addr) (*refreshToken, error) {
	var refreshToken refreshToken
	const q = `INSERT INTO refresh_token (account_id, token_hash, expires_at, user_agent, ip_address) VALUES ($1, $2, $3, $4, $5) RETURNING *`
	err := pgxscan.Get(ctx, repository.pool, &refreshToken, q, accountID, tokenHash, expiresAt, userAgent, ipAddress)
	return &refreshToken, err
}

func (repository *repository) findRefreshTokenByHash(ctx context.Context, hash []byte) (*refreshToken, error) {
	var refreshToken refreshToken
	const q = `SELECT * FROM refresh_token WHERE token_hash = $1`
	err := pgxscan.Get(ctx, repository.pool, &refreshToken, q, hash)
	return &refreshToken, err
}

func (repository *repository) revokeRefreshTokenByHash(ctx context.Context, hash []byte) (*refreshToken, error) {
	var refreshToken refreshToken
	const q = `UPDATE refresh_token SET revoked_at = $2 WHERE token_hash = $1 RETURNING *`
	err := pgxscan.Get(ctx, repository.pool, &refreshToken, q, hash, time.Now())
	return &refreshToken, err
}
