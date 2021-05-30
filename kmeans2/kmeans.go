package kmeans2

import (
	"math"
	"math/rand"
)

// ------------------------------------------------------------------------------------
// k-Means algorithm (naive or Lloyd's algorithm)
// ------------------------------------------------------------------------------------
// 1. Initialize means of each of k clusters with random data points.
// 2. Assign data points to each cluster.
// 3. Select a new mean for each cluster that minimizes the variance of the cluster.
//    If the cluster size is less than two, randomly choose a new cluster mean from
//    the data points.
// 4. Reassign any data points to their nearest cluster. Return the means when no
//    reassignments are made.
// ------------------------------------------------------------------------------------

// ------------------------------------------------------------------------------------
// TODO
// ------------------------------------------------------------------------------------
// 1. k-Means++ selects means more carefully before training to reduce convergence
//    iterations. DONE.
//    https://www.geeksforgeeks.org/ml-k-means-algorithm/
// 2. Use the triangle inequality to reduce distance calculations. DONE.
//    https://www.aaai.org/Papers/ICML/2003/ICML03-022.pdf
// 3. Implement a variant for mini batching to handle large data sets.
// ------------------------------------------------------------------------------------

// kMeans is a set of k points representing the means of their clusters. Implements the Model interface.
type kMeans []Point

// New returns a model ready for training initialized with a given set of points as the cluster means.
func New(means ...Point) Model {
	kmns := make(kMeans, 0, len(means))
	for i := 0; i < len(means); i++ {
		kmns = append(kmns, means[i].Copy())
	}

	return kmns
}

// KMeans returns a set of k points representing the means of their clusters.
func KMeans(k, trains int, data ...Point) Model {
	if len(data) < k {
		panic("insufficient number of data points")
	}

	var maxScoreKMns kMeans // Model having the highest score to return
	for maxScore := -math.MaxFloat64; 0 < trains; trains-- {
		var (
			kmns    = make(kMeans, 0, k) // Cluster means
			binSize = len(data) / k      // Partition size of data points each mean is selected from
		)

		if k*binSize < len(data) {
			// binSize needs to be ceil(n/k) but is currently floor(n/k)
			binSize++
		}

		// Random initialization
		/*
			for i := 0; i+binSize <= len(data); i += binSize {
				kmns = append(kmns, data[i+rand.Intn(binSize)])
			}

			if len(kmns) < cap(kmns) {
				kmns = append(kmns, data[len(kmns)+rand.Intn(cap(kmns)-len(kmns))])
			}
		*/

		// kmeans++ initialization
		// /*
		kmns = append(kmns, data[rand.Intn(len(data))])
		for i := 1; i < k; i++ {
			var (
				meanDists = newDistMtx(kmns...)
				maxJ      int
				maxDist   float64
			)

			for j := 0; j < len(data); j++ {
				_, d := kmns.classDistMem(data[j], meanDists)
				if maxDist < d {
					maxJ = j
					maxDist = d
				}
			}

			kmns = append(kmns, data[maxJ])
		}
		// */

		if s := kmns.Train(data...).Score(data...); maxScore < s {
			maxScoreKMns = kmns
			maxScore = s
		}
	}

	return maxScoreKMns
}

// Clusters ...
func (kmns kMeans) Clusters(data ...Point) [][]Point {
	clusters := make([][]Point, len(kmns))
	for i := 0; i < len(data); i++ {
		class := kmns.Class(data[i])
		clusters[class] = append(clusters[class], data[i].Copy())
	}

	return clusters
}

// Class returns the cluster classification a point belongs to (i.e. is closest to).
func (kmns kMeans) Class(datum Point) int {
	i, _ := kmns.classDist(datum)
	return i
}

// classDist returns the classification and distance between a data point and its mean.
func (kmns kMeans) classDist(datum Point) (int, float64) {
	var (
		i    int
		minD = kmns[i].Dist(datum)
	)

	for j := 1; j < len(kmns); j++ {
		if d := kmns[j].Dist(datum); d < minD {
			i = j
			minD = d
		}
	}

	return i, minD
}

