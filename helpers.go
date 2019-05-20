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

// copyFloat64s returns a copy of a []float64.
func copyFloat64s(s []float64) []float64 {
	cpy := make([]float64, len(s))
	copy(cpy, s)
	return cpy
}
