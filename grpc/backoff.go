package grpc

import (
	"sync"
	"time"
)

type Func func() time.Duration

// given (100ms, 1s) => [100ms, 200ms, 400ms, 800ms, 1s, 1s, 1s, ...]
func ExponentialWithCappedMax(base time.Duration, max time.Duration) Func {
	idx := uint(0)
	idxMu := sync.Mutex{}

	return func() time.Duration {
		idxMu.Lock()
		defer idxMu.Unlock()

		var multiplier time.Duration
		if idx == 0 {
			multiplier = 1
		} else {
			multiplier = 2 << (idx - 1)
		}

		sleepDuration := multiplier * base

		if sleepDuration > max {
			sleepDuration = max
		} else {
			idx++
		}

		return sleepDuration
	}
}
