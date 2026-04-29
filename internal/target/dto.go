package target

import "time"

type targetDto struct {
	ID           string    `json:"uuid"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	TargetTypeID string    `json:"targetTypeId"`
	AccountID    string    `json:"accountId"`
	CreatedAt    time.Time `json:"createdAt"`
}

type rtmpDto struct {
	TargetID  string    `json:"targetId"`
	URL       string    `json:"url"`
	StreamKey string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
}

type targetWithDetailsDto struct {
	Target targetDto `json:"target"`
	// add all available types here
	Rtmp *rtmpDto `json:"rtmp"`
}
