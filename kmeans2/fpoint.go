package kmeans2

import (
	"fmt"
	"math"
)

// FPoint implements the point interface.
type FPoint []float64

// Add ...
func (p FPoint) Add(q Point) Point {
	fq := q.(FPoint)
	if len(p) != len(fq) {
		panic("dimension mismatch")
	}

	for i := 0; i < len(p); i++ {
		p[i] += fq[i]
	}

	return p
}

func (p FPoint) Mult(v float64) Point {
	for i := 0; i < len(p); i++ {
		p[i] *= v
	}

	return p
}

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
	var (
		fq, ok = q.(FPoint)
		i      int
	)

	if ok && len(p) == len(fq) {
		for ; i < len(p) && p[i] == fq[i]; i++ {
		}
	}

	return ok && i == len(p) && i == len(fq)
}

// Get ...
func (p FPoint) Get(i int) interface{} {
	return p[i]
}

// Set ...
func (p FPoint) Set(i int, v interface{}) Point {
	p[i] = v.(float64)
	return p
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
