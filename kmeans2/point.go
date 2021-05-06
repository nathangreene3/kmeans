package kmeans2

// Point ...
type Point interface {
	Copy() Point
	Dist(Point) float64
	Equals(Point) bool
	SqDist(Point) float64
	// String() string
}
