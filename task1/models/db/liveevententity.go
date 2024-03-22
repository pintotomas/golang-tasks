package db

import "time"

// LiveEventEntity represents a live event database entity
type LiveEventEntity struct {
	ID          int
	Title       string
	Description string
	Timezone    string
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
