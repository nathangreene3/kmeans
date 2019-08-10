package kmeans

import (
	"sort"
	"testing"
)

var species = map[string]Points{
	"setosa": Points{
		FPoint{5.1, 3.5, 1.4, 0.2},
		FPoint{4.9, 3.0, 1.4, 0.2},
		FPoint{4.7, 3.2, 1.3, 0.2},
		FPoint{4.6, 3.1, 1.5, 0.2},
		FPoint{5.0, 3.6, 1.4, 0.2},
		FPoint{5.4, 3.9, 1.7, 0.4},
		FPoint{4.6, 3.4, 1.4, 0.3},
		FPoint{5.0, 3.4, 1.5, 0.2},
		FPoint{4.4, 2.9, 1.4, 0.2},
		FPoint{4.9, 3.1, 1.5, 0.1},
		FPoint{5.4, 3.7, 1.5, 0.2},
		FPoint{4.8, 3.4, 1.6, 0.2},
		FPoint{4.8, 3.0, 1.4, 0.1},
		FPoint{4.3, 3.0, 1.1, 0.1},
		FPoint{5.8, 4.0, 1.2, 0.2},
		FPoint{5.7, 4.4, 1.5, 0.4},
		FPoint{5.4, 3.9, 1.3, 0.4},
		FPoint{5.1, 3.5, 1.4, 0.3},
		FPoint{5.7, 3.8, 1.7, 0.3},
		FPoint{5.1, 3.8, 1.5, 0.3},
		FPoint{5.4, 3.4, 1.7, 0.2},
		FPoint{5.1, 3.7, 1.5, 0.4},
		FPoint{4.6, 3.6, 1.0, 0.2},
		FPoint{5.1, 3.3, 1.7, 0.5},
		FPoint{4.8, 3.4, 1.9, 0.2},
		FPoint{5.0, 3.0, 1.6, 0.2},
		FPoint{5.0, 3.4, 1.6, 0.4},
		FPoint{5.2, 3.5, 1.5, 0.2},
		FPoint{5.2, 3.4, 1.4, 0.2},
		FPoint{4.7, 3.2, 1.6, 0.2},
		FPoint{4.8, 3.1, 1.6, 0.2},
		FPoint{5.4, 3.4, 1.5, 0.4},
		FPoint{5.2, 4.1, 1.5, 0.1},
		FPoint{5.5, 4.2, 1.4, 0.2},
		FPoint{4.9, 3.1, 1.5, 0.1},
		FPoint{5.0, 3.2, 1.2, 0.2},
		FPoint{5.5, 3.5, 1.3, 0.2},
		FPoint{4.9, 3.1, 1.5, 0.1},
		FPoint{4.4, 3.0, 1.3, 0.2},
		FPoint{5.1, 3.4, 1.5, 0.2},
		FPoint{5.0, 3.5, 1.3, 0.3},
		FPoint{4.5, 2.3, 1.3, 0.3},
		FPoint{4.4, 3.2, 1.3, 0.2},
		FPoint{5.0, 3.5, 1.6, 0.6},
		FPoint{5.1, 3.8, 1.9, 0.4},
		FPoint{4.8, 3.0, 1.4, 0.3},
		FPoint{5.1, 3.8, 1.6, 0.2},
		FPoint{4.6, 3.2, 1.4, 0.2},
		FPoint{5.3, 3.7, 1.5, 0.2},
		FPoint{5.0, 3.3, 1.4, 0.2},
	},
	"versicolor": Points{
		FPoint{7.0, 3.2, 4.7, 1.4},
		FPoint{6.4, 3.2, 4.5, 1.5},
		FPoint{6.9, 3.1, 4.9, 1.5},
		FPoint{5.5, 2.3, 4.0, 1.3},
		FPoint{6.5, 2.8, 4.6, 1.5},
		FPoint{5.7, 2.8, 4.5, 1.3},
		FPoint{6.3, 3.3, 4.7, 1.6},
		FPoint{4.9, 2.4, 3.3, 1.0},
		FPoint{6.6, 2.9, 4.6, 1.3},
		FPoint{5.2, 2.7, 3.9, 1.4},
		FPoint{5.0, 2.0, 3.5, 1.0},
		FPoint{5.9, 3.0, 4.2, 1.5},
		FPoint{6.0, 2.2, 4.0, 1.0},
		FPoint{6.1, 2.9, 4.7, 1.4},
		FPoint{5.6, 2.9, 3.6, 1.3},
		FPoint{6.7, 3.1, 4.4, 1.4},
		FPoint{5.6, 3.0, 4.5, 1.5},
		FPoint{5.8, 2.7, 4.1, 1.0},
		FPoint{6.2, 2.2, 4.5, 1.5},
		FPoint{5.6, 2.5, 3.9, 1.1},
		FPoint{5.9, 3.2, 4.8, 1.8},
		FPoint{6.1, 2.8, 4.0, 1.3},
		FPoint{6.3, 2.5, 4.9, 1.5},
		FPoint{6.1, 2.8, 4.7, 1.2},
		FPoint{6.4, 2.9, 4.3, 1.3},
		FPoint{6.6, 3.0, 4.4, 1.4},
		FPoint{6.8, 2.8, 4.8, 1.4},
		FPoint{6.7, 3.0, 5.0, 1.7},
		FPoint{6.0, 2.9, 4.5, 1.5},
		FPoint{5.7, 2.6, 3.5, 1.0},
		FPoint{5.5, 2.4, 3.8, 1.1},
		FPoint{5.5, 2.4, 3.7, 1.0},
		FPoint{5.8, 2.7, 3.9, 1.2},
		FPoint{6.0, 2.7, 5.1, 1.6},
		FPoint{5.4, 3.0, 4.5, 1.5},
		FPoint{6.0, 3.4, 4.5, 1.6},
		FPoint{6.7, 3.1, 4.7, 1.5},
		FPoint{6.3, 2.3, 4.4, 1.3},
		FPoint{5.6, 3.0, 4.1, 1.3},
		FPoint{5.5, 2.5, 4.0, 1.3},
		FPoint{5.5, 2.6, 4.4, 1.2},
		FPoint{6.1, 3.0, 4.6, 1.4},
		FPoint{5.8, 2.6, 4.0, 1.2},
		FPoint{5.0, 2.3, 3.3, 1.0},
		FPoint{5.6, 2.7, 4.2, 1.3},
		FPoint{5.7, 3.0, 4.2, 1.2},
		FPoint{5.7, 2.9, 4.2, 1.3},
		FPoint{6.2, 2.9, 4.3, 1.3},
		FPoint{5.1, 2.5, 3.0, 1.1},
		FPoint{5.7, 2.8, 4.1, 1.3},
	},
	"virginica": Points{
		FPoint{6.3, 3.3, 6.0, 2.5},
		FPoint{5.8, 2.7, 5.1, 1.9},
		FPoint{7.1, 3.0, 5.9, 2.1},
		FPoint{6.3, 2.9, 5.6, 1.8},
		FPoint{6.5, 3.0, 5.8, 2.2},
		FPoint{7.6, 3.0, 6.6, 2.1},
		FPoint{4.9, 2.5, 4.5, 1.7},
		FPoint{7.3, 2.9, 6.3, 1.8},
		FPoint{6.7, 2.5, 5.8, 1.8},
		FPoint{7.2, 3.6, 6.1, 2.5},
		FPoint{6.5, 3.2, 5.1, 2.0},
		FPoint{6.4, 2.7, 5.3, 1.9},
		FPoint{6.8, 3.0, 5.5, 2.1},
		FPoint{5.7, 2.5, 5.0, 2.0},
		FPoint{5.8, 2.8, 5.1, 2.4},
		FPoint{6.4, 3.2, 5.3, 2.3},
		FPoint{6.5, 3.0, 5.5, 1.8},
		FPoint{7.7, 3.8, 6.7, 2.2},
		FPoint{7.7, 2.6, 6.9, 2.3},
		FPoint{6.0, 2.2, 5.0, 1.5},
		FPoint{6.9, 3.2, 5.7, 2.3},
		FPoint{5.6, 2.8, 4.9, 2.0},
		FPoint{7.7, 2.8, 6.7, 2.0},
		FPoint{6.3, 2.7, 4.9, 1.8},
		FPoint{6.7, 3.3, 5.7, 2.1},
		FPoint{7.2, 3.2, 6.0, 1.8},
		FPoint{6.2, 2.8, 4.8, 1.8},
		FPoint{6.1, 3.0, 4.9, 1.8},
		FPoint{6.4, 2.8, 5.6, 2.1},
		FPoint{7.2, 3.0, 5.8, 1.6},
		FPoint{7.4, 2.8, 6.1, 1.9},
		FPoint{7.9, 3.8, 6.4, 2.0},
		FPoint{6.4, 2.8, 5.6, 2.2},
		FPoint{6.3, 2.8, 5.1, 1.5},
		FPoint{6.1, 2.6, 5.6, 1.4},
		FPoint{7.7, 3.0, 6.1, 2.3},
		FPoint{6.3, 3.4, 5.6, 2.4},
		FPoint{6.4, 3.1, 5.5, 1.8},
		FPoint{6.0, 3.0, 4.8, 1.8},
		FPoint{6.9, 3.1, 5.4, 2.1},
		FPoint{6.7, 3.1, 5.6, 2.4},
		FPoint{6.9, 3.1, 5.1, 2.3},
		FPoint{5.8, 2.7, 5.1, 1.9},
		FPoint{6.8, 3.2, 5.9, 2.3},
		FPoint{6.7, 3.3, 5.7, 2.5},
		FPoint{6.7, 3.0, 5.2, 2.3},
		FPoint{6.3, 2.5, 5.0, 1.9},
		FPoint{6.5, 3.0, 5.2, 2.0},
		FPoint{6.2, 3.4, 5.4, 2.3},
		FPoint{5.9, 3.0, 5.1, 1.8},
	},
}

