package kmeans2

import (
	"math/rand"
	"sort"
)

// Points is a set of points.
type Points []Point

// Copy returns a copy of a set of points.
func (ps Points) Copy() Points {
	points := make(Points, 0, len(ps))
	for _, p := range ps {
		points = append(points, p.Copy())
	}

	return points
}

// Random returns a Random point from a set of points.
func (ps Points) Random() Point {
	seedRNG()
	return ps[rand.Intn(len(ps))]
}

// Shuffle randomly orders a set of points.
func (ps Points) Shuffle() {
	seedRNG()
	rand.Shuffle(len(ps), func(i, j int) { ps[i], ps[j] = ps[j], ps[i] })
}

// Sort a set of points.
func (ps Points) Sort(stable bool) {
	if stable {
		sort.SliceStable(ps, func(i, j int) bool { return ps[i].Compare(ps[j]) < 0 })
	} else {
		sort.Slice(ps, func(i, j int) bool { return ps[i].Compare(ps[j]) < 0 })
	}
}

// validate panics if there are no points or if any points are of unequal or
// zero dimension.
func (ps Points) validate() {
	// TODO: Make real errors instead of panicing.

	if 0 < len(ps) {
		d := ps[0].Len()
		if d == 0 {
			panic("dimensionless point")
		}

		for _, p := range ps[1:] {
			if d != p.Len() {
				panic("dimension mismatch")
			}
		}
	}
}

// Variance returns the mean distance of all points to a given point.
func (ps Points) Variance(point Point) float64 {
	var (
		n = len(ps)
		v float64
	)

	if n == 0 {
		return v
	}

	for _, p := range ps {
		v += point.Dist(p)
	}

	return v / float64(n)
}
