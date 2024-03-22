package models

import (
	"time"
)

// LiveEventAdapter is a wrapper for ClientLiveEvent implementing LiveEventInterface
type LiveEventAdapter struct {
	ClientLiveEvent *ClientLiveEvent
}

// GetTitle returns the name of the live event
func (a *LiveEventAdapter) GetTitle() string {
	return a.ClientLiveEvent.Name
}

// GetDescription returns the description of the live event
func (a *LiveEventAdapter) GetDescription() string {
	return a.ClientLiveEvent.Description + " " + a.ClientLiveEvent.AdditionalInfo
}

// GetTimezone returns the default timezone for the live event
func (a *LiveEventAdapter) GetTimezone() string {
	return time.UTC.String() // You can modify this to return a specific timezone if needed
}

// GetStartDate returns the start date of the live event
func (a *LiveEventAdapter) GetStartDate() time.Time {
	return a.ClientLiveEvent.StartDate
}

// GetEndDate returns the end date of the live event
func (a *LiveEventAdapter) GetEndDate() time.Time {
	return a.ClientLiveEvent.StartDate.Add(a.ClientLiveEvent.Duration)
}

// CalculateDuration calculates the duration of the live event
func (a *LiveEventAdapter) CalculateDuration() time.Duration {
	return a.ClientLiveEvent.Duration
}
