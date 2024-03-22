package task2

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	limiter := NewAPILimiter(10, time.Second) // Example rate Limit is 10 Requests per second
	numWorkers := 3
	numRequests := 50
	requestCh := make(chan int, numRequests)
	var wg sync.WaitGroup
	var throttledCount int32 // Throttled count across all workers
	var successCount int32   // Success count across all workers

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		worker := &Worker{
			ID:       i,
			Limiter:  limiter,
			Throttle: &throttledCount,
			Success:  &successCount,
			Requests: requestCh,
			wg:       &wg,
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
