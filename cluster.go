package main

import "sort"

// Cluster is a set of points.
type Cluster Points

// Clusters is a set of clusters.
type Clusters []Cluster

// initClusters returns a set of k-sorted clusters.
func initClusters(k int, ps Points, normalize bool) Clusters {
	ps.validate()
	ps.Shuffle()
	if normalize {
		ps.Normalize()
	}

	// Each cluster contains capacity/k points. Remainders will be added to the last cluster (index k-1).
	var (
		clusters = make(Clusters, 0, k) // Clusters to return
		capacity = len(ps)              // Cluster capacity
		length   = capacity / k         // Cluster length, not including remainder
		h        int                    // Indexer through points
	)
	for i := 0; i < k; i++ {
		clstr := make(Cluster, 0, capacity)
		for j := 0; j < length; j++ {
			clstr = append(clstr, ps[h])
			h++
		}

		clusters = append(clusters, clstr)
	}

	// The last cluster is potentially largest if k doesn't divide n evenly. The actual size is s + n mod k. For example, given 32 points on 5 clusters, the size would be 32/5 + 32 mod 5 = 8.
	for ; h < capacity; h++ {
		clusters[k-1] = append(clusters[k-1], ps[h])
	}

	clusters.Coalesce() // Sorts all
	return clusters
}

// Mean returns a point representing the mean (center) of the cluster.
func Mean(c Cluster) Point {
	n := float64(len(c))
	switch n {
	case 0:
		return nil
	case 1:
		return c[0].Copy()
	}

	d := len(c[0])
	mn := make(Point, d)
	for i := range c {
		for j := range c[i] {
			mn[j] += c[i][j]
		}
	}

	for i := range mn {
		mn[i] /= n
	}

	return mn
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
	for i := range c {
		for j := range c[i] {
			mn[j] += c[i][j]
		}
	}

	for i := range mn {
		mn[i] /= n
	}

	return mn
}

// Means returns the set of points each representing the mean (center) of each cluster.
func Means(clstrs Clusters) Points {
	mns := make(Points, 0, len(clstrs))
	for i := range clstrs {
		mns = append(mns, Mean(clstrs[i]))
	}

	return mns
}

// Means returns the set of points each representing the mean (center) of each cluster.
func (cs Clusters) Means() Points {
	mns := make(Points, 0, len(cs))
	for i := range cs {
		mns = append(mns, cs[i].Mean())
	}

	return mns
}

// Variance returns the Variance of the cluster to the mean.
func Variance(clstr Cluster) float64 {
	n := len(clstr)
	if n < 2 {
		return 0
	}

	// The variance is the sum of the squared Euclidean distances, divided by the number of points minus one.
	mn := Mean(clstr) // Cluster mean
	var v float64     // Variance to return
	for i := range clstr {
		v += SquaredDistance(mn, clstr[i])
	}

	return v / float64(n-1)
}

// Variance returns the Variance of the cluster to the mean.
func (c Cluster) Variance() float64 {
	n := len(c)
	if n < 2 {
		return 0
	}

	// The variance is the sum of the squared Euclidean distances, divided by the number of points minus one.
	mn := c.Mean() // Cluster mean
	var v float64  // Variance to return
	for i := range c {
		v += SquaredDistance(mn, c[i])
	}

	return v / float64(n-1)
}

// Variances returns the set of variances for each cluster.
func Variances(cs Clusters) []float64 {
	vars := make([]float64, 0, len(cs))
	for i := range cs {
		vars = append(vars, cs[i].Variance())
	}

	return vars
}

// Variances returns the set of variances for each cluster.
func (cs Clusters) Variances() []float64 {
	vars := make([]float64, 0, len(cs))
	for i := range cs {
		vars = append(vars, cs[i].Variance())
	}

	return vars
}

