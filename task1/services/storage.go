package services

import (
	"golang-tasks/task1/models"
	"golang-tasks/task1/models/db"
	"time"
)

// StorageService represents a storage service for live events
type StorageService struct {
	// You can add any additional fields if needed
}

// SaveLiveEvent saves the given live event and returns true on success
func (s *StorageService) SaveLiveEvent(event *models.LiveEvent) bool {
	currentTime := time.Now()
	_ = &db.LiveEventEntity{
		Title:       event.Title,
		Description: event.Description,
		Timezone:    event.Timezone,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}
	// Here you would typically save the entity to a database or some persistent storage
	return true
}
