package task2

import (
	"sync"
	"testing"
	"time"
)

func TestAPILimiter_Allow(t *testing.T) {
	tests := []struct {
		name          string
		limit         int
		interval      time.Duration
		lastTimestamp time.Time
		requests      int
		expectedAllow bool
	}{
		{
			name:          "RequestWithinLimit",
			limit:         5,
			interval:      time.Second * 10,
			lastTimestamp: time.Now().Add(-time.Second),
			requests:      4,
			expectedAllow: true,
		},
		{
			name:          "RequestExceedsLimit",
			limit:         5,
			lastTimestamp: time.Now().Add(-time.Second),
			interval:      time.Second * 10,
			requests:      6,
			expectedAllow: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limiter := NewAPILimiter(tt.limit, tt.interval)
			limiter.Requests = tt.requests
			limiter.LastTimestamp = tt.lastTimestamp
			allowed := limiter.Allow()
			if allowed != tt.expectedAllow {
				t.Errorf("Allow() = %v, expected %v", allowed, tt.expectedAllow)
			}
		})
	}
}

func TestAPILimiter_Allow_IntervalReset(t *testing.T) {
	limiter := NewAPILimiter(5, time.Second)
	limiter.Requests = 4
	limiter.LastTimestamp = time.Now().Add(-2 * time.Second) // Set LastTimestamp to 2 seconds ago

	// Allow should reset the request count and update the LastTimestamp
	allowed := limiter.Allow()
	if !allowed || limiter.Requests != 1 || time.Since(limiter.LastTimestamp) > time.Millisecond {
		t.Errorf("Allow() did not reset Requests count or update LastTimestamp")
	}
}

func TestAPILimiter_Allow_ConcurrentAccess(t *testing.T) {
	limiter := NewAPILimiter(2, time.Second)
	var wg sync.WaitGroup
	wg.Add(2)

	// Two goroutines concurrently accessing Allow()
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			_ = limiter.Allow()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			_ = limiter.Allow()
		}
	}()

	wg.Wait()

	// Current Requests should not exceed Limit
	if limiter.Requests > limiter.Limit {
		t.Errorf("Total Requests exceeded Limit")
	}
}
