package kmeans

import (
	"math"
	"math/rand"
	"sort"
	"strings"
)

// --------------------------------------------------------------------
// 	  k-Means algorithm (naive or Lloyd's algorithm)
// --------------------------------------------------------------------
// 1. Initialize means of each of k clusters with random data points.
// 2. Assign data points to each cluster.
// 3. Select a new mean for each cluster that minimizes the variance of
//    the cluster. If the cluster size is less than two, randomly
//    choose a new cluster mean from the data points.
// 4. Reassign any data points to their nearest cluster. Return the
//    means when no reassignments are made.
// --------------------------------------------------------------------

// --------------------------------------------------------------------
//    k-Means++
// --------------------------------------------------------------------
// 1. Initialize the first mean as a random data point.
// 2. Initialize remaining means with points that are farthest away
//    from initialized means.
// 3. Continue from step 2 of the naive algorithm.
// --------------------------------------------------------------------

// --------------------------------------------------------------------
//    Sources
// --------------------------------------------------------------------
//  * k-Means++ selects means more carefully before training to reduce
//    convergence iterations.
//    https://www.geeksforgeeks.org/ml-k-means-algorithm/
//  * The triangle inequality prevents unnecessary distance
//    calculations.
//    https://www.aaai.org/Papers/ICML/2003/ICML03-022.pdf
//  * Full explanation of k-means.
//    https://towardsdatascience.com/k-means-clustering-algorithm-applications-evaluation-methods-and-drawbacks-aa03e644b48a
// --------------------------------------------------------------------

// Model is a set of k points representing the means of their clusters.
// Implements the Model interface.
type Model []Point

// New returns a trained model. By default, the model initialized with
// random points from the data set and is trained once.
func New(k int, data []Point, opts ...Option) Model {
	if len(data) < k {
		panic(errDataSize)
	}

	var (
		cfg       = NewConfig(opts...)
		meanDists = newTriMatrix(k)
		mdl       = make(Model, 0, k)
		maxScrMdl = make(Model, 0, k)
		cls       = make(classes, len(data))
		maxScr    = -math.MaxFloat64
	)

	for i := 0; i < k; i++ {
		mdl = append(mdl, make(Point, len(data[0])))
		maxScrMdl = append(maxScrMdl, make(Point, len(data[0])))
	}

	for ; 0 < cfg.TrainRounds; cfg.TrainRounds-- {
		mdl.init(cfg.Mthd, meanDists, data)
		mdl.train(meanDists, cls, data)
		if score := mdl.Score(data...); maxScr < score {
			maxScrMdl.copyFrom(mdl)
			maxScr = score
		}
	}

	return maxScrMdl
}

// Class returns the classification of a point.
func (mdl Model) Class(datum Point) int {
	class, _ := mdl.classDist(datum)
	return class
}

// classDist returns the classification and distance between a data
// point and its mean.
func (mdl Model) classDist(datum Point) (int, float64) {
	var (
		class     int
		minSqDist = mdl[class].Dist(datum)
	)

	for i := 1; i < len(mdl); i++ {
		if dist := mdl[i].Dist(datum); dist < minSqDist {
			class = i
			minSqDist = dist
		}
	}

	return class, minSqDist
}

// classDistMem returns the classification and distance between a data
// point and its mean. Requires the distance between each mean be
// computed in advance, but otherwise behaves the same as classDist.
func (mdl Model) classDistMem(datum Point, meanDists triMatrix) (int, float64) {
	// Implements triangle inequality to prevent needless distance
	// calculations. If point p is assigned to cluster 0 with mean m0
	// and we want to know if cluster mean m1 is nearer, there's no
	// point in computing the distance from p to m1 if the distance
	// from p to m0 is less than half the distance from m0 to m1. That
	// is, if d(p, m0) <= d(m0, m1) / 2, then don't compute d(p, m1).

	// Note: SqDist fails here, so Dist must be used.

	var (
		class   int
		minDist = mdl[class].Dist(datum)
	)

	for i := 1; i < len(mdl); i++ {
		if meanDists.dist(class, i)/2.0 < minDist {
			if dist := mdl[i].Dist(datum); dist < minDist {
				class = i
				minDist = dist
			}
		}
	}

	return class, minDist
}

