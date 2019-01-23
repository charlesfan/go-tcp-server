package tools_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/charlesfan/go-tcp-server/tools"
)

func testFn(ch chan int, stop chan bool, l *tools.Limiter) {
	for i := 0; i < 10; i++ {
		if l.Limit() {
			ch <- 1
		}
	}
	stop <- true
}

func TestLimiter(t *testing.T) {
	du, err := time.ParseDuration("1s")
	assert.NoError(t, err)
	max := 10
	limiter := tools.NewLimiter(du, max)

	ch := make(chan int)
	stop := make(chan bool)

	go testFn(ch, stop, limiter)
	go testFn(ch, stop, limiter)

	count := 0
	b := 0
	for {
		select {
		case i := <-ch:
			count += i
			break
		case <-stop:
			b++
		}
		if b == 2 {
			break
		}
	}
	assert.Equal(t, count, max)
}
