package source

import "time"

type source struct {
	ID           string    `db:"uuid"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	SourceTypeID string    `db:"source_type_id"`
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
	SourceID  string    `db:"source_id"`
	URL       string    `db:"url"`
	StreamKey string    `db:"key"`
	CreatedAt time.Time `db:"created_at"`
}