// MeanVariance returns the mean variance of a set of clusters.
func MeanVariance(clstrs Clusters) float64 {
	var v float64 // Sum of variances
	for i := range clstrs {
		v += Variance(clstrs[i])
	}

	return v / float64(len(clstrs))
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
func MeanWeightedVariance(clstrs Clusters) float64 {
	var (
		n float64 // Number of points in clusters
		s float64 // Number of points in cluster i
		v float64 // Sum of size-weighted variances
	)
	for i := range clstrs {
		s = float64(len(clstrs[i]))
		n += s
		v += s * Variance(clstrs[i])
	}

	return v / n
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

// SortCluster sorts a cluster on the squared distance.
func SortCluster(clstr Cluster, st SortOpt) Cluster {
	switch st {
	case VarSort:
		if mn := Mean(clstr); mn != nil {
			sort.SliceStable(clstr, func(i, j int) bool { return SquaredDistance(mn, clstr[i]) < SquaredDistance(mn, clstr[j]) })
		}
	case LexiSort:
		sort.SliceStable(clstr, func(i, j int) bool { return ComparePoints(clstr[i], clstr[j]) < 0 })
	}

	return clstr
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

// SortClusters sorts a set of clusters.
func SortClusters(cs Clusters) {
	sort.SliceStable(cs, func(i, j int) bool { return CompareClusters(cs[i], cs[j]) < 0 })
}

// Sort sorts a set of clusters.
func (cs Clusters) Sort() {
	sort.SliceStable(cs, func(i, j int) bool { return cs[i].CompareTo(cs[j]) < 0 })
}

// SortAllClusters sorts each cluster. The set of clusters is NOT sorted.
func SortAllClusters(clstrs Clusters, st SortOpt) Clusters {
	for i := range clstrs {
		SortCluster(clstrs[i], st)
	}

	return clstrs
}

// SortAll sorts each cluster. The set of clusters is NOT sorted.
func (cs Clusters) SortAll(st SortOpt) {
	for i := range cs {
		cs[i].Sort(st)
	}
}

// CompareClusters returns -1, 0, or 1 indicating cluster 0 precedes, is equal to, or follows cluster 1.
func CompareClusters(clstr0, clstr1 Cluster) int {
	m, n := len(clstr0), len(clstr1)
	if m == 0 {
		if n == 0 {
			return 0
		}

		return -1
	}

	if n == 0 {
		return 1
	}

	maxIndex := min(m, n)
	var comparison int
	for i := 0; i < maxIndex; i++ {
		comparison = ComparePoints(clstr0[i], clstr1[i])
		if comparison != 0 {
			return comparison
		}
	}

	if m < n {
		// Point 0 is shorter than point 1 (while equal over the range [0,m))
		return -1
	}

	if n < m {
		// Point 1 is shorter than point 0 (while equal over the range [0,n))
		return 1
	}

	// Points are equal in length and in each value
	return 0
}

// CompareTo returns -1, 0, or 1 indicating cluster c precedes, is equal to, or follows another cluster.
func (c Cluster) CompareTo(clstr Cluster) int {
	m, n := len(c), len(clstr)
	if m == 0 {
		if n == 0 {
			return 0
		}
		return -1
	}

	if n == 0 {
		return 1
	}

	maxIndex := min(m, n)
	var comparison int
	for i := 0; i < maxIndex; i++ {
		if comparison = c[i].CompareTo(clstr[i]); comparison != 0 {
			return comparison
		}
	}

	if m < n {
		// Point 0 is shorter than point 1 (while equal over the range [0,m))
		return -1
	}

	if n < m {
		// Point 1 is shorter than point 0 (while equal over the range [0,n))
		return 1
	}

	// Points are equal in length and in each value
	return 0
}

// CopyCluster returns a copy of a cluster.
func CopyCluster(clstr Cluster) Cluster {
	cpy := make(Cluster, 0, len(clstr))
	for i := range clstr {
		cpy = append(cpy, CopyPoint(clstr[i]))
	}

	return cpy
}

// Copy a cluster.
func (c Cluster) Copy() Cluster {
	cpy := make(Cluster, 0, len(c))
	for i := range c {
		cpy = append(cpy, c[i].Copy())
	}

	return cpy
}

// CopyClusters returns a copy of a set of clusters.
func CopyClusters(clstrs Clusters) Clusters {
	cpy := make(Clusters, 0, len(clstrs))
	for i := range clstrs {
		cpy = append(cpy, CopyCluster(clstrs[i]))
	}

	return cpy
}

// Copy returns a copy of a set of clusters.
func (cs Clusters) Copy() Clusters {
	cpy := make(Clusters, 0, len(cs))
	for i := range cs {
		cpy = append(cpy, cs[i].Copy())
	}

	return cpy
}

// coalesceClusters ensures clusters containing equivalent points are placed in one cluster. For each cluster i, the points in each cluster j are compared and equivalent points are moved to cluster i.
func coalesceClusters(clstrs Clusters) Clusters {
	clstrs = SortAllClusters(clstrs, VarSort)

	var (
		numClstrs  = len(clstrs) // Number of clusters
		numPntsI   int           // Number of points in cluster i
		comparison int           // Comparison result
	)
	for i := 0; i < numClstrs; i++ {
		numPntsI = len(clstrs[i])
		for j := 0; j < numPntsI; j++ {
			if j+1 < numPntsI && ComparePoints(clstrs[i][j], clstrs[i][j+1]) == 0 {
				// Points j and j+1 are equal, so keep iterating until the last equal point is found.
				continue
			}

			for k := i + 1; k < numClstrs; k++ {
				for b := 0; b < len(clstrs[k]); b++ {
					comparison = ComparePoints(clstrs[i][j], clstrs[k][b])
					if comparison == 0 {
						clstrs[i], clstrs[k] = Transfer(b, clstrs[i], clstrs[k])
						continue
					}

					if 0 < comparison {
						break
					}
				}
			}
		}

		clstrs[i] = SortCluster(clstrs[i], VarSort)
	}

	return clstrs
}

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

// joinClusters into a single, sorted cluster (sorted on variance).
func joinClusters(clstrs Clusters) Cluster {
	var n int
	for i := range clstrs {
		n += len(clstrs[i])
	}

	clstr := make(Cluster, 0, n)
	for i := range clstrs {
		clstr = append(clstr, clstrs[i]...)
	}

	return SortCluster(clstr, VarSort)
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

// ToPoints returns a set of points from a cluster.
func (c Cluster) ToPoints() Points {
	ps := make(Points, len(c))
	copy(ps, c)
	return ps
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
