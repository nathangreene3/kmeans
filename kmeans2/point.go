package kmeans2

// Point ...
type Point interface {
	Copy() Point
	Dist(Point) float64
	Equals(Point) bool
	SqDist(Point) float64

	// These are needed if mean representation is desired over discrete representation and if mean representation can be made to actually work
	Add(Point) Point
	Mult(float64) Point
}
