package models

import "time"

// LiveEventInterface represents the interface for a live event
type LiveEventInterface interface {
	GetTitle() string
	GetDescription() string
	GetTimezone() string
	GetStartDate() time.Time
	GetEndDate() time.Time
	CalculateDuration() time.Duration
}
