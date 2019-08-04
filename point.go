package kmeans

import (
	"math"
)

// Point is an n-dimensional point in n-space.
type Point []float64

// SqDist returns the squared distance between two points.
func (p Point) SqDist(pnt Point) float64 {
	n := len(p)
	if n != len(pnt) {
		panic("dimension mismatch")
	}

	var sd, d float64 // Squared distance, Difference in each dimension
	for i := 0; i < n; i++ {
		d = p[i] - pnt[i]
		sd += d * d
	}

	return sd
}

// Dist returns the Euclidean Dist between two points.
func (p Point) Dist(pnt Point) float64 {
	return math.Sqrt(p.SqDist(pnt))
}

// Normalize a point given a maximum representing point.
func (p Point) Normalize(maxPoint Point) {
	for i, v := range maxPoint {
		if 0 < v {
			p[i] /= v
		}
	}
}

// CompareTo returns -1, 0, or 1 indicating point 0 precedes, is equal to, or follows point 1.
func (p Point) CompareTo(q Point) int {
	n := len(p)
	if n != len(q) {
		panic("dimension mismatch")
	}

	for i := 0; i < n; i++ {
		switch {
		case p[i] < q[i]:
			return -1
		case q[i] < p[i]:
			return 1
		}
	}

	return 0
}

// Copy a point.
func (p Point) Copy() Point {
	cpy := make(Point, len(p))
	copy(cpy, p)
	return cpy
}
