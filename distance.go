package main

import "math"

// SquaredDistance returns the Euclidean metric on two points.
func SquaredDistance(pnt0, pnt1 Point) float64 {
	var (
		sd float64 // Squared distance
		d  float64 // Difference in each dimension
	)
	for i := range pnt0 {
		d = pnt0[i] - pnt1[i]
		sd += d * d
	}

	return sd
}

// Distance returns the Euclidean Distance between two points.
func Distance(pnt0, pnt1 Point) float64 {
	return math.Sqrt(SquaredDistance(pnt0, pnt1))
}
