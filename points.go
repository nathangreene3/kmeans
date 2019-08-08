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

// MaxRep returns a point in which each dimension is the largest non-negative value observed in the set of points.
func (ps Points) MaxRep() Point {
	if len(ps) == 0 {
		return nil
	}

	var (
		v      float64
		n      = len(ps[0])
		maxPnt = make(Point, n)
	)
	for _, p := range ps {
		if n != len(p) {
			panic("dimension mismatch")
		}

		for j := 0; j < n; j++ {
			if v = p[j]; v < 0 {
				v = -v
			}

			if maxPnt[j] < v {
				maxPnt[j] = v
			}
		}
	}

	return maxPnt
}

// Normalize each dimension in each point to the range [-1,1] assuming the largest value for each dimension is within the scope of the points provided.
func (ps Points) Normalize() {
	// Check if max point is normal. If it is, the points are already normalized.
	maxPnt := ps.MaxRep()
	for _, v := range maxPnt {
		if 1 < v {
			for _, p := range ps {
				p.Normalize(maxPnt)
			}
			break
		}
	}
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
		d := len(ps[0])
		if d == 0 {
			panic("dimensionless point")
		}

		for i := 1; i < n; i++ {
			if d != len(ps[i]) {
				panic("dimension mismatch")
			}
		}
	}
}
