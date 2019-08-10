package kmeans

// Point is an n-dimensional point in n-space.
type Point interface {
	At(i int) float64
	CompareTo(p Point) int
	Copy() Point
	Dist(p Point) float64
	Len() int
	SqDist(p Point) float64
}
