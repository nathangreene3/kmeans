package kmeans

import "math"

// FPoint is an n-dimensional point in n-space. It implements the k-means point
// interface.
type FPoint []float64

// At returns the ith value.
func (p FPoint) At(i int) float64 {
	return p[i]
}

// Compare returns -1, 0, or 1 indicating point 0 precedes, is equal to, or
// follows point 1.
func (p FPoint) Compare(x Interface) int {
	switch {
	case len(p) == 0:
		if x == nil {
			return 0
		}
		return 1
	case x == nil:
		return -1
	}

	n := len(p)
	if n != x.Len() {
		panic("dimension mismatch")
	}

	for i := 0; i < n; i++ {
		switch {
		case p[i] < x.At(i):
			return -1
		case x.At(i) < p[i]:
			return 1
		}
	}

	return 0
}

// Copy a point.
func (p FPoint) Copy() Interface {
	point := make(FPoint, len(p))
	copy(point, p)
	return point
}

// Dist returns the Euclidean Dist between two points.
func (p FPoint) Dist(x Interface) float64 {
	n := len(p)
	if n != x.Len() {
		panic("dimension mismatch")
	}

	var sqDist, diffAt float64
	for i := 0; i < n; i++ {
		diffAt = p[i] - x.At(i)
		sqDist += diffAt * diffAt
	}

	return math.Sqrt(sqDist)
}

// Len returns the number of dimensions of a point.
func (p FPoint) Len() int {
	return len(p)
}
