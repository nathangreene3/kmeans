package kmeans2

import "sort"

// Clusters ...
type Clusters []Cluster

// Coalesce ensures clusters containing equivalent points are placed in one
// cluster. For each cluster i, the points in each cluster j are compared and
// equivalent points are moved to cluster i.
func (cs Clusters) Coalesce() {
	cs.SortAll(SortByVariance, false)

	m := len(cs)
	for i := 0; i < m; i++ {
		n := cs[i].size
		for j := 0; j < n; j++ {
			if j+1 < n && cs[i].points[j].Compare(cs[i].points[j+1]) == 0 {
				// Points j and j+1 are equal, so keep iterating until the last
				// equal point is found.
				continue
			}

			for k := i + 1; k < m; k++ {
				for b := 0; b < cs[k].size; b++ {
					c := cs[i].points[j].Compare(cs[k].points[b])
					if c == 0 {
						cs[k].transfer(b, &cs[i])
					} else if 0 < c {
						break
					}
				}
			}
		}

		cs[i].Sort(SortByVariance, false)
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
	var joinedCluster Cluster
	for _, c := range cs {
		joinedCluster.points = append(joinedCluster.points, c.points...)
	}

	joinedCluster.Sort(SortByVariance, false)
	joinedCluster.Update()
	return joinedCluster
}

// Means returns the set of points each representing the mean (center) of each
// cluster.
func (cs Clusters) Means() Points {
	means := make(Points, 0, len(cs))
	for _, c := range cs {
		means = append(means, c.mean)
	}

	return means
}

// MeanVariance returns the mean variance of a set of clusters.
func (cs Clusters) MeanVariance() float64 {
	var v float64
	for _, c := range cs {
		v += c.variance
	}

	return v / float64(len(cs))
}

// MeanWeightedVariance returns the mean of the variances weighted by the number
// of points in each cluster.
func (cs Clusters) MeanWeightedVariance() float64 {
	var n, v float64
	for _, c := range cs {
		s := float64(c.size)
		n += s
		v += s * c.variance
	}

	return v / n
}

// Sort sorts a set of clusters.
func (cs Clusters) Sort(stable bool) {
	if stable {
		sort.SliceStable(cs, func(i, j int) bool { return cs[i].Compare(cs[j]) < 0 })
	} else {
		sort.Slice(cs, func(i, j int) bool { return cs[i].Compare(cs[j]) < 0 })
	}
}

// SortAll sorts each cluster. The set of clusters is NOT sorted.
func (cs Clusters) SortAll(sortOpt SortOption, stable bool) {
	for i := range cs {
		cs[i].Sort(sortOpt, stable)
	}
}

// Variances returns the set of variances for each cluster.
func (cs Clusters) Variances() []float64 {
	variances := make([]float64, 0, len(cs))
	for _, c := range cs {
		variances = append(variances, c.variance)
	}

	return variances
}
