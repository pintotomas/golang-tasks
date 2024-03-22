package task2

import (
	"sync"
	"sync/atomic"
	"time"
)

type Worker struct {
	ID       int
	Limiter  *APILimiter
	Requests <-chan int
	Throttle *int32 // Throttle count
	Success  *int32 // Success count
	wg       *sync.WaitGroup
}

func (w *Worker) Process() {
	defer w.wg.Done()
	for range w.Requests {
		if !w.Limiter.Allow() {
			atomic.AddInt32(w.Throttle, 1)
		} else {
			atomic.AddInt32(w.Success, 1)
		}
		// Simulate sending request to AWS
		time.Sleep(100 * time.Millisecond)
	}
}