// classDistMem returns the classification and distance between a data point and its
// mean. Requires the distance between each mean be computed in advance, but otherwise
// behaves the same as classDist.
func (kmns kMeans) classDistMem(datum Point, meanDists triMatrix) (int, float64) {
	// Implements triangle inequality to prevent needless distance calculations. If
	// point p is assigned to cluster 0 with mean m0 and we want to know if cluster
	// mean m1 is nearer, there's no point in computing the distance from p to m1 if
	// the distance from p to m0 is less than half the distance from m0 to m1. That
	// is, if d(p, m0) <= d(m0, m1) / 2, then don't compute d(p, m1).

	var (
		class int
		dist  = kmns[class].Dist(datum)
	)

	for j := 1; j < len(kmns); j++ {
		if dist <= meanDists.dist(class, j)/2 {
			continue
		}

		if d := kmns[j].Dist(datum); d < dist {
			class, dist = j, d
		}
	}

	return class, dist
}

// Dist returns the distance from cluster class to a given point.
func (kmns kMeans) Dist(class int, datum Point) float64 { return kmns[class].Dist(datum) }

// K returns the number of clusters k.
func (kmns kMeans) K() int { return len(kmns) }

// Mean returns the mean of the specified class.
func (kmns kMeans) Mean(class int) Point { return kmns[class].Copy() }

// Train returns a trained model on a set of data.
func (kmns kMeans) Train(data ...Point) Model {
	var (
		classes   = make([]int, len(data)) // Cluster assignments; the ith data point is assigned to the ith value in assigns; that is, data[i] is closest to kmns[assigns[i]]
		meanDists = newDistMtx(kmns...)    // Distances between cluster means
	)

	// Train until no more reassignments are made
	for {
		var c bool // Indicates if assignments are changed

		// Update assignments
		for i := 0; i < len(data); i++ {
			j := classes[i]
			classes[i], _ = kmns.classDistMem(data[i], meanDists)
			c = c || classes[i] != j
		}

		if !c {
			// No assignments were changed; Deep copy and return
			for i := 0; i < len(kmns); i++ {
				kmns[i] = kmns[i].Copy()
			}

			return kmns
		}

		// Update cluster means
		// /*
		for h := 0; h < len(kmns); h++ {
			var (
				minI    int               // Index i minimizing the variance of cluster h
				minVars = math.MaxFloat64 // Variance of cluster h after mean is updated
			)

			for i := 0; i < len(data); i++ {
				if h == classes[i] {
					// Compute the variance of cluster h to data point i
					var vars, size float64
					for j := 0; j < len(data); j++ {
						if classes[i] == classes[j] {
							vars += data[i].SqDist(data[j])
							size++
						}
					}

					if 1 < size {
						vars /= (size - 1)
					}

					// Update the index i that minimizes the variance for cluster h
					if vars < minVars {
						minI = i
						minVars = vars
					}
				}
			}

			// Update the mean of cluster h
			kmns[h] = data[minI]
		}
		// */

		// This updates each center as the actual mean of its cluster, not as a representative of the cluster. However, it doesn't work and usually results in an infinite cycle.
		/*
			for i := 0; i < len(kmns); i++ {
				kmns[i] = kmns[i].Mult(0)

				var size float64
				for j := 0; j < len(data); j++ {
					if i == classes[j] {
						kmns[i] = kmns[i].Add(data[j])
						size++
					}
				}

				if size == 0 {
					kmns[i] = data[rand.Intn(len(data))].Copy()
				} else {
					kmns[i] = kmns[i].Mult(1.0 / size)
				}
			}
		*/

		meanDists.update(kmns...)
	}
}

// Err returns the sum of squared distances of the subset of a given set of data that is classified as the given class.
func (kmns kMeans) Err(class int, data ...Point) (float64, int) {
	var (
		err  float64
		size int
	)

	for i := 0; i < len(data); i++ {
		if class == kmns.Class(data[i]) {
			err += kmns[class].SqDist(data[i])
			size++
		}
	}

	return err, size
}

// Errs classifies a given set of data and returns the variance for each cluster.
func (kmns kMeans) Errs(data ...Point) ([]float64, []int) {
	var (
		errs  = make([]float64, len(kmns))
		sizes = make([]int, len(kmns))
	)

	for i := 0; i < len(data); i++ {
		class, dist := kmns.classDist(data[i])
		errs[class] += dist * dist
		sizes[class]++
	}

	return errs, sizes
}

// Score indicates how well a model clusters data. A higher score inidicates the model is a better fit.
func (kmns kMeans) Score(data ...Point) float64 {
	var (
		errs, sizes = kmns.Errs(data...)
		meanErrs    float64 // e := (e0/n0 + e1/n1 + ...) / k
	)

	for i := 0; i < len(errs); i++ {
		meanErrs += errs[i] / float64(sizes[i])
	}

	return -meanErrs / float64(len(kmns)) // Lower mean weighted error means higher score
}
