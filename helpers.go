package kmeans

// sum returns the sum of a list of values.
func sum(xs ...float64) float64 {
	var s float64
	for i := 0; i < len(xs); i++ {
		s += xs[i]
	}

	return s
}
