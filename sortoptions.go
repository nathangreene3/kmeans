package kmeans

// SortOption indicates how a cluster is sorted.
type SortOption int

const (
	// SortByVariance dictates points will be compared by variance
	// to the cluster mean.
	SortByVariance SortOption = 1 + iota
	// SortByDimension dictates points will be compared by
	// dictionary-order.
	SortByDimension
)
