package jitterbug

import (
	"math/rand"
	"time"
)

// Uniform distribution
type Uniform struct {
	Source *rand.Rand
	Limit  time.Duration
}

// Jitter the duration by drawing from a uniform distribution
func (u Uniform) Jitter(d time.Duration) time.Duration {
	f := rand.Int63n
	if u.Source != nil {
		f = u.Source.Int63n
	}

	sign := rand.Intn
	if u.Source != nil {
		sign = u.Source.Intn
	}

	samp := time.Duration(f(int64(u.Limit)))

	if sign(1) == 0 {
		return d - samp
	}

	return d + samp
}