// Classes returns the classification of each data point.
func (mdl Model) Classes(data ...Point) []int {
	var (
		classes   = make([]int, 0, len(data))
		meanDists = newTriMatrix(len(mdl))
	)

	meanDists.update(mdl)
	for i := 0; i < len(data); i++ {
		class, _ := mdl.classDistMem(data[i], meanDists)
		classes = append(classes, class)
	}

	return classes
}

// Cluster returns the data that is classified in the given class.
func (mdl Model) Cluster(class int, data ...Point) []Point {
	var (
		cluster   = make([]Point, 0, len(data))
		meanDists = newTriMatrix(len(mdl))
	)

	meanDists.update(mdl)
	for i := 0; i < len(data); i++ {
		if c, _ := mdl.classDistMem(data[i], meanDists); c == class {
			cluster = append(cluster, data[i].Copy())
		}
	}

	return cluster
}

// Clusters returns the data classified into k clusters.
func (mdl Model) Clusters(data ...Point) [][]Point {
	var (
		classes = mdl.Classes(data...)
		sizes   = make([]int, len(mdl))
	)

	for i := 0; i < len(classes); i++ {
		sizes[classes[i]]++
	}

	clusters := make([][]Point, len(mdl))
	for i := 0; i < len(mdl); i++ {
		clusters[i] = make([]Point, 0, sizes[i])
	}

	for i := 0; i < len(data); i++ {
		clusters[classes[i]] = append(clusters[classes[i]], data[i].Copy())
	}

	return clusters
}

// copyFrom copies the source into a given model.
func (mdl Model) copyFrom(src Model) {
	for i := 0; i < len(mdl); i++ {
		copy(mdl[i], src[i])
	}
}

// Copy returns a copy of a model.
func (mdl Model) Copy() Model {
	cpy := make(Model, 0, len(mdl))
	for i := 0; i < len(mdl); i++ {
		cpy = append(cpy, mdl[i].Copy())
	}

	return cpy
}

// Dist returns the distance from a class mean to a given point.
func (mdl Model) Dist(class int, datum Point) float64 {
	return mdl[class].Dist(datum)
}

// Err returns the sum of squared distances of the subset of a given
// set of data that is classified as the given class and the number of
// points in the subset that were assigned to the given class.
func (mdl Model) Err(class int, data ...Point) float64 {
	var err float64 // e = sum((xi-m)^2, i = 0, 1, 2,...)
	for i := 0; i < len(data); i++ {
		if c, d := mdl.classDist(data[i]); c == class {
			err += d * d
		}
	}

	return err
}

// Errs classifies a given set of data and returns the variance for
// each cluster.
func (mdl Model) Errs(data ...Point) []float64 {
	errs := make([]float64, len(mdl))
	for i := 0; i < len(data); i++ {
		class, dist := mdl.classDist(data[i])
		errs[class] += dist * dist
	}

	return errs
}

// init initializes a model by a given method. The mean distances may
// be updated as the method requires.
func (mdl Model) init(mthd InitMethod, meanDists triMatrix, data []Point) {
	switch mthd {
	case Random:
		var (
			partSize = len(data) / len(mdl)
			i, j     int
		)

		for i < len(mdl)-1 {
			copy(mdl[i], data[rand.Intn(partSize)+j])
			i++
			j += partSize
		}

		copy(mdl[i], data[rand.Intn(len(data)-j)+j])
		meanDists.update(mdl)
	case PlusPlus:
		copy(mdl[0], data[rand.Intn(len(data))])
		meanDists.update(mdl)

		for i := 1; i < len(mdl); i++ {
			var (
				maxJ    int
				maxDist float64
			)

			for j := 0; j < len(data); j++ {
				if _, dist := mdl.classDistMem(data[j], meanDists); maxDist < dist {
					maxJ = j
					maxDist = dist
				}
			}

			copy(mdl[i], data[maxJ])
			meanDists.update(mdl)
		}
	case FirstK:
		mdl.copyFrom(data[:len(mdl)])
		meanDists.update(data)
	default:
		panic(errInitMthd)
	}
}

