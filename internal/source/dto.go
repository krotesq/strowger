package source

import "time"

type sourceDto struct {
	UUID         string    `json:"uuid"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	SourceTypeID string    `json:"sourceTypeId"`
	AccountID    string    `json:"accountId"`
	CreatedAt    time.Time `json:"createdAt"`
}

type rtmpDto struct {
	SourceID  string    `json:"sourceId"`
	URL       string    `json:"url"`
	StreamKey string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
}

type sourceWithDetailsDto struct {
	Source sourceDto `json:"source"`
	// add all available types here
	Rtmp *rtmpDto `json:"rtmp"`
}
