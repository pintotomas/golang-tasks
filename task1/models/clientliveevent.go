package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// ClientLiveEvent represents a live event for a client
type ClientLiveEvent struct {
	ClientTraceID  string        `json:"client_trace_id" validate:"required"`
	Name           string        `json:"name" validate:"required,max=25"`
	Description    string        `json:"description" validate:"required,max=2000"`
	AdditionalInfo string        `json:"additional_info" validate:"max=2000"`
	StartDate      time.Time     `json:"start_date" validate:"required"`
	Duration       time.Duration `json:"duration" validate:"required"`
}

// ValidateClientLiveEvent validates the ClientLiveEvent struct
func ValidateClientLiveEvent(event *ClientLiveEvent) error {
	validate := validator.New()
	err := validate.Struct(event)
	return err
}
