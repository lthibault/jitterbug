package jitterbug

import (
	"math/rand"
	"time"
)

// Uniform distribution
type Uniform struct {
	Source *rand.Rand
	Limit  time.Duration

	// If true, only a positive jitter will be added to the base duration
	NonNegative bool
}

// Jitter the duration by drawing from a uniform distribution
func (u Uniform) Jitter(d time.Duration) time.Duration {
	f := rand.Int63n
	if u.Source != nil {
		f = u.Source.Int63n
	}

	samp := time.Duration(f(int64(u.Limit)))

	if u.NonNegative {
		sign := rand.Intn
		if u.Source != nil {
			sign = u.Source.Intn
		}

		if sign(1) == 0 {
			samp = -samp
		}
	}

	return d + samp
}
