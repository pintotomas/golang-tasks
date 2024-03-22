package models

import (
	"testing"
	"time"
)

func TestValidateClientLiveEvent(t *testing.T) {
	tests := []struct {
		name            string
		event           *ClientLiveEvent
		successExpected bool
	}{
		{
			name: "ValidClientLiveEvent",
			event: &ClientLiveEvent{
				ClientTraceID:  "1234567890",
				Name:           "Test Client Live Event",
				Description:    "This is a test client live event",
				AdditionalInfo: "Additional information for the test event",
				StartDate:      time.Now(),
				Duration:       time.Hour,
			},
			successExpected: true,
		},
		{
			name: "MissingOptionalField",
			event: &ClientLiveEvent{
				Name:        "Test Client Live Event",
				Description: "This is a test client live event",
				StartDate:   time.Now(),
				Duration:    time.Hour,
			},
			successExpected: false,
		},
		{
			name: "NameExceedsMaxLength",
			event: &ClientLiveEvent{
				ClientTraceID:  "1234567890",
				Name:           "This name exceeds the maximum allowed length of twenty-five characters",
				Description:    "This is a test event",
				AdditionalInfo: "Additional information for the test event",
				StartDate:      time.Now(),
				Duration:       time.Hour,
			},
			successExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateClientLiveEvent(test.event)

			if (err == nil) != test.successExpected {
				t.Errorf("Test case '%s' failed: successExpected validation result %t, but got %t", test.name, test.successExpected, err == nil)
			}
		})
	}
}
