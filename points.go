package kmeans

import (
	"math/rand"
	"sort"
)

// Points is a set of points.
type Points []Point

// Copy returns a copy of a set of points.
func (ps Points) Copy() Points {
	cpy := make(Points, 0, len(ps))
	for i := range ps {
		cpy = append(cpy, ps[i].Copy())
	}

	return cpy
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
	c := make(Cluster, 0, len(ps))
	for _, p := range ps {
		c = append(c, p.Copy())
	}

	return c
}

// validate panics if there are no points or if any points are of unequal or zero dimension.
func (ps Points) validate() {
	// TODO: Make real errors instead of panicing.

	if n := len(ps); 0 < n {
		d := ps[0].Len()
		if d == 0 {
			panic("dimensionless point")
		}

		for i := 1; i < n; i++ {
			if d != ps[i].Len() {
				panic("dimension mismatch")
			}
		}
	}
}
