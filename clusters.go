package main

import "sort"

// Clusters is a set of clusters.
type Clusters []Cluster

// Coalesce ensures clusters containing equivalent points are placed in one cluster. For each cluster i, the points in each cluster j are compared and equivalent points are moved to cluster i.
func (cs Clusters) Coalesce() {
	cs.SortAll(VarSort)

	var (
		numClstrs  = len(cs) // Number of clusters
		numPntsI   int       // Number of points in cluster i
		comparison int       // Comparison result
	)
	for i := 0; i < numClstrs; i++ {
		numPntsI = len(cs[i])
		for j := 0; j < numPntsI; j++ {
			if j+1 < numPntsI && cs[i][j].CompareTo(cs[i][j+1]) == 0 {
				// Points j and j+1 are equal, so keep iterating until the last equal point is found.
				continue
			}

			for k := i + 1; k < numClstrs; k++ {
				for b := 0; b < len(cs[k]); b++ {
					if comparison = cs[i][j].CompareTo(cs[k][b]); comparison == 0 {
						cs[i], cs[k] = Transfer(b, cs[i], cs[k])
						continue
					}

					if 0 < comparison {
						break
					}
				}
			}
		}

		cs[i].Sort(VarSort)
	}
}

// Copy returns a copy of a set of clusters.
func (cs Clusters) Copy() Clusters {
	cpy := make(Clusters, 0, len(cs))
	for i := range cs {
		cpy = append(cpy, cs[i].Copy())
	}

	return cpy
}

// Join into a single, sorted cluster (sorted on variance).
func (cs Clusters) Join() Cluster {
	var i, n int
	for i := range cs {
		n += len(cs[i])
	}

	c := make(Cluster, n)
	for _, cluster := range cs {
		n = len(cluster)
		copy(c[i:i+n], cluster)
		i += n
	}

	c.Sort(VarSort)
	return c
}

// Means returns the set of points each representing the mean (center) of each cluster.
func (cs Clusters) Means() Points {
	mns := make(Points, 0, len(cs))
	for i := range cs {
		mns = append(mns, cs[i].Mean())
	}

	return mns
}

// MeanVariance returns the mean variance of a set of clusters.
func (cs Clusters) MeanVariance() float64 {
	var v float64 // Sum of variances
	for i := range cs {
		v += cs[i].Variance()
	}

	return v / float64(len(cs))
}

// MeanWeightedVariance returns the mean of the variances weighted by the number of points in each cluster.
func (cs Clusters) MeanWeightedVariance() float64 {
	var (
		n float64 // Number of points in clusters
		s float64 // Number of points in cluster i
		v float64 // Sum of size-weighted variances
	)
	for i := range cs {
		s = float64(len(cs[i]))
		n += s
		v += s * cs[i].Variance()
	}

	return v / n
}

// Sort sorts a set of clusters.
func (cs Clusters) Sort() {
	sort.SliceStable(cs, func(i, j int) bool { return cs[i].CompareTo(cs[j]) < 0 })
}

// SortAll sorts each cluster. The set of clusters is NOT sorted.
func (cs Clusters) SortAll(st SortOpt) {
	for i := range cs {
		cs[i].Sort(st)
	}
}

// Variances returns the set of variances for each cluster.
func (cs Clusters) Variances() []float64 {
	vars := make([]float64, 0, len(cs))
	for i := range cs {
		vars = append(vars, cs[i].Variance())
	}

	return vars
}
