package task1

import (
	"encoding/xml"
	"fmt"
	"golang-tasks/task1/models"
	"golang-tasks/task1/services"
)

// Run executes task 1 example
func Run() {
	xmlData := `
	<LiveEvent>
		<title>Example Event</title>
		<description>This is an example event</description>
		<timezone>UTC</timezone>
		<start_date>2024-03-21T09:00:00Z</start_date>
		<end_date>2024-03-21T12:00:00Z</end_date>
	</LiveEvent>`

	event := &models.LiveEvent{}
	err := xml.Unmarshal([]byte(xmlData), event)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Validate the event
	err = models.ValidateLiveEvent(event)
	if err != nil {
		fmt.Println("Validation Error:", err)
		return
	}

	service := services.StorageService{}

	if service.SaveLiveEvent(event) {
		fmt.Println("Successfully saved the live event!")
	} else {
		fmt.Println("Something went wrong saving the live event")
	}
}
