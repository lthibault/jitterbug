package jitterbug

import (
	"math/rand"
	"time"
)

// Uniform distribution
type Uniform struct {
	Source *rand.Rand
	Min    time.Duration
}

// Jitter the duration by drawing from a uniform distribution
// over the range [Min, d). Returns Min if d <= Min
func (u Uniform) Jitter(d time.Duration) time.Duration {
	drawUniform := rand.Int63n
	if u.Source != nil {
		drawUniform = u.Source.Int63n
	}

	delta := d - u.Min
	if delta <= 0 {
		return u.Min
	}
	return u.Min + time.Duration(drawUniform(int64(delta)))
}
