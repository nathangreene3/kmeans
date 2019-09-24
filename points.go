package kmeans

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

// Sort a set of points given a sort option.
func (ps Points) Sort() {
	sort.Slice(ps, func(i, j int) bool { return ps[i].CompareTo(ps[j]) < 0 })
}

// ToCluster returns a cluster converted from a set of points.
func (ps Points) ToCluster() Cluster {
	cluster := make(Cluster, 0, len(ps))
	for _, p := range ps {
		cluster = append(cluster, p.Copy())
	}

	return cluster
}

// validate panics if there are no points or if any points are of unequal or
// zero dimension.
func (ps Points) validate() {
	// TODO: Make real errors instead of panicing.

	if n := len(ps); 0 < n {
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
