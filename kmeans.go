package kmeans

import (
	"fmt"
	"math"

	graph "github.com/guptarohit/asciigraph"
)

// Model holds k-means clusters and their meta data.
type Model struct {
	k         int       // Number of clusters
	clusters  Clusters  // Clusters returned from k-means
	means     Points    // Means of clusters
	variances []float64 // Variances of clusters
}

// New returns a trained k-means model.
func New(k int, ps Points) *Model {
	var (
		mdl                     = &Model{}
		meanWeightedVariance    float64
		minMeanWeightedVariance = math.MaxFloat64
		minClusters             Clusters
	)

	for i := 0; i < 5; i++ {
		mdl.Train(k, ps)
		if meanWeightedVariance = mdl.MeanWeightedVariance(); meanWeightedVariance < minMeanWeightedVariance {
			minClusters = mdl.clusters
			minMeanWeightedVariance = meanWeightedVariance
		}
	}

	mdl.clusters = minClusters
	mdl.update()
	return mdl
}

// Assignment returns the index of the cluster a point belongs to.
func (mdl *Model) Assignment(pnt Point) int {
	var (
		assignment int
		sqDist     float64
		minSqDist  = math.MaxFloat64
	)

	for i := range mdl.means {
		if sqDist = pnt.SqDist(mdl.means[i]); sqDist < minSqDist {
			minSqDist = sqDist
			assignment = i
		}
	}

	return assignment
}

// initClusters returns a set of k-sorted clusters.
func (mdl *Model) initialize(k int, ps Points) {
	ps.validate()
	ps.Shuffle()

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

// K returns the number of clusters.
func (mdl *Model) K() int {
	return mdl.k
}

// Mean returns the mean of cluster i.
func (mdl *Model) Mean(i int) Point {
	return mdl.means[i].Copy()
}

// Means returns the mean points of the clusters.
func (mdl *Model) Means() Points {
	return mdl.means.Copy()
}

// MeanWeightedVariance returns the mean of the variances weighted by the number of points in each cluster.
func (mdl *Model) MeanWeightedVariance() float64 {
	return mdl.clusters.MeanWeightedVariance()
}

// PlotMeanWeightedVars returns a string representing a chart of the mean variances of several models over a range of k in [kMin, kMax].
func PlotMeanWeightedVars(kMin, kMax int, ps Points) string {
	if kMax < kMin {
		kMin, kMax = kMax, kMin
	}

	mnVars := make([]float64, 0, kMax-kMin+1)
	for k := kMin; k <= kMax; k++ {
		mnVars = append(mnVars, New(k, ps).MeanWeightedVariance())
	}

	caption := fmt.Sprintf("Mean Variances of k-Means Trials for k in [%d,%d]", kMin, kMax)
	return graph.Plot(mnVars, graph.Caption(caption))
}

// sort clusters.
func (mdl *Model) sort() {
	mdl.clusters.Sort()
	mdl.update()
}

// sortAll sorts each cluster in a model. The order of the clusters is NOT sorted.
func (mdl *Model) sortAll(st SortOpt) {
	mdl.clusters.SortAll(st)
}

// Train clusters a set of points into k groups. Potentially, clusters can be empty, so multiple attempts should be made.
func (mdl *Model) Train(k int, ps Points) {
	mdl.initialize(k, ps)

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

// update the means and variances.
func (mdl *Model) update() {
	mdl.means = mdl.clusters.Means()
	mdl.variances = mdl.clusters.Variances()
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
