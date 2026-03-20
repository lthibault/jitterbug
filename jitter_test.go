package jitterbug

import (
	"testing"
	"testing/synctest"
	"time"

	"go.uber.org/goleak"
)

type constantTimeJitter struct {
	jitter time.Duration
}

func (m *constantTimeJitter) Jitter(d time.Duration) time.Duration {
	return d + m.jitter
}

func TestNew(t *testing.T) {
	defer goleak.VerifyNone(t)

	interval := 100 * time.Millisecond
	j := &constantTimeJitter{}

	ticker := New(interval, j)
	if ticker == nil {
		t.Fatal("New() returned nil ticker")
	}
	defer ticker.Stop()

	if ticker.Interval != interval {
		t.Errorf("expected interval %v, got %v", interval, ticker.Interval)
	}
	if ticker.Jitter != j {
		t.Error("jitter implementation not set correctly")
	}
	if ticker.C == nil {
		t.Error("ticker channel C is nil")
	}
}

func TestTicker_Stop(t *testing.T) {
	defer goleak.VerifyNone(t)

	interval := 10 * time.Millisecond
	j := &constantTimeJitter{}

	ticker := New(interval, j)
	channelClosed := false

	synctest.Test(t, func(t *testing.T) {
		synctest.Wait() // Give the ticker time to start
		ticker.Stop()
		synctest.Wait() // Give it time to stop
		// After Stop(), the channel should eventually close
		select {
		case _, ok := <-ticker.C:
			if !ok {
				channelClosed = true
			}
		case <-time.After(200 * time.Millisecond):
			t.Fatal("ticker channel did not close after Stop()")
		}
	})

	if !channelClosed {
		t.Fatal("ticker channel should have been closed")
	}
}

func TestTicker_Ticking(t *testing.T) {
	defer goleak.VerifyNone(t)

	interval := 50 * time.Millisecond
	j := &constantTimeJitter{}

	ticker := New(interval, j)
	defer ticker.Stop()

	tickCount := 0

	// Wait for  3 ticks
	synctest.Test(t, func(t *testing.T) {
		synctest.Wait()
		select {
		case tick := <-ticker.C:
			if tick.IsZero() {
				t.Error("received zero time value")
			}
			tickCount++
		}
		synctest.Wait()
		select {
		case tick := <-ticker.C:
			if tick.IsZero() {
				t.Error("received zero time value")
			}
			tickCount++
		}
		synctest.Wait()
		select {
		case tick := <-ticker.C:
			if tick.IsZero() {
				t.Error("received zero time value")
			}
			tickCount++
		}

	})

	if tickCount != 3 {
		t.Fatalf("expected 3 ticks, got %d", tickCount)
	}
}

func TestTicker_MultipleStops(t *testing.T) {
	defer goleak.VerifyNone(t)

	interval := 50 * time.Millisecond
	j := &constantTimeJitter{}

	ticker := New(interval, j)

	synctest.Test(t, func(t *testing.T) {
		ticker.Stop()   // First stop should work fine
		synctest.Wait() // Wait for channel to close
		ticker.Stop()   // Close again should also work fine
	})
}
