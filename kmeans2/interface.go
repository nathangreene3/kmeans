package kmeans2

// Interface ...
type Interface interface {
	Compare(p Interface) int
	Dist(p Interface) float64
	Len() int
}
