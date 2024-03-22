package task1

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"golang-tasks/task1/models"
	"golang-tasks/task1/services"
)

// Run executes task 1 example
func Run() {
	xmlData := `
	<LiveEvent>
		<title>Example Event (XML)</title>
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

	if service.SaveLiveEvent(event) != nil {
		fmt.Println("Successfully saved the live event!")
	} else {
		fmt.Println("Something went wrong saving the live event")
	}

	// JSON data representing the client live event
	jsonData := `
	{
		"client_trace_id": "1234567890",
		"name": "Test Client Event",
		"description": "This is a test event",
		"additional_info": "Additional information for the test event",
		"start_date": "2024-03-25T12:00:00Z",
		"duration": 10000
	}
	`

	// Initialize a ClientLiveEvent instance from JSON
	clientLiveEvent := &models.ClientLiveEvent{}
	err = json.Unmarshal([]byte(jsonData), clientLiveEvent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a LiveEventAdapter wrapping the ClientLiveEvent
	adapter := &models.LiveEventAdapter{ClientLiveEvent: clientLiveEvent}

	if service.SaveLiveEvent(adapter) != nil {
		fmt.Println("Successfully saved the client live event!")
	} else {
		fmt.Println("Something went wrong saving the client live event")
	}

}
