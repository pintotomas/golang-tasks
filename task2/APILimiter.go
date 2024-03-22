package task2

import (
	"sync"
	"time"
)

// APILimiter struct
type APILimiter struct {
	Limit         int
	Interval      time.Duration
	Requests      int
	LastTimestamp time.Time
	mu            sync.Mutex
}

// NewAPILimiter APILimiter constructor
func NewAPILimiter(limit int, interval time.Duration) *APILimiter {
	return &APILimiter{
		Limit:    limit,
		Interval: interval,
	}
}

// Allow decides whether a request is allowed if it's within the limits in the time Interval
func (l *APILimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	if now.Sub(l.LastTimestamp) > l.Interval {
		l.Requests = 0
		l.LastTimestamp = now
	}

	if l.Requests < l.Limit {
		l.Requests++
		return true
	}

	return false
}
