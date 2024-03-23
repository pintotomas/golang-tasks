package task2

import (
	"sync"
	"sync/atomic"
	"time"
)

// Worker struct
type Worker struct {
	ID                          int
	Limiter                     *APILimiter
	Requests                    <-chan int
	Backoff                     <-chan struct{}
	BackoffNotificationChannels []chan<- struct{}
	Throttle                    *int32 // Throttle count
	Success                     *int32 // Success count
	wg                          *sync.WaitGroup
}

// Process worker process
func (w *Worker) Process() {
	defer w.wg.Done()

	for {
		select {
		case <-w.Backoff:
			// Back off
			time.Sleep(time.Second)
		case _, isOpen := <-w.Requests:
			if isOpen {
				if !w.Limiter.Allow() {
					atomic.AddInt32(w.Throttle, 1)
					// Notification to other workers
					for _, ch := range w.BackoffNotificationChannels {
						// attempt a non-blocking send operation to avoid a deadlock
						select {
						case ch <- struct{}{}:
							// Value sent successfully
						default:
							// Channel is full, Coworker was already notified
						}
					}
				} else {
					atomic.AddInt32(w.Success, 1)
				}
				// Simulate sending request to AWS
				time.Sleep(100 * time.Millisecond)
			} else {
				return
			}
		default:
			return
		}
	}
}
