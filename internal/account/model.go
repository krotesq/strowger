package account

import (
	"net/netip"
	"time"
)

type account struct {
	ID                  string    `db:"id"`
	Username            string    `db:"username"`
	PasswordHash        string    `db:"password_hash"`
	Active              bool      `db:"active"`
	FailedLoginAttempts int       `db:"failed_login_attempts"`
	CreatedAt           time.Time `db:"created_at"`
}

type refreshToken struct {
	ID        string     `db:"id"`
	AccountID string     `db:"account_id"`
	TokenHash []byte     `db:"token_hash"`
	ExpiresAt time.Time  `db:"expires_at"`
	RevokedAt *time.Time `db:"revoked_at"`
	UserAgent string     `db:"user_agent"`
	IPAddress netip.Addr `db:"ip_address"`
	CreatedAt time.Time  `db:"created_at"`
}
