package task2

import (
	"sync"
	"testing"
	"time"
)

// TestWorker_Process this test creates a worker which will send a total of 5 requests to a limiter with 3 reqs per second, so only 4 should success after one fails and we backoff
func TestWorker_Process(t *testing.T) {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := NewAPILimiter(3, time.Second)
	var wg sync.WaitGroup
	var throttleCount int32
	var successCount int32

	// Create a Worker instance
	worker := &Worker{
		ID:       1,
		Limiter:  limiter,
		Requests: requests,
		Throttle: &throttleCount,
		Success:  &successCount,
		wg:       &wg,
	}

	wg.Add(1)
	worker.Process()

	wg.Wait()

	// Ensure throttle count and success count are updated correctly
	expectedThrottle := int32(1)
	expectedSuccess := int32(4)

	if throttleCount != expectedThrottle {
		t.Errorf("Throttle count = %d; want %d", throttleCount, expectedThrottle)
	}
	if successCount != expectedSuccess {
		t.Errorf("Success count = %d; want %d", successCount, expectedSuccess)
	}
}