// sortMap sorts each point set in the map.
func sortMap(m map[string]Points) {
	for _, p := range m {
		p.Sort()
	}
}

// mapToPoints returns a single sorted set of points.
func mapToPoints(m map[string]Points) Points {
	var n int
	for _, p := range m {
		n += len(p)
	}

	points := make(Points, 0, n)
	for _, p := range m {
		points = append(points, p...)
	}

	points.Sort()
	return points
}

// TestModel tests a small set of points.
func TestModel(t *testing.T) {
	tests := []struct {
		k                int
		pnts             Points
		expectedClusters Clusters
	}{
		{
			k:    2,
			pnts: Points{FPoint{1, 1}, FPoint{2, 1}, FPoint{3, 3}, FPoint{4, 2}, FPoint{4, 3}},
			expectedClusters: Clusters{
				Cluster{FPoint{1, 1}, FPoint{2, 1}},
				Cluster{FPoint{3, 3}, FPoint{4, 2}, FPoint{4, 3}},
			},
		},
	}

	for _, test := range tests {
		mdl := New(test.k, test.pnts)
		mdl.sort()
		mdl.sortAll(LexiSort)
		sort.SliceStable(test.expectedClusters, func(i, j int) bool { return test.expectedClusters[i].CompareTo(test.expectedClusters[j]) < 0 })
		for h := range test.expectedClusters {
			sort.SliceStable(test.expectedClusters[h], func(i, j int) bool { return test.expectedClusters[h][i].CompareTo(test.expectedClusters[h][j]) < 0 })
		}

		for i := range mdl.clusters {
			if test.expectedClusters[i].CompareTo(mdl.clusters[i]) != 0 {
				t.Fatalf("Expected: %0.2f\nReceived: %0.2f\nMeans: %0.2f\n", test.expectedClusters, mdl.clusters, mdl.clusters.Means())
			}
		}
	}
}

