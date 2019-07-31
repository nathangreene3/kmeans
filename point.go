package main

import (
	"math"
	"math/rand"
	"sort"
)

// Point is an n-dimensional point in n-space.
type Point []float64

// Points is a set of points.
type Points []Point

// SqDist returns the squared distance between two points.
func (p Point) SqDist(pnt Point) float64 {
	var sd, d float64 // Squared distance, Difference in each dimension
	for i := range p {
		d = p[i] - pnt[i]
		sd += d * d
	}

	return sd
}

// Dist returns the Euclidean Dist between two points.
func (p Point) Dist(pnt Point) float64 {
	return math.Sqrt(p.SqDist(pnt))
}

// AssignPoint returns the index of the closest cluster mean to a given point.
func AssignPoint(pnt Point, mns Points) int {
	var (
		assignment int
		sd         float64
		minSD      = math.MaxFloat64
	)
	for i := range mns {
		sd = SquaredDistance(mns[i], pnt)
		if sd < minSD {
			minSD = sd
			assignment = i
		}
	}

	return assignment
}

// AssignedTo returns the index of the closest cluster mean to a given point.
func (p Point) AssignedTo(pnts Points) int {
	var (
		assignment int
		sd         float64
		minSD      = math.MaxFloat64
	)
	for i := range pnts {
		sd = SquaredDistance(pnts[i], p)
		if sd < minSD {
			minSD = sd
			assignment = i
		}
	}

	return assignment
}

// validate panics if there are no points or if any points are of unequal or zero dimension.
func validate(pnts Points) {
	numPnts := len(pnts)
	if numPnts == 0 {
		panic("validate: no points")
	}

	dims := len(pnts[0])
	if dims == 0 {
		panic("validate: dimensionless point")
	}

	for i := 1; i < numPnts; i++ {
		if dims != len(pnts[i]) {
			panic("validate: dimension mismatch")
		}
	}
}

// validate panics if there are no points or if any points are of unequal or zero dimension.
func (ps Points) validate() {
	numPnts := len(ps)
	if numPnts == 0 {
		panic("validate: no points")
	}

	dims := len(ps[0])
	if dims == 0 {
		panic("validate: dimensionless point")
	}

	for i := 1; i < numPnts; i++ {
		if dims != len(ps[i]) {
			panic("validate: dimension mismatch")
		}
	}
}

// NormalizePoint returns a point normalized by a maximum representing point. Assumes the point and max point are of equal length.
func NormalizePoint(pnt, maxPoint Point) Point {
	normPnt := make(Point, len(pnt))
	for i, v := range maxPoint {
		if 0 < v {
			normPnt[i] = pnt[i] / v
		}
	}

	return normPnt
}

// Normalize a point given a maximum representing point.
func (p Point) Normalize(maxPoint Point) {
	for i, v := range maxPoint {
		if 0 < v {
			p[i] /= v
		}
	}
}

// NormalizePoints each dimension in each point to the range [-1,1] assuming the largest value for each dimension is within the scope of the points provided.
func NormalizePoints(pnts Points) Points {
	maxPnt := pnts.MaxRep()

	// Check if max point is normal. If it is, the points are already normalized.
	var notNormal bool
	for i := range maxPnt {
		if 1 < maxPnt[i] {
			notNormal = true
			break
		}
	}

	if notNormal {
		normPnts := make(Points, 0, len(pnts))
		for i := range pnts {
			normPnts = append(normPnts, NormalizePoint(pnts[i], maxPnt))
		}

		return normPnts
	}

	return CopyPoints(pnts)
}

// Normalize each dimension in each point to the range [-1,1] assuming the largest value for each dimension is within the scope of the points provided.
func (ps Points) Normalize() {
	// Check if max point is normal. If it is, the points are already normalized.
	maxPnt := ps.MaxRep()
	for i := range maxPnt {
		if 1 < maxPnt[i] {
			for j := range ps {
				ps[j].Normalize(maxPnt)
			}
			break
		}
	}
}

