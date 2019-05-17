package main

import (
	"math"
	"math/rand"
	"sort"
)

// Point is an n-dimensional point in n-space.
type Point []float64

// Cluster is a set of points.
type Cluster []Point

// SortOpt indicates how a cluster is sorted.
type SortOpt int

const (
	// VarSort dictates points will be compared by variance to the cluster mean.
	VarSort SortOpt = iota
	// LexiSort dictates points will be compared by the default comparer, which is lexicographic.
	LexiSort
)

// assignPoint returns the index of the closest point (a cluster mean) to a given point.
func assignPoint(pnt Point, mns []Point) int {
	var assignment int
	var (
		sd    float64
		minSD = math.MaxFloat64
	)
	for i := range mns {
		sd = SquaredDistance(mns[i], pnt)
		if sd < minSD {
			minSD = sd
			assignment = i
		}
	}

	return assignment
}

// KMeans clusters a set of points into k groups. Potentially, clusters can be empty, so multiple attempts should be made.
func KMeans(k int, pnts []Point, normalize bool) []Cluster {
	// Move points to their nearest cluster until they no longer move with each pass (indicated by the changed boolean).
	var (
		clstrs    = initClusters(k, pnts, normalize) // Clusters to return
		changed   = true                             // Indicates if a cluster was altered
		mns       []Point                            // Means of clusters
		minIndex  int                                // Index of cluster having smallest squared distance to a point
		sd, minSD float64                            // Squared distance; minimum squared distance
	)

	for changed {
		changed = false

		// Update the means. If any cluster is empty, its mean, which is nil, will be reassigned to be a random point on the space spanned by the maximum point.
		mns = Means(clstrs)
		for i := range mns {
			if mns[i] == nil {
				randPnt(pnts)
			}
		}

		for h := range clstrs {
			// Each cluster is sorted, so the most variant point is at index clstrs[h][sizes[h]-1]. When the point clstrs[h][i] variance is small enough to not move to another cluster, then we can move on to the next cluster and ignore the other points on the range [0,i). So, h counts down and halts early.
			for i := 0; i < len(clstrs[h]); i++ {
				// Find the index of the cluster closest to point i in cluster h.
				minIndex = h
				minSD = SquaredDistance(mns[h], clstrs[h][i])
				for j := range clstrs {
					if h == j {
						// If h = j, then we are comparing the same cluster. If the size of cluster j is zero, then, obviously, there's no points to compare.
						continue
					}

					sd = SquaredDistance(mns[j], clstrs[h][i])
					if sd < minSD {
						minSD = sd
						minIndex = j
					}
				}

				if h == minIndex {
					// Point i in cluster h is closest to cluster h, so we don't need to move it.
					continue
				}

				// Move point i in cluster h to the nearest cluster and update sizes and changed.
				clstrs[minIndex], clstrs[h] = movePoint(i, clstrs[minIndex], clstrs[h])
				changed = true
			}
		}
	}

	return clstrs
}

// OptimalKMeans determines the optimal number of clusters and returns the clustering result.
func OptimalKMeans(pnts []Point, normalize bool) (int, []Cluster) {
	var (
		n        = len(pnts) // Number of points
		vars     []float64
		varsPnts []Point // Vars converted to points
	)

	if n <= 10 {
		// Run k-means on k = 1..n and track the mean variance.
		varsPnts = make([]Point, 0, n)
		for numClstrs := 1; numClstrs <= n; numClstrs++ {
			vars = append(vars, MeanVariance(KMeans(numClstrs, pnts, normalize)))
			varsPnts = append(varsPnts, Point{vars[numClstrs-1]})
		}

		// Run 2-means on the variances to find the elbow point at which smaller numbers of clusters increases the mean variance.
		var (
			k          int                              // Optimal k
			minIndex   int                              // Index of cluster having the smallest mean variance
			v          float64                          // Variance
			maxVars    float64                          // Maximum variance
			varsClstrs = KMeans(2, varsPnts, normalize) // Variance clusters run on k = 2 to find the elbow point k
		)

		if 0 < ComparePoints(Mean(varsClstrs[0]), Mean(varsClstrs[1])) {
			// Mean var of cluster 0 > mean var of cluster 1
			minIndex = 1
		}

		// Find the index of the largest value in the cluster with the smaller mean variance.
		for i := range varsClstrs[minIndex] {
			v = varsClstrs[minIndex][i][0]
			if maxVars < v {
				maxVars = v
				k = i + 1
			}
		}

		return k, KMeans(k, pnts, normalize)
	}

	return 0, nil
}