// TestSepals tests the classic sepal data set.
func TestSepals(t *testing.T) {
	// Sort so species can be searched.
	sortMap(species)

	// Get the assignment of the species means and verify each flower in each species is assigned to the correct cluster
	var (
		points      = mapToPoints(species)
		model       = New(len(species), points)
		assignments = map[string]int{
			"setosa":     model.Assignment(species["setosa"].ToCluster().Mean()),
			"versicolor": model.Assignment(species["versicolor"].ToCluster().Mean()),
			"virginica":  model.Assignment(species["virginica"].ToCluster().Mean()),
		}
	)

	// Check that assignments are distinct.
	for species0, assignment0 := range assignments {
		for species1, assignment1 := range assignments {
			if species0 != species1 && assignment0 == assignment1 {
				t.Fatal("failed to categorize test species means")
			}
		}
	}

	var (
		numSetosa     = len(species["setosa"])
		numVersicolor = len(species["versicolor"])
		numVirginica  = len(species["virginica"])
		correct       float64
	)

	for _, p := range points {
		switch model.Assignment(p) {
		case assignments["setosa"]:
			if sort.Search(numSetosa, func(j int) bool { return p.CompareTo(species["setosa"][j]) <= 0 }) < numSetosa {
				correct++
			}
		case assignments["versicolor"]:
			if sort.Search(numVersicolor, func(j int) bool { return p.CompareTo(species["versicolor"][j]) <= 0 }) < numVersicolor {
				correct++
			}
		case assignments["virginica"]:
			if sort.Search(numVirginica, func(j int) bool { return p.CompareTo(species["virginica"][j]) <= 0 }) < numVirginica {
				correct++
			}
		}
	}

	if correct < 100 {
		t.Fatalf("k-means model trained and was correct only %0.2f%% of the time", correct)
	}
}

// TestMaxPow tests the maxPow function.
func TestMaxPow(t *testing.T) {
	tests := []struct {
		b        int
		n        int
		expected int
		actual   int
	}{
		{
			b:        10,
			n:        0,
			expected: 0,
			actual:   0,
		},
		{
			b:        10,
			n:        1,
			expected: 0,
			actual:   0,
		},
		{
			b:        10,
			n:        9,
			expected: 0,
			actual:   0,
		},
		{
			b:        10,
			n:        10,
			expected: 1,
			actual:   0,
		},
		{
			b:        10,
			n:        11,
			expected: 1,
			actual:   0,
		},
		{
			b:        2,
			n:        7,
			expected: 2,
			actual:   0,
		},
	}

	for _, test := range tests {
		test.actual = maxPow(test.b, test.n)
		if test.expected != test.actual {
			t.Fatalf("expected: %d\nactual: %d\n", test.expected, test.actual)
		}
	}
}
