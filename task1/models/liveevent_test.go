package models

import (
	"testing"
	"time"
)

func TestLiveEventValidation(t *testing.T) {
	tests := []struct {
		name            string
		event           *LiveEvent
		successExpected bool
	}{
		{
			name: "ValidEvent",
			event: &LiveEvent{
				Title:       "Example Live Event",
				Description: "This is an example of a live event",
				Timezone:    "UTC",
				StartDate:   time.Now(),
				EndDate:     time.Now().Add(time.Hour),
			},
			successExpected: true,
		},
		{
			name: "EmptyTitle",
			event: &LiveEvent{
				Description: "This is an example of a live event",
				Timezone:    "UTC",
				StartDate:   time.Now(),
				EndDate:     time.Now().Add(time.Hour),
			},
			successExpected: false,
		},
		{
			name: "TitleExceedsMaxLength",
			event: &LiveEvent{
				Title:       "This title exceeds the maximum allowed length of twenty-five characters",
				Description: "This is an example of a live event",
				Timezone:    "UTC",
				StartDate:   time.Now(),
				EndDate:     time.Now().Add(time.Hour),
			},
			successExpected: false,
		},
		{
			name: "ValidTimezone",
			event: &LiveEvent{
				Title:       "Valid Timezone Event",
				Description: "This is an example of a live event",
				Timezone:    "America/New_York",
				StartDate:   time.Now(),
				EndDate:     time.Now().Add(time.Hour),
			},
			successExpected: true,
		},
		{
			name: "InvalidTimezone",
			event: &LiveEvent{
				Title:       "Invalid Timezone Event",
				Description: "This is an example of a live event",
				Timezone:    "Invalid/Timezone",
				StartDate:   time.Now(),
				EndDate:     time.Now().Add(time.Hour),
			},
			successExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateLiveEvent(test.event)

			if (err == nil) != test.successExpected {
				t.Errorf("Test case '%s' failed: successExpected validation result %t, but got %t", test.name, test.successExpected, err == nil)
			}
		})
	}
}

func TestCalculateDuration(t *testing.T) {
	start := time.Date(2024, time.March, 21, 9, 0, 0, 0, time.UTC)
	end := time.Date(2024, time.March, 21, 12, 0, 0, 0, time.UTC)

	event := LiveEvent{
		Title:       "Example Live Event",
		Description: "This is an example of a live event",
		Timezone:    "UTC",
		StartDate:   start,
		EndDate:     end,
	}

	expectedDuration := 3 * time.Hour

	actualDuration := event.CalculateDuration()

	if actualDuration != expectedDuration {
		t.Errorf("Expected duration: %s, but got: %s", expectedDuration, actualDuration)
	}
}
