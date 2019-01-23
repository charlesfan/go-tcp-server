package tools

import (
	"sync"
	"time"
)

type Limiter struct {
	startTime time.Time
	duration  time.Duration
	count     int
	max       int
	mux       sync.Mutex
}

func (l *Limiter) Limit() bool {
	du := time.Since(l.startTime)
	if du > l.duration {
		l.mux.Lock()
		l.startTime = time.Now()
		l.count = 0
		l.mux.Unlock()
		return true
	}

	if l.count < l.max {
		l.mux.Lock()
		l.count = l.count + 1
		l.mux.Unlock()
		return true
	}

	return false
}

func NewLimiter(du time.Duration, max int) *Limiter {
	return &Limiter{
		startTime: time.Now(),
		duration:  du,
		count:     0,
		max:       max,
	}
}
