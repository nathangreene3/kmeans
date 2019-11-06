package kmeans2

import (
	"fmt"
	"math"

	graph "github.com/guptarohit/asciigraph"
)

// Model holds k-means clusters and their meta data.
type Model struct {
	k        int
	clusters []Cluster
}

// New returns a trained k-means model.
func New(k int, points Points) *Model {
	var (
		model                   Model
		minClusters             []Cluster
		minMeanWeightedVariance = math.MaxFloat64
	)

	// Number of training sessions is arbitrarily set to five.
	for i := 0; i < 5; i++ {
		model.Train(k, points)
		if meanWeightedVariance := model.MeanWeightedVariance(); meanWeightedVariance < minMeanWeightedVariance {
			minClusters = model.clusters
			minMeanWeightedVariance = meanWeightedVariance
		}
	}

	model.clusters = minClusters
	model.update()
	return &model
}

// Assignment returns the index of the cluster a point belongs to.
func (mdl *Model) Assignment(point Point) int {
	var (
		assignment int
		minDist    = math.MaxFloat64
	)

	for i, clstr := range mdl.clusters {
		if dist := point.Dist(clstr.mean); dist < minDist {
			minDist = dist
			assignment = i
		}
	}

	return assignment
}

// initialize a set of clusters.
func (mdl *Model) initialize(k int, ps Points) {
	ps.validate()
	ps.Shuffle()

	mdl.k = k
	mdl.clusters = make([]Cluster, 0, k)

	// Each cluster contains capacity/k points. Remainders will be added to the
	// last cluster (index k-1).
	var (
		capacity = len(ps)      // Cluster capacity
		length   = capacity / k // Cluster length, not including remainder
		h        int            // Indexer through points
	)

	for i := 0; i < k; i++ {
		mdl.clusters = append(mdl.clusters, NewCluster(ps[h:h+length]))
	}

	// The last cluster is potentially largest if k doesn't divide n evenly. The
	// actual size is s + n mod k. For example, given 32 points on 5 clusters,
	// the size would be 32/5 + 32 mod 5 = 8.
	for ; h < capacity; h++ {
		mdl.clusters[k-1].Append(ps[h])
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
	return mdl.clusters[i].mean.Copy()
}

// Means returns the mean points of the clusters.
func (mdl *Model) Means() Points {
	means := make(Points, 0, mdl.k)
	for _, c := range mdl.clusters {
		means = append(means, c.mean.Copy())
	}

	return means
}

// MeanWeightedVariance returns the mean of the variances weighted by the number
// of points in each cluster.
func (mdl *Model) MeanWeightedVariance() float64 {
	return mdl.clusters.MeanWeightedVariance()
}

// PlotMeanWeightedVars returns a string representing a chart of the mean
// variances of several models over a range of k in [kMin, kMax].
func PlotMeanWeightedVars(kMin, kMax int, points Points) string {
	if kMax < kMin {
		kMin, kMax = kMax, kMin
	}

	meanVariances := make([]float64, 0, kMax-kMin+1)
	for k := kMin; k <= kMax; k++ {
		meanVariances = append(meanVariances, New(k, points).MeanWeightedVariance())
	}

	caption := fmt.Sprintf("Mean Variances of k-Means Trials for k in [%d,%d]", kMin, kMax)
	return graph.Plot(meanVariances, graph.Caption(caption))
}

// sort clusters.
func (mdl *Model) sort() {
	mdl.clusters.Sort()
	mdl.update()
}

// sortAll sorts each cluster in a model. The order of the clusters is NOT
// sorted.
func (mdl *Model) sortAll(sortOpt SortOption) {
	mdl.clusters.SortAll(sortOpt)
}

// Train clusters a set of points into k groups. Potentially, clusters can be
// empty, so multiple attempts should be made.
func (mdl *Model) Train(k int, points Points) {
	mdl.initialize(k, points)

	// Move points to their nearest cluster until they no longer move with each
	// pass (indicated by the changed boolean).
	z := Zero()
	for changed := true; changed; {
		changed = false

		// Ensure means are defined, even on empty clusters.
		mdl.update()
		for i, mn := range mdl.means {
			if mn.Compare(z) == 0 {
				mdl.means[i] = points.Random()
				changed = true
			}
		}

		// Update each cluster h.
		for h := 0; h < mdl.k; h++ {
			// Update cluster assignments for each point p in cluster h.
			for i := 0; i < len(mdl.clusters[h]); i++ {
				// Find the index j of the cluster closest to point i that is in cluster h.
				var (
					p        = mdl.clusters[h][i]
					minIndex = h
					minDist  = mdl.means[h].Dist(p)
				)

				for j, mn := range mdl.means {
					if h != j {
						if dist := mn.Dist(p); dist < minDist {
							// Cluster j is closer to point i than its current cluster h.
							minDist = dist
							minIndex = j
						}
					}
				}

				if h != minIndex {
					// Transfer point from cluster h to minIndex.
					mdl.clusters[minIndex], mdl.clusters[h] = Transfer(i, mdl.clusters[minIndex], mdl.clusters[h])
					changed = true
				}
			}
		}
	}
}

// update the means and variances.
func (mdl *Model) update() {
	for i := range mdl.clusters {
		mdl.clusters[i].Update()
	}
}

// Variance returns the variance of cluster i.
func (mdl *Model) Variance(i int) float64 {
	return mdl.clusters[i].variance
}

// Variances returns a copy of the variances of the clusters.
func (mdl *Model) Variances() []float64 {
	variances := make([]float64, 0, mdl.k)
	for _, c := range mdl.clusters {
		variances = append(variances, c.variance)
	}

	return variances
}
