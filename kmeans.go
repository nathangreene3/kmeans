package main

import (
	"math"
)

// Model holds k-means clusters and their meta data.
type Model struct {
	k          int       // Number of clusters
	clusters   Clusters  // Clusters returned from k-means
	means      Points    // Means of clusters
	variances  []float64 // Variances of clusters
	maxPoint   Point     // Max representing point
	normalized bool      // Indicates if normalized
}

// New returns a trained k-means model.
func New(k int, ps Points, normalize bool) *Model {
	mdl := &Model{}
	for i := 0; i < 5; i++ {
		mdl.Train(k, ps, normalize)
	}

	return mdl
}

// Assignment returns the index of the cluster a point belongs to.
func (mdl *Model) Assignment(pnt Point) int {
	var (
		assignment int
		sd         float64
		minSD      = math.MaxFloat64
	)

	if mdl.normalized {
		pnt.Normalize(mdl.maxPoint)
	}

	for i := range mdl.means {
		sd = SquaredDistance(mdl.means[i], pnt)
		if sd < minSD {
			minSD = sd
			assignment = i
		}
	}

	return assignment
}

// K returns the number of clusters.
func (mdl *Model) K() int {
	return mdl.k
}

// Mean returns the Mean of cluster i.
func (mdl *Model) Mean(i int) Point {
	return mdl.means[i].Copy()
}

// Means returns the mean points of the clusters.
func (mdl *Model) Means() Points {
	return mdl.means.Copy()
}

// Variance returns the variance of cluster i.
func (mdl *Model) Variance(i int) float64 {
	return mdl.variances[i]
}

// Variances returns a copy of the variances of the clustes.
func (mdl *Model) Variances() []float64 {
	cpy := make([]float64, 0, len(mdl.variances))
	copy(cpy, mdl.variances)
	return cpy
}

// MaxRepPoint returns the max representing point.
func (mdl *Model) MaxRepPoint() Point {
	return mdl.maxPoint.Copy()
}

// IsNormed returns true if the model was normalized and false if otherwise.
func (mdl *Model) IsNormed() bool {
	return mdl.normalized
}

// sortAll each cluster in a model.
func (mdl *Model) sortAll(st SortOpt) {
	mdl.clusters.SortAll(st)
}

// sort clusters.
func (mdl *Model) sort() {
	mdl.clusters.Sort()
	mdl.update()
}

