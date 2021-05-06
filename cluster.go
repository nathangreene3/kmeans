package kmeans

import (
	gomath "math"
	"sort"
)

// Cluster is a set of points.
type Cluster Points

// Compare returns -1, 0, or 1 indicating cluster c precedes, is equal to, or
// follows another cluster.
func (c Cluster) Compare(cluster Cluster) int {
	m, n := len(c), len(cluster)
	switch {
	case m == 0:
		if n == 0 {
			return 0
		}

		return -1
	case n == 0:
		return 1
	}

	var (
		maxIndex   = minInt(m, n)
		comparison int
	)

	for i := 0; i < maxIndex; i++ {
		if comparison = c[i].Compare(cluster[i]); comparison != 0 {
			return comparison
		}
	}

	switch {
	case m < n:
		return -1 // c is shorter (while equal over the range [0,m))
	case n < m:
		return 1 // clustr is shorter (while equal over the range [0,n))
	default:
		return 0 // c and clster are equal in length and in each point
	}
}

// Copy a cluster.
func (c Cluster) Copy() Cluster {
	cluster := make(Cluster, 0, len(c))
	for i := 0; i < len(c); i++ {
		cluster = append(cluster, c[i].Copy())
	}

	return cluster
}

// Mean returns a point representing the mean (center) of the cluster.
func (c Cluster) Mean() Point {
	switch len(c) {
	case 0:
		return nil
	case 1:
		return c[0].Copy()
	}

	var (
		meanVariance = gomath.MaxFloat64
		variance     float64
		mean         Point
	)

	for i := 0; i < len(c); i++ {
		if variance = c.Variance(c[i]); variance < meanVariance {
			mean = c[i].Copy()
			meanVariance = variance
		}
	}

	return mean
}

// Sort a cluster by a sorting option.
func (c Cluster) Sort(sortOpt SortOption) {
	switch sortOpt {
	case SortByVariance:
		if mean := c.Mean(); mean != nil {
			sort.SliceStable(c, func(i, j int) bool { return mean.Dist(c[i]) < mean.Dist(c[j]) })
		}
	case SortByDimension:
		sort.SliceStable(c, func(i, j int) bool { return c[i].Compare(c[j]) < 0 })
	}
}

// ToPoints returns a set of points from a cluster.
func (c Cluster) ToPoints() Points {
	points := make(Points, 0, len(c))
	for i := 0; i < len(c); i++ {
		points = append(points, c[i].Copy())
	}

	return points
}

// Variance returns the Variance of the cluster to the mean.
func (c Cluster) Variance(mean Point) float64 {
	n := len(c)
	if n < 2 {
		return 0
	}

	// The variance is the sum of the squared Euclidean distances, divided by
	// the number of points minus one.
	var v float64
	for i := 0; i < len(c); i++ {
		v += mean.Dist(c[i]) // That's not squared... unless the distance function is squared.  That's defined by the user, though.
	}

	return v / float64(n-1)
}

// Transfer ith point from the source cluster to the destination cluster.
// Returns (dest, src).
func Transfer(i int, destination, source Cluster) (Cluster, Cluster) {
	destination = append(destination, source[i])
	destination.Sort(SortByVariance)
	if i+1 < len(source) {
		return destination, append(source[:i], source[i+1:]...)
	}

	return destination, source[:i]
}
