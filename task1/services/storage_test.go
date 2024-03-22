package services

import (
	"testing"
	"time"

	"golang-tasks/task1/models"
)

func TestSaveLiveEventWithLiveEventAdapter(t *testing.T) {
	adapter := &models.LiveEventAdapter{
		ClientLiveEvent: &models.ClientLiveEvent{
			ClientTraceID:  "12345",
			Name:           "Test Live Event Adapter",
			Description:    "Test Description",
			AdditionalInfo: "Additional Info",
			StartDate:      time.Now(),
			Duration:       time.Hour,
		},
	}

	storage := &StorageService{}

	savedEvent := storage.SaveLiveEvent(adapter)

	// Ensure that the event was saved
	if savedEvent == nil {
		t.Error("Expected event to be saved")
	}

	// Validate the saved event's fields
	if savedEvent.Title != adapter.GetTitle() {
		t.Errorf("Expected title to match '%s', got '%s'", adapter.GetTitle(), savedEvent.Title)
	}

	if savedEvent.Description != adapter.GetDescription() {
		t.Errorf("Expected description to match '%s', got '%s'", adapter.GetDescription(), savedEvent.Description)
	}

	if savedEvent.StartDate != adapter.GetStartDate() {
		t.Errorf("Expected start date to match '%s', got '%s'", adapter.GetStartDate().String(), savedEvent.StartDate.String())
	}

	if savedEvent.EndDate != adapter.GetEndDate() {
		t.Errorf("Expected end date to match '%s', got '%s'", adapter.GetEndDate().String(), savedEvent.EndDate.String())
	}

	if savedEvent.Timezone != adapter.GetTimezone() {
		t.Errorf("Expected end date to match '%s', got '%s'", adapter.GetTimezone(), savedEvent.Timezone)
	}
}

func TestSaveLiveEventWithLiveEvent(t *testing.T) {
	event := &models.LiveEvent{
		Title:       "Test Event",
		Description: "Test Description",
		Timezone:    "UTC",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(time.Hour),
	}

	storage := &StorageService{}

	savedEvent := storage.SaveLiveEvent(event)

	// Ensure that the event was saved
	if savedEvent == nil {
		t.Error("Expected event to be saved")
	}

	// Validate the saved event's fields
	if savedEvent.Title != event.GetTitle() {
		t.Errorf("Expected title to match '%s', got '%s'", event.GetTitle(), savedEvent.Title)
	}

	if savedEvent.Description != event.GetDescription() {
		t.Errorf("Expected description to match '%s', got '%s'", event.GetDescription(), savedEvent.Description)
	}

	if savedEvent.StartDate != event.GetStartDate() {
		t.Errorf("Expected start date to match '%s', got '%s'", event.GetStartDate().String(), savedEvent.StartDate.String())
	}

	if savedEvent.EndDate != event.GetEndDate() {
		t.Errorf("Expected end date to match '%s', got '%s'", event.GetEndDate().String(), savedEvent.EndDate.String())
	}

	if savedEvent.Timezone != event.GetTimezone() {
		t.Errorf("Expected end date to match '%s', got '%s'", event.GetTimezone(), savedEvent.Timezone)
	}
}
