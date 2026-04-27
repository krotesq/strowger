package account

import "time"

type account struct {
	ID                  string    `db:"id"`
	Username            string    `db:"username"`
	PasswordHash        string    `db:"password_hash"`
	Active              bool      `db:"active"`
	FailedLoginAttempts int       `db:"failed_login_attempts"`
	CreatedAt           time.Time `db:"created_at"`
}
