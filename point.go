package kmeans

// Point is an n-dimensional point in n-space.
type Point interface {
	At(i int) float64      // Return the ith dimension
	CompareTo(p Point) int // Return a value from {-1,0,1}
	Copy() Point           // Return a copy of a point
	Dist(p Point) float64  // Return the distance between points
	Len() int              // Return the number of dimensions of a point
}
