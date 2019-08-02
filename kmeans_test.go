package main

import (
	"sort"
	"testing"
)

func TestModel(t *testing.T) {
	tests := []struct {
		k                int
		normalize        bool
		pnts             Points
		expectedClusters Clusters
	}{
		{
			k:         2,
			normalize: false,
			pnts:      Points{Point{1, 1}, Point{2, 1}, Point{3, 3}, Point{4, 2}, Point{4, 3}},
			expectedClusters: Clusters{
				Cluster{Point{1, 1}, Point{2, 1}},
				Cluster{Point{3, 3}, Point{4, 2}, Point{4, 3}},
			},
		},
	}

	for _, test := range tests {
		mdl := New(test.k, test.pnts, test.normalize)
		mdl.clusters.Sort()
		mdl.sortAll(LexiSort)

		sort.SliceStable(test.expectedClusters, func(i, j int) bool { return test.expectedClusters[i].CompareTo(test.expectedClusters[j]) < 0 })
		for h := range test.expectedClusters {
			sort.SliceStable(test.expectedClusters[h], func(i, j int) bool { return test.expectedClusters[h][i].CompareTo(test.expectedClusters[h][j]) < 0 })
		}

		for i := range mdl.clusters {
			if test.expectedClusters[i].CompareTo(mdl.clusters[i]) != 0 {
				t.Fatalf("KMeans failed.\nExpected: %0.2f\nReceived: %0.2f\nMeans: %0.2f\n", test.expectedClusters, mdl.clusters, mdl.clusters.Means())
			}
		}
	}
}

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
