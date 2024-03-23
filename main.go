package main

import (
	"flag"
	"fmt"
	"golang-tasks/task1"
	"golang-tasks/task2"
	"golang-tasks/task3"
	"os"
)

func main() {
	// Define flags
	task := flag.Int("task", 0, "Task number (1, 2, or 3)")
	url := flag.String("url", "", "URL for task 3")
	depth := flag.Int("depth", 1, "Depth for task 3")
	timeout := flag.Int("timeout", 10, "Timeout in seconds for task 3")

	// Parse command-line arguments
	flag.Parse()

	// Check if task number is provided
	if *task == 0 {
		fmt.Println("Error: Task number is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check if task 3 is selected and URL is provided
	if *task == 3 && *url == "" {
		fmt.Println("Error: URL is required for task 3")
		os.Exit(1)
	}

	// Execute the selected task
	switch *task {
	case 1:
		task1.Run()
	case 2:
		task2.Run()
	case 3:
		task3.Run(*url, *depth, *timeout)
	default:
		fmt.Println("Error: Invalid task number")
		os.Exit(1)
	}
}