// validate panics if there are no points or if any points are of unequal or zero dimension.
func validate(pnts []Point) {
	numPnts := len(pnts)
	if numPnts == 0 {
		panic("validate: no points")
	}

	dims := len(pnts[0])
	if dims == 0 {
		panic("validate: dimensionless point")
	}

	for i := 1; i < numPnts; i++ {
		if dims != len(pnts[i]) {
			panic("validate: dimension mismatch")
		}
	}
}

// initClusters returns a set of k sorted clusters.
func initClusters(k int, pnts []Point, normalize bool) []Cluster {
	validate(pnts)
	pnts = shufflePoints(pnts)
	if normalize {
		pnts = NormalizePoints(pnts)
	}

	// Each cluster contains n/k points. Remainders will be added to the last cluster (index k-1).
	var (
		clstrs = make([]Cluster, 0, k) // Clusters to return
		n      = len(pnts)             // Cluster capacity
		s      = n / k                 // Cluster length, not including remainder
		h      int                     // Indexer through points
	)

	for i := 0; i < k; i++ {
		clstr := make(Cluster, 0, n)
		for j := 0; j < s; j++ {
			clstr = append(clstr, pnts[h])
			h++
		}

		SortCluster(clstr, VarSort)
		clstrs = append(clstrs, clstr)
	}

	// The last cluster is potentially largest if k doesn't divide n evenly. The actual size is s + n mod k. For example, given 32 points on 5 clusters, the size would be 32/5 + 32 mod 5 = 8.
	for ; h < n; h++ {
		clstrs[k-1] = append(clstrs[k-1], pnts[h])
	}

	SortCluster(clstrs[k-1], VarSort)
	return coalesceClusters(clstrs)
}

// NormalizePoints each dimension in each point to the range [-1,1] assuming the largest value for each dimension is within the scope of the points provided.
func NormalizePoints(pnts []Point) []Point {
	maxPnt := maxPoint(pnts)

	// Check if max point is normal. If it is, the points are already normalized.
	var notNormal bool
	for i := range maxPnt {
		if 1 < maxPnt[i] {
			notNormal = true
			break
		}
	}

	if notNormal {
		for i := range pnts {
			for j, v := range maxPnt {
				if v != 0 {
					pnts[i][j] /= v
				}
			}
		}
	}

	return pnts
}

