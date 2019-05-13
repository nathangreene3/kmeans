package main

import (
	"math/rand"
	"time"
)

// seed the random number generator.
func seed() {
	if !seeded {
		rand.Seed(int64(time.Now().Nanosecond()))
		seeded = true
	}
}

// maxPow returns the largest power p such that b^p <= n for a given base b > 0. Assumes b,n > 0.
func maxPow(b, n int) int {
	var p int                       // Power to return
	for bp := b; bp <= n; bp *= b { // bp = b^p
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
