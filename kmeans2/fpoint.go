package kmeans2

import (
	"fmt"
	"math"
)

// FPoint ...
type FPoint []float64

// Copy ...
func (p FPoint) Copy() Point {
	return append(make(FPoint, 0, len(p)), p...)
}

// Dist ...
func (p FPoint) Dist(q Point) float64 {
	return math.Sqrt(p.SqDist(q))
}

// Equals ...
func (p FPoint) Equals(q Point) bool {
	fq := q.(FPoint)
	if len(p) != len(fq) {
		return false
	}

	for i := 0; i < len(p); i++ {
		if p[i] != fq[i] {
			return false
		}
	}

	return true
}

// SqDist ...
func (p FPoint) SqDist(q Point) float64 {
	fq := q.(FPoint)
	if len(p) != len(fq) {
		panic("dimension mismatch")
	}

	var sd float64
	for i := 0; i < len(p); i++ {
		d := p[i] - fq[i]
		sd += d * d
	}

	return sd
}

// String ...
func (p FPoint) String() string {
	return fmt.Sprintf("%g", p)
}
