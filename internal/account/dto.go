package account

import "time"

type accountDto struct {
	UUID                string    `json:"uuid"`
	Username            string    `json:"username"`
	Active              bool      `json:"active"`
	FailedLoginAttempts int       `json:"failedLoginAttempts"`
	CreatedAt           time.Time `json:"createdAt"`
}

type loginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
