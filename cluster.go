package kmeans

import "sort"

// Cluster is a set of points.
type Cluster Points

// CompareTo returns -1, 0, or 1 indicating cluster c precedes, is equal to, or follows another cluster.
func (c Cluster) CompareTo(clstr Cluster) int {
	m, n := len(c), len(clstr)
	switch {
	case m == 0:
		if n == 0 {
			return 0
		}
		return -1
	case n == 0:
		return 1
	}

	maxIndex := min(m, n)
	var comparison int
	for i := 0; i < maxIndex; i++ {
		if comparison = c[i].CompareTo(clstr[i]); comparison != 0 {
			return comparison
		}
	}

	switch {
	case m < n:
		// c is shorter (while equal over the range [0,m))
		return -1
	case n < m:
		// clstr is shorter (while equal over the range [0,n))
		return 1
	default:
		// c and clster are equal in length and in each point
		return 0
	}
}

// Copy a cluster.
func (c Cluster) Copy() Cluster {
	cpy := make(Cluster, 0, len(c))
	for _, p := range c {
		cpy = append(cpy, p.Copy())
	}

	return cpy
}

// Mean returns a point representing the mean (center) of the cluster.
func (c Cluster) Mean() Point {
	n := float64(len(c))
	switch n {
	case 0:
		return nil
	case 1:
		return c[0].Copy()
	}

	d := len(c[0])
	mn := make(Point, d)
	for _, clstr := range c {
		if d != len(clstr) {
			panic("dimension mismatch")
		}

		for j := 0; j < d; j++ {
			mn[j] += clstr[j]
		}
	}

	for i := range mn {
		mn[i] /= n
	}

	return mn
}

// Sort a cluster by a sorting option.
func (c Cluster) Sort(st SortOpt) {
	switch st {
	case VarSort:
		if mn := c.Mean(); mn != nil {
			sort.SliceStable(c, func(i, j int) bool { return mn.SqDist(c[i]) < mn.SqDist(c[j]) })
		}
	case LexiSort:
		sort.SliceStable(c, func(i, j int) bool { return c[i].CompareTo(c[j]) < 0 })
	}
}

// ToPoints returns a set of points from a cluster.
func (c Cluster) ToPoints() Points {
	ps := make(Points, len(c))
	copy(ps, c)
	return ps
}

// Variance returns the Variance of the cluster to the mean.
func (c Cluster) Variance() float64 {
	n := len(c)
	if n < 2 {
		return 0
	}

	// The variance is the sum of the squared Euclidean distances, divided by the number of points minus one.
	mn := c.Mean()
	var v float64
	for _, p := range c {
		v += mn.SqDist(p)
	}

	return v / float64(n-1)
}

// Transfer ith point from the source cluster to the destination cluster. Returns (dest, src).
func Transfer(i int, dest, src Cluster) (Cluster, Cluster) {
	dest = append(dest, src[i])
	dest.Sort(VarSort)
	if i+1 < len(src) {
		return dest, append(src[:i], src[i+1:]...)
	}

	return dest, src[:i]
}
