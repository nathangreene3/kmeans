package kmeans

import "math"

// FPoint is an n-dimensional point in n-space.
type FPoint []float64

// At returns the ith value.
func (p FPoint) At(i int) float64 {
	return p[i]
}

// CompareTo returns -1, 0, or 1 indicating point 0 precedes, is equal to, or follows point 1.
func (p FPoint) CompareTo(q Point) int {
	n := len(p)
	if n != q.Len() {
		panic("dimension mismatch")
	}

	for i := 0; i < n; i++ {
		switch {
		case p[i] < q.At(i):
			return -1
		case q.At(i) < p[i]:
			return 1
		}
	}

	return 0
}

// Copy a point.
func (p FPoint) Copy() Point {
	cpy := make(FPoint, len(p))
	copy(cpy, p)
	return cpy
}

// Dist returns the Euclidean Dist between two points.
func (p FPoint) Dist(pnt Point) float64 {
	return math.Sqrt(p.SqDist(pnt))
}

// Length returns the number of dimensions of a point.
func (p FPoint) Len() int {
	return len(p)
}

// SqDist returns the squared distance between two points.
func (p FPoint) SqDist(pnt Point) float64 {
	n := len(p)
	if n != pnt.Len() {
		panic("dimension mismatch")
	}

	var sd, d float64 // Squared distance, Difference in each dimension
	for i := 0; i < n; i++ {
		d = p[i] - pnt.At(i)
		sd += d * d
	}

	return sd
}
