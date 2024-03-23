package task2

import (
	"fmt"
	"sync"
	"time"
)

const (
	Workers                     = 3
	Requests                    = 50
	APILimiterRequestsPerSecond = 10
)

// Run task2
func Run() {
	fmt.Println("Processing requests..")
	limiter := NewAPILimiter(APILimiterRequestsPerSecond, time.Second) // Example rate Limit is 10 Requests per second
	numWorkers := Workers
	numRequests := Requests
	requestCh := make(chan int, numRequests)
	var wg sync.WaitGroup
	var throttledCount int32 // Throttled count across all workers
	var successCount int32   // Success count across all workers

	// Create backoff channels for each worker
	var backoffChannels [3]chan struct{}
	for i := range backoffChannels {
		backoffChannels[i] = make(chan struct{}, 1)
	}

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {

		// Create a slice to store coworkers backoff channels
		var coworkersBackoffChannels []chan<- struct{}
		for _, ch := range backoffChannels {
			coworkersBackoffChannels = append(coworkersBackoffChannels, ch)
		}

		worker := &Worker{
			ID:                          i,
			Limiter:                     limiter,
			Throttle:                    &throttledCount,
			Success:                     &successCount,
			Backoff:                     backoffChannels[i],
			BackoffNotificationChannels: coworkersBackoffChannels,
			Requests:                    requestCh,
			wg:                          &wg,
		}
		wg.Add(1)
		go worker.Process()
	}

	// Enqueue Requests
	for i := 0; i < numRequests; i++ {
		requestCh <- i
	}
	close(requestCh)

	// Wait for all workers to finish
	wg.Wait()

	// Print results
	fmt.Printf("Succeeded: %d\n", successCount)
	fmt.Printf("Throttled: %d\n", throttledCount)
}
