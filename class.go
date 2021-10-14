package kmeans

// classes holds classifications for a set of data points to a k-means
// model. The ith data point belongs to the ith classification. That
// is, given a list of data points and means, point i belongs to
// classes i corresponding to one of the k means.
type classes []int

// update updates each classification and determines if any changes
// were made.
func (cls classes) update(mdl Model, meanDists triMatrix, data []Point) bool {
	var changed bool
	for i := 0; i < len(data); i++ {
		prevClass := cls[i]
		cls[i], _ = mdl.classDistMem(data[i], meanDists)
		changed = changed || cls[i] != prevClass
	}

	return changed
}
