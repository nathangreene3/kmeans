package kmeans

// Point is an n-dimensional point in n-space.
type Point interface {
	At(i int) float64
	CompareTo(pnt Point) int
	Copy() Point
	Dist(pnt Point) float64
	Len() int
	SqDist(pnt Point) float64
}