// KMeans clusters a set of points into k groups. Potentially, clusters can be empty, so multiple attempts should be made.
func KMeans(k int, pnts Points, normalize bool) Clusters {
	// Move points to their nearest cluster until they no longer move with each pass (indicated by the changed boolean).
	var (
		clusters  = initClusters(k, pnts, normalize) // Clusters to return
		changed   = true                             // Indicates if a cluster was altered
		means     Points                             // Means of clusters
		minIndex  int                                // Index of cluster having smallest squared distance to a point
		sd, minSD float64                            // Squared distance; minimum squared distance
	)

	for changed {
		changed = false

		// Update the means. If any cluster is empty, its mean, which is nil, will be reassigned to be a random point on the space spanned by the maximum point.
		means = Means(clusters)
		for i := range means {
			if means[i] == nil {
				randomPoint(pnts)
			}
		}

		for h := range clusters {
			// Each cluster is sorted, so the most variant point is at index clstrs[h][sizes[h]-1]. When the point clstrs[h][i] variance is small enough to not move to another cluster, then we can move on to the next cluster and ignore the other points on the range [0,i). So, h counts down and halts early.
			for i := 0; i < len(clusters[h]); i++ {
				// Find the index of the cluster closest to point i in cluster h.
				minIndex = h
				minSD = SquaredDistance(means[h], clusters[h][i])
				for j := range clusters {
					if h == j {
						// If h = j, then we are comparing the same cluster. If the size of cluster j is zero, then, obviously, there's no points to compare.
						continue
					}

					sd = SquaredDistance(means[j], clusters[h][i])
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
				clusters[minIndex], clusters[h] = Transfer(i, clusters[minIndex], clusters[h])
				changed = true
			}
		}
	}

	return clusters
}

// Train clusters a set of points into k groups. Potentially, clusters can be empty, so multiple attempts should be made.
func (mdl *Model) Train(k int, ps Points, normalize bool) {
	mdl.initialize(k, ps, normalize)

	// Move points to their nearest cluster until they no longer move with each pass (indicated by the changed boolean).
	var (
		changed   = true  // Indicates if a cluster was altered
		minIndex  int     // Index of cluster having smallest squared distance to a point
		sd, minSD float64 // Squared distance; minimum squared distance
	)

	for changed {
		changed = false

		// Update the means and variances. If any cluster is empty, its mean, which is nil, will be reassigned to be a random point on the space spanned by the maximum point.
		mdl.update()
		for i := range mdl.means {
			if mdl.means[i] == nil {
				ps.Random()
			}
		}

		for h := range mdl.clusters {
			// Each cluster is sorted, so the most variant point is at index clstrs[h][sizes[h]-1]. When the point clstrs[h][i] variance is small enough to not move to another cluster, then we can move on to the next cluster and ignore the other points on the range [0,i). So, h counts down and halts early.
			for i := 0; i < len(mdl.clusters[h]); i++ {
				// Find the index of the cluster closest to point i in cluster h.
				minIndex = h
				minSD = mdl.means[h].SqDist(mdl.clusters[h][i])
				for j := range mdl.clusters {
					if h == j {
						// If h = j, then we are comparing the same cluster. If the size of cluster j is zero, then, obviously, there's no points to compare.
						continue
					}

					sd = mdl.means[j].SqDist(mdl.clusters[h][i])
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
				mdl.clusters[minIndex], mdl.clusters[h] = Transfer(i, mdl.clusters[minIndex], mdl.clusters[h])
				changed = true
			}
		}
	}
}

// initClusters returns a set of k-sorted clusters.
func (mdl *Model) initialize(k int, ps Points, normalize bool) {
	ps.validate()
	ps.Shuffle()
	if normalize {
		ps.Normalize()
	}

	mdl.k = k
	mdl.clusters = make(Clusters, 0, k)

	// Each cluster contains capacity/k points. Remainders will be added to the last cluster (index k-1).
	var (
		capacity = len(ps)      // Cluster capacity
		length   = capacity / k // Cluster length, not including remainder
		h        int            // Indexer through points
	)
	for i := 0; i < k; i++ {
		c := make(Cluster, 0, capacity)
		for j := 0; j < length; j++ {
			c = append(c, ps[h])
			h++
		}

		mdl.clusters = append(mdl.clusters, c)
	}

	// The last cluster is potentially largest if k doesn't divide n evenly. The actual size is s + n mod k. For example, given 32 points on 5 clusters, the size would be 32/5 + 32 mod 5 = 8.
	for ; h < capacity; h++ {
		mdl.clusters[k-1] = append(mdl.clusters[k-1], ps[h])
	}

	mdl.clusters.Coalesce() // Sorts all
	mdl.update()
}

// update the means and variances.
func (mdl *Model) update() {
	mdl.means = mdl.clusters.Means()
	mdl.variances = mdl.clusters.Variances()
}

// OptimalKMeans determines (attempts to anyway...) the optimal number of clusters and returns the clustering result.
func OptimalKMeans(pnts Points, normalize bool) (int, Clusters) {
	var (
		n        = len(pnts) // Number of points
		vars     []float64
		varsPnts Points // Vars converted to points
	)
	n = 10 // Halt at k = 1..n until testing is done.

	// if n <= 10 {
	// Run k-means on k = 1..n and track the mean variance.
	varsPnts = make(Points, 0, n)
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
	// }

	// return 0, nil
}
