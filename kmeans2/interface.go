package kmeans2

// Interface defines the requirements for clustering.
type Interface interface {
	Compare(p Interface) int
	Dist(p Interface) float64
	Len() int
}
