package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// LiveEvent struct representing a live event
type LiveEvent struct {
	Title       string    `xml:"title" validate:"required,max=25"`
	Description string    `xml:"description" validate:"required,max=2000"`
	Timezone    string    `xml:"timezone" validate:"required,max=30,validTimezone"`
	StartDate   time.Time `xml:"start_date" validate:"required"`
	EndDate     time.Time `xml:"end_date" validate:"required"`
}

func (e *LiveEvent) CalculateDuration() time.Duration {
	return e.EndDate.Sub(e.StartDate)
}

// ValidateLiveEvent Helper function to validate a LiveEvent instance
func ValidateLiveEvent(event *LiveEvent) error {
	validate := validator.New()
	err := validate.RegisterValidation("validTimezone", validateTimezone)
	if err != nil {
		return err
	}
	err = validate.Struct(event)
	return err
}

// Custom validation function for timezone
func validateTimezone(fl validator.FieldLevel) bool {
	timezone := fl.Field().String()
	_, err := time.LoadLocation(timezone)
	return err == nil
}
