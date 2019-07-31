package main

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

// seedRNG the random number generator.
func seedRNG() {
	if !seeded {
		seed = int64(time.Now().Nanosecond())
		rand.Seed(seed)
		seeded = true
	}
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
func mean(x []float64) float64 {
	if len(x) == 0 {
		panic("mean: cannot take the mean of an empty set")
	}

	var v float64 // Sum of values in x
	for i := range x {
		v += x[i]
	}

	return v / float64(len(x))
}

// max returns the maximum value.
func max(m, n int) int {
	if m < n {
		return n
	}

	return m
}

// min returns the minimum value.
func min(m, n int) int {
	if m < n {
		return m
	}

	return n
}