// K returns the number of clusters k.
func (mdl Model) K() int {
	return len(mdl)
}

// Mean returns the mean of the specified class.
func (mdl Model) Mean(class int) Point {
	return mdl[class].Copy()
}

// Means returns the means.
func (mdl Model) Means() []Point {
	mns := make([]Point, 0, len(mdl))
	for i := 0; i < len(mdl); i++ {
		mns = append(mns, mdl[i].Copy())
	}

	return mns
}

// Score indicates how well a model clusters data. A higher score
// inidicates the model is a better fit.
func (mdl Model) Score(data ...Point) float64 {
	return -sum(mdl.Errs(data...)...)
}

// Size returns the size of the specified cluster.
func (mdl Model) Size(class int, data ...Point) int {
	var size int
	for i := 0; i < len(data); i++ {
		if c, _ := mdl.classDist(data[i]); c == class {
			size++
		}
	}

	return size
}

// Sizes returns the sizes of each cluster.
func (mdl Model) Sizes(data ...Point) []int {
	sizes := make([]int, len(mdl))
	for i := 0; i < len(data); i++ {
		class, _ := mdl.classDist(data[i])
		sizes[class]++
	}

	return sizes
}

// Sort the model.
func (mdl Model) Sort() {
	sort.Slice(mdl, func(i, j int) bool { return mdl[i].Compare(mdl[j]) < 0 })
}

// SqDist returns the squared distance from a class mean to a given
// point.
func (mdl Model) SqDist(class int, datum Point) float64 {
	return mdl[class].SqDist(datum)
}

// String returns a representation of a model formatted as
// ((0.0, ..., 0.0), ..., (0.0, ..., 0.0))
func (mdl Model) String() string {
	var sb strings.Builder
	if len(mdl) != 0 {
		sb.WriteString("(" + mdl[0].String())
		for i := 1; i < len(mdl); i++ {
			sb.WriteString(", " + mdl[i].String())
		}

		sb.WriteByte(')')
	}

	return sb.String()
}

// Train updates the means using the given data set.
func (mdl Model) Train(data ...Point) {
	var (
		meanDists = newTriMatrix(len(mdl))
		cls       = make(classes, len(data))
	)

	meanDists.update(mdl)
	mdl.train(meanDists, cls, data)
}

// train updates the means using the given data set and mean distance
// lookup table.
func (mdl Model) train(meanDists triMatrix, cls classes, data []Point) {
	for i := 0; ; i++ {
		if !cls.update(mdl, meanDists, data) {
			return
		}

		mdl.update(cls, data)
		meanDists.update(mdl)
	}
}

// update the model with data points as new means that have the
// smallest variance in their respective class.
func (mdl Model) update(cls classes, data []Point) {
	if len(cls) != len(data) {
		panic(errDims)
	}

	for i := 0; i < len(mdl); i++ {
		for j := 0; j < len(mdl[i]); j++ {
			mdl[i][j] = 0.0
		}

		var size float64
		for j := 0; j < len(data); j++ {
			if i == cls[j] {
				mdl[i].Add(data[j])
				size++
			}
		}

		if size == 0 {
			// Cluster is empty; randomly select a representative
			mdl[i].Add(data[rand.Intn(len(data))])
			continue
		}

		mdl[i].ScalMult(1.0 / size)
	}
}
