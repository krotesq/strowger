package account

import "time"

type accountDTO struct {
	UUID                string    `json:"uuid"`
	Username            string    `json:"username"`
	Active              bool      `json:"active"`
	FailedLoginAttempts int       `json:"failedLoginAttempts"`
	CreatedAt           time.Time `json:"createdAt"`
}

type loginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type createDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
