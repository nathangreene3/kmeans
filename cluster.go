package kmeans

import (
	"math"
	"sort"
)

// Cluster is a set of points.
type Cluster Points

// CompareTo returns -1, 0, or 1 indicating cluster c precedes, is equal to, or follows another cluster.
func (c Cluster) CompareTo(cluster Cluster) int {
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

	maxIndex := min(m, n)
	var comparison int
	for i := 0; i < maxIndex; i++ {
		if comparison = c[i].CompareTo(cluster[i]); comparison != 0 {
			return comparison
		}
	}

	switch {
	case m < n:
		return -1 // c is shorter (while equal over the range [0,m))
	case n < m:
		return 1 // clstr is shorter (while equal over the range [0,n))
	default:
		return 0 // c and clster are equal in length and in each point
	}
}

// Copy a cluster.
func (c Cluster) Copy() Cluster {
	cluster := make(Cluster, 0, len(c))
	for _, p := range c {
		cluster = append(cluster, p.Copy())
	}

	return cluster
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

	var (
		mean         Point
		meanVariance = math.MaxFloat64
		variance     float64
	)

	for _, p := range c {
		if variance = c.Variance(p); variance < meanVariance {
			mean = p
			meanVariance = variance
		}
	}

	return mean
}

// Sort a cluster by a sorting option.
func (c Cluster) Sort(sortOpt SortOption) {
	switch sortOpt {
	case SortByVariance:
		if mn := c.Mean(); mn != nil {
			sort.SliceStable(c, func(i, j int) bool { return mn.SqDist(c[i]) < mn.SqDist(c[j]) })
		}
	case SortByDimension:
		sort.SliceStable(c, func(i, j int) bool { return c[i].CompareTo(c[j]) < 0 })
	}
}

// ToPoints returns a set of points from a cluster.
func (c Cluster) ToPoints() Points {
	ps := make(Points, 0, len(c))
	for _, p := range c {
		ps = append(ps, p.Copy())
	}

	return ps
}

// Variance returns the Variance of the cluster to the mean.
func (c Cluster) Variance(mean Point) float64 {
	n := len(c)
	if n < 2 {
		return 0
	}

	// The variance is the sum of the squared Euclidean distances, divided by the number of points minus one.
	var v float64
	for _, p := range c {
		v += mean.SqDist(p)
	}

	return v / float64(n-1)
}

// Transfer ith point from the source cluster to the destination cluster. Returns (dest, src).
func Transfer(i int, destination, source Cluster) (Cluster, Cluster) {
	destination = append(destination, source[i])
	destination.Sort(SortByVariance)
	if i+1 < len(source) {
		return destination, append(source[:i], source[i+1:]...)
	}

	return destination, source[:i]
}
