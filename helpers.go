package kmeans

// maxInt returns the largest value of a set of numbers.
func maxInt(n ...int) int {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if m < n[i] {
			m = n[i]
		}
	}

	return m
}

// minInt returns the smallest value of a set of numbers.
func minInt(n ...int) int {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] < m {
			m = n[i]
		}
	}

	return m
}
