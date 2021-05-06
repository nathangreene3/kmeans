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
//    iterations.
// 2. Use the triangle inequality to reduce distance calculations.
// 3. Implement a variant for mini batching to handle large data sets.
// ------------------------------------------------------------------------------------

// Model defines the k-means model.
type Model interface {
	Class(Point) int
	Clusters(...Point) [][]Point
	Dist(int, Point) float64
	Err(int, ...Point) (float64, int)
	Errs(...Point) ([]float64, []int)
	K() int
	Mean(int) Point
	Score(...Point) float64
	Train(...Point) Model
}

// kMeans is a set of k points representing the means of their clusters. Implements the Model interface.
type kMeans []Point

// Init returns a model ready for training initialized with a given set of points as the cluster means.
func Init(means ...Point) Model {
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

	var (
		maxScoreKMns kMeans             // Model having the highest score to return
		maxScore     = -math.MaxFloat64 // Score of model
	)

	for ; 0 < trains; trains-- {
		var (
			kmns    = make(kMeans, 0, k) // Cluster means
			binSize = len(data) / k      // Partition size of data points each mean is selected from
		)

		if k*binSize < len(data) {
			// binSize needs to be ceil(n/k) but is currently floor(n/k)
			binSize++
		}

		// Initialize cluster means
		for j := 0; len(kmns) < cap(kmns); j += binSize {
			if len(data) < j+binSize {
				binSize = len(data) - j
			}

			kmns = append(kmns, data[j+rand.Intn(binSize)])
		}

		if s := kmns.Train(data...).Score(data...); maxScore < s {
			maxScoreKMns = kmns
			maxScore = s
		}
	}

	return maxScoreKMns
}

// Class returns the cluster classification a point belongs to (i.e. is closest to).
func (kmns kMeans) Class(datum Point) int {
	i, _ := kmns.classDist(datum)
	return i
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

// classDist returns the classification and distance between a data point and its mean.
func (kmns kMeans) classDist(datum Point) (int, float64) {
	var (
		i    int
		minD = math.MaxFloat64
	)

	for j := 0; j < len(kmns); j++ {
		if d := kmns[j].Dist(datum); d < minD {
			i = j
			minD = d
		}
	}

	return i, minD
}

// Dist returns the distance from cluster class to a given point.
func (kmns kMeans) Dist(class int, datum Point) float64 { return kmns[class].Dist(datum) }

// K returns the number of clusters k.
func (kmns kMeans) K() int { return len(kmns) }

// Mean returns the mean of the specified class.
func (kmns kMeans) Mean(class int) Point { return kmns[class].Copy() }

// Train returns a trained model on a set of data.
func (kmns kMeans) Train(data ...Point) Model {
	classes := make([]int, len(data)) // Cluster assignments; the ith data point is assigned to the ith value in assigns; that is, data[i] is closest to kmns[assigns[i]]

	// Train until no more reassignments are made
	for {
		var c bool // Indicates if assignments are changed

		// Update assignments
		for i := 0; i < len(data); i++ {
			j := classes[i]
			classes[i] = kmns.Class(data[i])
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
		for h := 0; h < len(kmns); h++ {
			var (
				minI    int               // Index i minimizing the variance of cluster h
				minVars = math.MaxFloat64 // Variance of cluster h after mean is updated
			)

			for i := 0; i < len(data); i++ {
				if h == classes[i] {
					var (
						vars float64      // Variance of cluster h with respect to data point i
						size float64 = -1 // Size of cluster h
					)

					// Compute the variance of cluster h to data point i
					for j := 0; j < len(data); j++ {
						if classes[i] == classes[j] {
							d := data[i].SqDist(data[j])
							vars += d
							size++
						}
					}

					if 0 < size {
						vars /= size
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

// Score ...
func (kmns kMeans) Score(data ...Point) float64 {
	var (
		errs, sizes  = kmns.Errs(data...)
		meanWghtErrs float64
	)

	for i := 0; i < len(errs); i++ {
		meanWghtErrs += errs[i] / float64(sizes[i])
	}

	return -meanWghtErrs / float64(len(data)) // Lower mean weighted error means higher score
}