// randomPoint returns a random point in the space spanned by the maximum point on a set of points. Each dimension is a random value on (-r,r) where r is the maximum value in each dimension on the set of points. It does NOT return a point in the set.
func randomPoint(pnts Points) Point {
	seedRNG()
	pnt := pnts.MaxRep()
	for i := range pnt {
		pnt[i] *= 2*rand.Float64() - 1
	}

	return pnt
}

// Random returns a Random point in the space spanned by the maximum point on a set of points. Each dimension is a Random value on (-r,r) where r is the maximum value in each dimension on the set of points. It does NOT return a point in the set.
func (ps Points) Random() Point {
	seedRNG()
	maxPnt := ps.MaxRep()
	for i := range maxPnt {
		maxPnt[i] *= 2*rand.Float64() - 1
	}

	return maxPnt
}

// MaxRepPoint returns a point in which each dimension is the largest non-negative value observed in the set of points.
func MaxRepPoint(pnts Points) Point {
	if len(pnts) == 0 {
		return nil
	}

	maxPnt := make(Point, len(pnts[0]))
	for i := range pnts {
		for j, v := range pnts[i] {
			if v < 0 {
				v = -v
			}

			if maxPnt[j] < v {
				maxPnt[j] = v
			}
		}
	}

	return maxPnt
}

// MaxRep returns a point in which each dimension is the largest non-negative value observed in the set of points.
func (ps Points) MaxRep() Point {
	if len(ps) == 0 {
		return nil
	}

	var (
		maxPnt = make(Point, len(ps[0]))
		n      = len(maxPnt)
		v      float64
	)
	for i := range ps {
		for j := 0; j < n; j++ {
			v = ps[i][j]
			if v < 0 {
				v = -v
			}

			if maxPnt[j] < v {
				maxPnt[j] = v
			}
		}
	}

	return maxPnt
}

// shufflePoints randomly orders a set of points.
func shufflePoints(pnts Points) Points {
	seedRNG()
	rand.Shuffle(len(pnts), func(i, j int) {
		temp := pnts[i]
		pnts[i] = pnts[j]
		pnts[j] = temp
	})

	return pnts
}

// Shuffle randomly orders a set of points.
func (ps Points) Shuffle() {
	seedRNG()
	rand.Shuffle(len(ps), func(i, j int) { ps[i], ps[j] = ps[j], ps[i] })
}

// ComparePoints returns -1, 0, or 1 indicating point 0 precedes, is equal to, or follows point 1.
func ComparePoints(pnt0, pnt1 Point) int {
	for i := range pnt0 {
		if pnt0[i] < pnt1[i] {
			return -1
		}

		if pnt1[i] < pnt0[i] {
			return 1
		}
	}

	return 0
}

// CompareTo returns -1, 0, or 1 indicating point 0 precedes, is equal to, or follows point 1.
func (p Point) CompareTo(q Point) int {
	n := len(p)
	if n != len(q) {
		panic("dimension mismatch")
	}

	for i := 0; i < n; i++ {
		if p[i] < q[i] {
			return -1
		}

		if q[i] < p[i] {
			return 1
		}
	}

	return 0
}

// CopyPoint returns a copy of a point.
func CopyPoint(pnt Point) Point {
	cpy := make(Point, len(pnt))
	copy(cpy, pnt)
	return cpy
}

// Copy a point.
func (p Point) Copy() Point {
	cpy := make(Point, len(p))
	copy(cpy, p)
	return cpy
}

// CopyPoints returns a copy of a set of points.
func CopyPoints(pnts Points) Points {
	cpy := make(Points, 0, len(pnts))
	for i := range pnts {
		cpy = append(cpy, pnts[i].Copy())
	}

	return cpy
}

// Copy returns a copy of a set of points.
func (ps Points) Copy() Points {
	cpy := make(Points, 0, len(ps))
	for i := range ps {
		cpy = append(cpy, ps[i].Copy())
	}

	return cpy
}

// Sort a set of points given a sort option.
func (ps Points) Sort() {
	sort.Slice(ps, func(i, j int) bool { return ps[i].CompareTo(ps[j]) < 0 })
}

// ToCluster returns a cluster converted from a set of points.
func (ps Points) ToCluster() Cluster {
	c := make(Cluster, len(ps))
	copy(c, ps)
	return c
}
