package kmeans

import (
	"math/rand"
	"time"
)

var (
	// seeded indicates if the random number generator has been seeded.
	seeded bool
	// seed is the seed used to seed the random number generator.
	seed int64
)

// max returns the maximum value.
func max(values ...int) int {
	var (
		maximum = values[0]
		n       = len(values)
		v       int
	)

	for i := 1; i < n; i++ {
		if v = values[i]; maximum < v {
			maximum = v
		}
	}

	return maximum
}

// maxPow returns the largest power p such that b^p <= n for a given base b > 0. Assumes b,n > 0.
func maxPow(b, n int) int {
	var p int // Power to return
	for bp := b; bp <= n; bp *= b {
		// bp = b^p
		p++
	}

	return p
}

// mean returns the mean of a set of numbers. Assumes x is not empty.
func mean(x ...float64) float64 {
	var v float64 // Sum of values in x
	for i := range x {
		v += x[i]
	}

	return v / float64(len(x))
}

// min returns the minimum value.
func min(values ...int) int {
	var (
		minimum = values[0]
		n       = len(values)
		v       int
	)

	for i := 1; i < n; i++ {
		if v = values[i]; v < minimum {
			minimum = v
		}
	}

	return minimum
}

// seedRNG the random number generator.
func seedRNG() {
	if !seeded {
		seed = int64(time.Now().Nanosecond())
		rand.Seed(seed)
		seeded = true
	}
}

// roundDown to the nearest integer.
func roundDown(x float64) int {
	return int(x)
}

// roundUp to the nearest integer.
func roundUp(x float64) int {
	n := int(x)
	if float64(n) < x {
		n++
	}

	return n
}

// roundUpToMult rounds x up to the next multiple of n.
func roundUpToMult(x float64, n int) int {
	m := int(x)
	if mod := m % n; 0 < mod {
		m += n - mod
	}

	return m
}