// coalesceClusters ensures clusters containing equivalent points are placed in one cluster.
func coalesceClusters(clstrs []Cluster) []Cluster {
	clstrs = SortAllClusters(clstrs, VarSort)

	var (
		numClstrs  = len(clstrs)
		maxIJ      int
		comparison int
	)

	for i := 0; i < numClstrs; i++ {
		maxIJ = len(clstrs[i])
		for j := 0; j < maxIJ; j++ {
			if j+1 < maxIJ && ComparePoints(clstrs[i][j], clstrs[i][j+1]) == 0 {
				continue
			}

			for k := i + 1; k < numClstrs; k++ {
				for b := 0; b < len(clstrs[k]); b++ {
					comparison = ComparePoints(clstrs[i][j], clstrs[k][b])
					if comparison == 0 {
						clstrs[i], clstrs[k] = movePoint(b, clstrs[i], clstrs[k])
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

// randPnt returns a random point in the space spanned by the maximum point on a set of points. Each dimension is a random value on (-r,r) where r is the maximum value in each dimension on the set of points. It does NOT return a point in the set.
func randPnt(pnts []Point) Point {
	seed()
	pnt := maxPoint(pnts)
	for i := range pnt {
		pnt[i] *= 2*rand.Float64() - 1
	}

	return pnt
}

// maxPoint returns a point in which each dimension is the largest non-negative value observed in the set of points.
func maxPoint(pnts []Point) Point {
	if len(pnts) == 0 {
		return nil
	}

	maxPnt := make(Point, len(pnts[0]))
	for i := range pnts {
		for j, v := range pnts[i] {
			if v < 0 {
				v = -v
			}

			if maxPnt[j] < v {
				maxPnt[j] = v
			}
		}
	}

	return maxPnt
}

// shufflePoints randomly orders a set of points.
func shufflePoints(pnts []Point) []Point {
	seed()
	rand.Shuffle(
		len(pnts),
		func(i, j int) {
			temp := pnts[i]
			pnts[i] = pnts[j]
			pnts[j] = temp
		},
	)

	return pnts
}

// movePoint i from the source cluster to the destination cluster.
func movePoint(i int, destClstr, srcClstr Cluster) (Cluster, Cluster) {
	destClstr = append(destClstr, srcClstr[i])
	if i+1 < len(srcClstr) {
		return SortCluster(destClstr, VarSort), append(srcClstr[:i], srcClstr[i+1:]...)
	}

	return SortCluster(destClstr, VarSort), srcClstr[:i]
}

// Means returns the set of points each representing the mean (center) of each cluster.
func Means(clstrs []Cluster) []Point {
	mns := make([]Point, 0, len(clstrs))
	for i := range clstrs {
		mns = append(mns, Mean(clstrs[i]))
	}

	return mns
}

// Mean returns a point representing the mean (center) of the cluster.
func Mean(clstr Cluster) Point {
	n := float64(len(clstr))
	if n == 0 {
		return nil
	}

	d := len(clstr[0])
	var mn Point
	if n == 1 {
		mn = make(Point, d)
		copy(mn, clstr[0])
		return mn
	}

	mn = make(Point, d)
	for i := range clstr {
		for j := range clstr[i] {
			mn[j] += clstr[i][j]
		}
	}

	for i := range mn {
		mn[i] /= n
	}

	return mn
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

// Variances returns the set of variances for each cluster.
func Variances(clstrs []Cluster) []float64 {
	vars := make([]float64, 0, len(clstrs))
	for i := range clstrs {
		vars = append(vars, Variance(clstrs[i]))
	}

	return vars
}

// MeanVariance returns the mean variance of a set of clusters.
func MeanVariance(clstrs []Cluster) float64 {
	var v float64 // Sum of variances
	for i := range clstrs {
		v += Variance(clstrs[i])
	}

	return v / float64(len(clstrs))
}

// SquaredDistance returns the Euclidean metric on two points.
func SquaredDistance(pnt0, pnt1 Point) float64 {
	var (
		sd float64 // Squared distance
		d  float64 // Difference in each dimension
	)
	for i := range pnt0 {
		d = pnt0[i] - pnt1[i]
		sd += d * d
	}

	return sd
}

// Distance returns the Euclidean Distance between two points.
func Distance(pnt0, pnt1 Point) float64 {
	return math.Sqrt(SquaredDistance(pnt0, pnt1))
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

// SortClusters sorts a set of clusters.
func SortClusters(clstrs []Cluster) {
	sort.SliceStable(clstrs, func(i, j int) bool { return CompareClusters(clstrs[i], clstrs[j]) < 0 })
}

// SortAllClusters sorts each cluster. The set of clusters is NOT sorted.
func SortAllClusters(clstrs []Cluster, st SortOpt) []Cluster {
	for i := range clstrs {
		SortCluster(clstrs[i], st)
	}

	return clstrs
}

// ComparePoints returns -1, 0, or 1 indicating point 0 precedes, is equal to, or follows point 1.
func ComparePoints(pnt0, pnt1 Point) int {
	for i := range pnt0 {
		if pnt0[i] < pnt1[i] {
			return -1
		}

		if pnt1[i] < pnt0[i] {
			return 1
		}
	}

	return 0
}

// CompareClusters returns -1, 0, or 1 indicating cluster 0 precedes, is equal to, or follows cluster 1.
func CompareClusters(clstr0, clstr1 Cluster) int {
	m, n := len(clstr0), len(clstr1)
	if m == 0 {
		if n == 0 {
			return 0 // Nothing to compare
		}

		return -1 // Point 0 is empty
	}

	if n == 0 {
		return 1 // Point 1 is empty
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
		return -1 // Point 0 is shorter than point 1 (while equal over the range [0,m))
	}

	if n < m {
		return 1 // Point 1 is shorter than point 0 (while equal over the range [0,n))
	}

	return 0 // Points are equal in length and in each value
}
