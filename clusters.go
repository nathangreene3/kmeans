package kmeans

import "sort"

// Clusters is a set of clusters.
type Clusters []Cluster

// Coalesce ensures clusters containing equivalent points are placed in one
// cluster. For each cluster i, the points in each cluster j are compared and
// equivalent points are moved to cluster i.
func (cs Clusters) Coalesce() {
	cs.SortAll(SortByVariance)

	var (
		m = len(cs)
		n int // Number of points in cluster i
		c int // Comparison result
	)

	for i := 0; i < m; i++ {
		n = len(cs[i])
		for j := 0; j < n; j++ {
			if j+1 < n && cs[i][j].Compare(cs[i][j+1]) == 0 {
				// Points j and j+1 are equal, so keep iterating until the last
				// equal point is found.
				continue
			}

			for k := i + 1; k < m; k++ {
				for b := 0; b < len(cs[k]); b++ {
					if c = cs[i][j].Compare(cs[k][b]); c == 0 {
						cs[i], cs[k] = Transfer(b, cs[i], cs[k])
						continue
					}

					if 0 < c {
						break
					}
				}
			}
		}

		cs[i].Sort(SortByVariance)
	}
}

// Copy returns a copy of a set of clusters.
func (cs Clusters) Copy() Clusters {
	clusters := make(Clusters, 0, len(cs))
	for _, c := range cs {
		clusters = append(clusters, c.Copy())
	}

	return clusters
}

// Join into a single, sorted cluster (sorted on variance).
func (cs Clusters) Join() Cluster {
	var n int
	for _, c := range cs {
		n += len(c)
	}

	joinedCluster := make(Cluster, 0, n)
	for _, c := range cs {
		joinedCluster = append(joinedCluster, c...)
	}

	joinedCluster.Sort(SortByVariance)
	return joinedCluster
}

// Means returns the set of points each representing the mean (center) of each
// cluster.
func (cs Clusters) Means() Points {
	means := make(Points, 0, len(cs))
	for _, c := range cs {
		means = append(means, c.Mean())
	}

	return means
}

// MeanVariance returns the mean variance of a set of clusters.
func (cs Clusters) MeanVariance() float64 {
	var v float64
	for _, c := range cs {
		v += c.Variance(c.Mean())
	}

	return v / float64(len(cs))
}

// MeanWeightedVariance returns the mean of the variances weighted by the number
// of points in each cluster.
func (cs Clusters) MeanWeightedVariance() float64 {
	var (
		n float64 // Number of points in clusters
		s float64 // Number of points in cluster i
		v float64 // Sum of size-weighted variances
	)

	for _, c := range cs {
		s = float64(len(c))
		n += s
		v += s * c.Variance(c.Mean())
	}

	return v / n
}

// Sort sorts a set of clusters.
func (cs Clusters) Sort() {
	sort.SliceStable(cs, func(i, j int) bool { return cs[i].Compare(cs[j]) < 0 })
}

// SortAll sorts each cluster. The set of clusters is NOT sorted.
func (cs Clusters) SortAll(sortOpt SortOption) {
	for _, c := range cs {
		c.Sort(sortOpt)
	}
}

// Variances returns the set of variances for each cluster.
func (cs Clusters) Variances() []float64 {
	variances := make([]float64, 0, len(cs))
	for _, c := range cs {
		variances = append(variances, c.Variance(c.Mean()))
	}

	return variances
}
