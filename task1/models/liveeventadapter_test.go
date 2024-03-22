package models

import (
	"testing"
	"time"
)

func TestLiveEventAdapter_GetTitle(t *testing.T) {
	event := &ClientLiveEvent{
		Name: "Test Event",
	}
	adapter := &LiveEventAdapter{ClientLiveEvent: event}

	title := adapter.GetTitle()

	if title != "Test Event" {
		t.Errorf("Expected title to match 'Test Event', got %s", title)
	}
}

func TestLiveEventAdapter_GetDescription(t *testing.T) {
	event := &ClientLiveEvent{
		Description:    "Test Description",
		AdditionalInfo: "Additional Info",
	}
	adapter := &LiveEventAdapter{ClientLiveEvent: event}

	description := adapter.GetDescription()

	expectedDescription := "Test Description Additional Info"
	if description != expectedDescription {
		t.Errorf("Expected description to match '%s', got '%s'", expectedDescription, description)
	}
}

func TestLiveEventAdapter_GetTimezone(t *testing.T) {
	event := &ClientLiveEvent{}
	adapter := &LiveEventAdapter{ClientLiveEvent: event}

	timezone := adapter.GetTimezone()

	expectedTimezone := time.UTC.String()
	if timezone != expectedTimezone {
		t.Errorf("Expected timezone to match '%s', got '%s'", expectedTimezone, timezone)
	}
}

func TestLiveEventAdapter_GetStartDate(t *testing.T) {
	startDate := time.Now()
	event := &ClientLiveEvent{StartDate: startDate}
	adapter := &LiveEventAdapter{ClientLiveEvent: event}

	startDateFromAdapter := adapter.GetStartDate()

	if !startDateFromAdapter.Equal(startDate) {
		t.Errorf("Expected start dates to match, got %s and %s", startDate, startDateFromAdapter)
	}
}

func TestLiveEventAdapter_GetEndDate(t *testing.T) {
	startDate := time.Now()
	duration := time.Hour
	event := &ClientLiveEvent{StartDate: startDate, Duration: duration}
	adapter := &LiveEventAdapter{ClientLiveEvent: event}

	endDate := adapter.GetEndDate()

	expectedEndDate := startDate.Add(duration)
	if !endDate.Equal(expectedEndDate) {
		t.Errorf("Expected end dates to match, got %s and %s", expectedEndDate, endDate)
	}
}

func TestLiveEventAdapter_CalculateDuration(t *testing.T) {
	duration := time.Hour
	event := &ClientLiveEvent{Duration: duration}
	adapter := &LiveEventAdapter{ClientLiveEvent: event}

	calculatedDuration := adapter.CalculateDuration()

	if calculatedDuration != duration {
		t.Errorf("Expected durations to match, got %s and %s", duration, calculatedDuration)
	}
}
