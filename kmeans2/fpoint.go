package kmeans2

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
	n := len(p)
	switch {
	case n == 0:
		if x == nil {
			return 0
		}
		return 1
	case x == nil:
		return -1
	case n != x.Len():
		panic("dimension mismatch")
	}

	q := x.(FPoint)
	for i := 0; i < n; i++ {
		pi, xi := p[i], q[i]
		switch {
		case pi < xi:
			return -1
		case xi < pi:
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
	var (
		n = len(p)
		q = x.(FPoint)
	)

	if n != len(q) {
		panic("dimension mismatch")
	}

	var sd float64
	for i := 0; i < n; i++ {
		d := p[i] - q[i]
		sd += d * d
	}

	return math.Sqrt(sd)
}

// Len returns the number of dimensions of a point.
func (p FPoint) Len() int {
	return len(p)
}
