package jitter

import (
	"math/rand"
	"time"
)

// Norm is a Jitterer that draws from a normal distribution.
type Norm struct {
	Source      *rand.Rand
	Mean, Stdev time.Duration
}

// Jitter the duration by drawing form a normal distribution
func (n Norm) Jitter(d time.Duration) time.Duration {
	f := rand.NormFloat64
	if n.Source != nil {
		f = n.Source.NormFloat64
	}

	samp := time.Duration(f())
	return d * (samp*n.Stdev + n.Mean)
}
