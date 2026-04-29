package target

import "time"

type target struct {
	ID           string    `db:"uuid"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	TargetTypeID string    `db:"target_type_id"`
	AccountID    string    `db:"account_id"`
	CreatedAt    time.Time `db:"created_at"`
}

type _type struct {
	ID          string    `db:"uuid"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

type rtmp struct {
	TargetID  string    `db:"target_id"`
	URL       string    `db:"url"`
	StreamKey string    `db:"key"`
	CreatedAt time.Time `db:"created_at"`
}
