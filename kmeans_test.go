package main

import (
	"sort"
	"testing"
)

func TestKMeans(t *testing.T) {
	tests := []struct {
		k         int
		pnts      []Point
		expClstrs []Cluster
		actClstrs []Cluster
	}{
		{
			k:    2,
			pnts: []Point{Point{1, 1}, Point{2, 1}, Point{3, 3}, Point{4, 2}, Point{4, 3}},
			expClstrs: []Cluster{
				Cluster{Point{1, 1}, Point{2, 1}},
				Cluster{Point{3, 3}, Point{4, 2}, Point{4, 3}},
			},
		},
	}

	for _, test := range tests {
		test.actClstrs = KMeans(test.k, test.pnts)
		for h := range test.expClstrs {
			sort.SliceStable(test.expClstrs[h], func(i, j int) bool { return ComparePoints(test.expClstrs[h][i], test.expClstrs[h][j]) < 0 })
		}

		for h := range test.actClstrs {
			sort.SliceStable(test.actClstrs[h], func(i, j int) bool { return ComparePoints(test.actClstrs[h][i], test.actClstrs[h][j]) < 0 })
		}

		for i := range test.actClstrs {
			if CompareClusters(test.expClstrs[i], test.actClstrs[i]) != 0 {
				t.Fatalf("KMeans failed.\nExpected: %0.2f\nReceived: %0.2f\n", test.expClstrs, test.actClstrs)
			}
		}
	}
}

// func TestNormalizeKMeans(t *testing.T){
// 	tests := []struct {
// 		k         int
// 		pnts      []Point
// 		expClstrs []Cluster
// 		actClstrs []Cluster
// 	}{
// 		{
// 			k:    2,
// 			pnts: []Point{Point{1, 1}, Point{2, 1}, Point{3, 3}, Point{4, 2}, Point{4, 3}},
// 			expClstrs: []Cluster{
// 				Cluster{Point{1, 1}, Point{2, 1}},
// 				Cluster{Point{3, 3}, Point{4, 2}, Point{4, 3}},
// 			},
// 		},
// 	}
// }
