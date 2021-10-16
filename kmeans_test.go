package kmeans

import (
	"sort"
	"testing"
)

func TestClass(t *testing.T) {
	tests := []struct {
		means []Point
		pnt   Point
		exp   int
	}{
		{
			means: []Point{{2.0}, {5.0}, {8.0}},
			pnt:   Point{1.5},
			exp:   0,
		},
		{
			means: []Point{{2.0}, {5.0}, {8.0}},
			pnt:   Point{2.0},
			exp:   0,
		},
		{
			means: []Point{{2.0}, {5.0}, {8.0}},
			pnt:   Point{2.5},
			exp:   0,
		},
		{
			means: []Point{{2.0}, {5.0}, {8.0}},
			pnt:   Point{4.5},
			exp:   1,
		},
		{
			means: []Point{{2.0}, {5.0}, {8.0}},
			pnt:   Point{5.0},
			exp:   1,
		},
		{
			means: []Point{{2.0}, {5.0}, {8.0}},
			pnt:   Point{5.5},
			exp:   1,
		},
		{
			means: []Point{{2.0}, {5.0}, {8.0}},
			pnt:   Point{7.5},
			exp:   2,
		},
		{
			means: []Point{{2.0}, {5.0}, {8.0}},
			pnt:   Point{8.0},
			exp:   2,
		},
		{
			means: []Point{{2.0}, {5.0}, {8.0}},
			pnt:   Point{8.5},
			exp:   2,
		},
	}

	for _, test := range tests {
		mdl := New(len(test.means), test.means, SetInitMethod(FirstK))
		if rec := mdl.Class(test.pnt); test.exp != rec {
			t.Errorf("\nexpected %d\nreceived %d\n", test.exp, rec)
		}

		meanDists := newTriMatrix(mdl.K())
		meanDists.update(mdl.Means())

		if rec, _ := mdl.classDistMem(test.pnt, meanDists); test.exp != rec {
			t.Errorf("\nexpected %d\nreceived %d\n", test.exp, rec)
		}
	}
}

func TestKMeans(t *testing.T) {
	/*
		   |
		   +   x   x
		   |
		   +   x               x
		   |
		   +               x   x
		   |
		   +       x           x
		   |
		   +   x       x
		   |
		---+---+---+---+---+---+---
		   |
	*/

	const tol = 1e-09
	data := []Point{
		{1.0, 1.0},
		{2.0, 2.0},
		{3.0, 1.0},
		// Exp mean: (2, 1.3...)

		{1.0, 4.0},
		{1.0, 5.0},
		{2.0, 5.0},
		// Exp mean: (1.3..., 4.6...)

		{4.0, 3.0},
		{5.0, 2.0},
		{5.0, 3.0},
		{5.0, 4.0},
		// Exp mean: (4.75, 3.0)
	}

	{
		var (
			expMeans = []Point{
				{4.0 / 3.0, 14.0 / 3.0},
				{2.0, 4.0 / 3.0},
				{19.0 / 4.0, 3.0},
			}
			mdl      = New(3, data)
			recMeans = mdl.Means()
		)

		sort.Slice(recMeans, func(i, j int) bool { return recMeans[i].Compare(recMeans[j]) < 0 })

		if len(expMeans) != len(recMeans) {
			t.Errorf("\nexpected %v\nreceived %v\n", expMeans, recMeans)
		} else {
			for i := 0; i < len(expMeans); i++ {
				if !expMeans[i].Near(recMeans[i], tol) {
					t.Errorf("\nexpected %v\nreceived %v\n", expMeans[i], recMeans[i])
				}
			}
		}

		t.Logf("\nKMeans.Random: %v\n", recMeans)
	}

	{
		var (
			expMeans = []Point{
				{4.0 / 3.0, 14.0 / 3.0},
				{2.0, 4.0 / 3.0},
				{19.0 / 4.0, 3.0},
			}
			mdl      = New(3, data, SetInitMethod(PlusPlus))
			recMeans = mdl.Means()
		)

		sort.Slice(recMeans, func(i, j int) bool { return recMeans[i].Compare(recMeans[j]) < 0 })

		if len(expMeans) != len(recMeans) {
			t.Errorf("\nexpected %v\nreceived %v\n", expMeans, recMeans)
		} else {
			for i := 0; i < len(expMeans); i++ {
				if !expMeans[i].Near(recMeans[i], tol) {
					t.Errorf("\nexpected %v\nreceived %v\n", expMeans[i], recMeans[i])
				}
			}
		}

		t.Logf("\nKMeans.PlusPlus: %v\n", recMeans)
	}
}
