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

// KMeans clusters a set of points into k groups. Potentially, clusters can be empty, so multiple attempts should be made.
func KMeans(k int, pnts []Point, normalize bool) []Cluster {
	// Move points to their nearest cluster until they no longer move with each pass (indicated by the changed boolean).
	var (
		clstrs    = initClusters(k, pnts, normalize) // Clusters to return
		sizes     = make([]int, k)                   // Sizes of clusters; to be updated with each pass in outer-most for loop on changed
		changed   = true                             // Indicates if a cluster was altered
		mns       []Point                            // Means of clusters
		minIndex  int                                // Index of cluster having smallest squared distance to a point
		n         int                                // Number of points
		sd, minSD float64                            // Squared distance; minimum squared distance
	)

	for h := range sizes {
		sizes[h] = len(clstrs[h])
	}

	for changed {
		changed = false
		mns = Means(clstrs)

		for h := range clstrs {
			// Each cluster is sorted, so the most variant point is at index clstrs[h][sizes[h]-1]. When the point clstrs[h][i] variance is small enough to not move to another cluster, then we can move on to the next cluster and ignore the other points on the range [0,i). So, h counts down and halts early.
			for i := sizes[h] - 1; 0 <= i && 0 < sizes[h]; i-- {
				// Find the index of the cluster closest to point i in cluster h.
				minIndex = h
				minSD = SquaredDistance(mns[h], clstrs[h][i])
				for j := range clstrs {
					if h == j { // if h == j || sizes[j] == 0 { // Experimental. See below comment.
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
				clstrs[minIndex] = append(clstrs[minIndex], clstrs[h][i])
				if i < sizes[h]-1 {
					clstrs[h] = append(clstrs[h][:i], clstrs[h][i+1:]...)
				} else {
					clstrs[h] = clstrs[h][:i]
				}

				sizes[minIndex]++
				sizes[h]--
				changed = true
			}

			// Experimental: Choose a random point as a new mean. This point won't be in the jth cluster, but its value will hopefully reset the empty cluster.
			if sizes[h] == 0 {
				seed()
				copy(mns[h], pnts[int(rand.Intn(n))])
			}
		}

		if changed {
			SortAllClusters(clstrs)
		}
	}

	SortClusters(clstrs)
	return clstrs
}

// OptimalKMeans determines the optimal number of clusters and returns the clustering result.
func OptimalKMeans(pnts []Point, normalize bool) (int, []Cluster) {
	var (
		n        = len(pnts) // Number of points
		vars     []float64   // Mean variances for each run
		varsPnts []Point     // Vars converted to points
		clstrs   []Cluster   // Clusters returned with each run
	)
	if n <= 10 {
		// Run and track the mean variance.
		vars = make([]float64, 0, n)
		varsPnts = make([]Point, 0, n)
		for numClstrs := 1; numClstrs < n; numClstrs++ {
			vars = append(vars, MeanVariances(KMeans(numClstrs, pnts, normalize)))
			varsPnts = append(varsPnts, Point{vars[numClstrs-1]})
		}

		// Run 2-means on the variances to find the elbow point at which smaller numbers of clusters increases the mean variance.
		varsClstrs := KMeans(2, varsPnts, normalize)

		// Find the maximum variance in second cluster, which contains variances on the right of the elbow.
		var v, maxVars float64 // Variance; maximum variance
		for _, varPnt := range varsClstrs[1] {
			v = varPnt[0]
			if maxVars < v {
				maxVars = v
			}
		}

		// Find the index of the maximum variance.
		var k int // Optimal k
		for i := range vars {
			if vars[i] == maxVars {
				k = i + 1 // Index is one less than the cluster count
				break
			}
		}

		// Return the optimal clustering solution.
		clstrs = KMeans(k, pnts, normalize)
		return k, clstrs
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

	if normalize {
		Normalize(pnts)
	}

	// Each cluster contains n/k points. Remainders will be added to the last cluster (index k-1).
	var (
		clstrs = make([]Cluster, 0, k) // Clusters to return
		n      = len(pnts)             // Cluster capacity
		s      = n / k                 // Cluster length, not including remainder
		h      int                     // Indexer through points
	)

	seed()
	rand.Shuffle(
		n,
		func(i, j int) {
			temp := pnts[i]
			pnts[i] = pnts[j]
			pnts[j] = temp
		},
	)

	for i := 0; i < k; i++ {
		clstr := make(Cluster, 0, n)
		for j := 0; j < s; j++ {
			clstr = append(clstr, pnts[h])
			h++
		}

		SortCluster(clstr)
		clstrs = append(clstrs, clstr)
	}

	// The last cluster is potentially largest if k doesn't divide n evenly. The actual size is s + n mod k. For example, given 32 points on 5 clusters, the size would be 32/5 + 32 mod 5 = 8.
	for ; h < n; h++ {
		clstrs[k-1] = append(clstrs[k-1], pnts[h])
	}

	SortCluster(clstrs[k-1])
	return clstrs
}

// Normalize each dimension in each point to the range [-1,1] assuming the largest value for each dimension is within the scope of the points provided.
func Normalize(pnts []Point) {
	maxPnt := MaxPoint(pnts)

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
}

// MaxPoint returns a point in which each dimension is the largest non-negative value observed in the set of points.
func MaxPoint(pnts []Point) Point {
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

// MeanVariances returns the mean variance of a set of clusters.
func MeanVariances(clstrs []Cluster) float64 {
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
func SortCluster(clstr Cluster) {
	if mn := Mean(clstr); mn != nil {
		sort.SliceStable(clstr, func(i, j int) bool { return SquaredDistance(mn, clstr[i]) < SquaredDistance(mn, clstr[j]) })
	}
}

// SortClusters sorts a set of clusters.
func SortClusters(clstrs []Cluster) {
	sort.SliceStable(clstrs, func(i, j int) bool { return CompareClusters(clstrs[i], clstrs[j]) < 0 })
}

// SortAllClusters sorts each cluster. The set of clusters is NOT sorted.
func SortAllClusters(clstrs []Cluster) {
	for i := range clstrs {
		SortCluster(clstrs[i])
	}
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

	var maxIndex, comparison int
	if m < n {
		maxIndex = m
	} else {
		maxIndex = n
	}

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
