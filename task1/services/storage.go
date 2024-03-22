package services

import (
	"fmt"
	"golang-tasks/task1/models"
	"golang-tasks/task1/models/db"
	"time"
)

// StorageService represents a storage service for live events
type StorageService struct {
}

// SaveLiveEvent saves the given live event and returns a live event entity
func (s *StorageService) SaveLiveEvent(event models.LiveEventInterface) *db.LiveEventEntity {
	currentTime := time.Now()
	e := &db.LiveEventEntity{
		Title:       event.GetTitle(),
		Description: event.GetDescription(),
		Timezone:    event.GetTimezone(),
		StartDate:   event.GetStartDate(),
		EndDate:     event.GetEndDate(),
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}
	fmt.Println("Saving live event: " + e.Title)
	// Here you would typically save the entity to a database or some persistent storage
	return e
}
