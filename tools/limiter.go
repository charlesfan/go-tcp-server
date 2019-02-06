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
	defer l.mux.Unlock()
	l.mux.Lock()
	du := time.Since(l.startTime)
	if du > l.duration {
		l.startTime = time.Now()
		l.count = 0
		return true
	}

	if l.count < l.max {
		l.count = l.count + 1
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
