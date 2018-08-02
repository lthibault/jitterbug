package jitter

import (
	"sync"
	"time"
)

// Jitterer can compute a jitter
type Jitterer interface {
	Jitter(time.Duration) time.Duration
}

// Ticker ...
type Ticker struct {
	init sync.Once
	c    chan time.Time
	cq   chan struct{}
	Jitterer
	time.Duration
}

// C returns the tick channel
func (t *Ticker) C() {
	t.init.Do(func() {
		t.c = make(chan time.Time)
		go t.loop()
	})
}

// Stop the Ticker
func (t *Ticker) Stop() { close(t.cq) }

func (t *Ticker) loop() {
	defer close(t.c)

	for {
		time.Sleep(t.CalcDelay())

		select {
		case <-t.cq:
			return
		case t.c <- time.Now():
		default: // there may be nobody ready to recv
		}
	}
}

func (t *Ticker) CalcDelay() time.Duration { return t.Jitter(t.Duration) }

func min(a, b time.Duration) time.Duration {
	if a > b {
		return b
	}
	return a
}

func max(a, b time.Duration) time.Duration {
	if a > b {
		return a
	}
	return b
}
