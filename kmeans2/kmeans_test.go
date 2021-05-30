package kmeans2

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/guptarohit/asciigraph"
	"github.com/nathangreene3/table"
)

var species = map[string][]Point{
	"setosa": {
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
	"versicolor": {
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
	"virginica": {
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

func TestAssignment(t *testing.T) {
	tests := []struct {
		kmns Model
		pnt  Point
		exp  int
	}{
		{
			kmns: New(FPoint{2.0}, FPoint{5.0}, FPoint{8.0}),
			pnt:  FPoint{1.5},
			exp:  0,
		},
		{
			kmns: New(FPoint{2.0}, FPoint{5.0}, FPoint{8.0}),
			pnt:  FPoint{2.0},
			exp:  0,
		},
		{
			kmns: New(FPoint{2.0}, FPoint{5.0}, FPoint{8.0}),
			pnt:  FPoint{2.5},
			exp:  0,
		},
		{
			kmns: New(FPoint{2.0}, FPoint{5.0}, FPoint{8.0}),
			pnt:  FPoint{4.5},
			exp:  1,
		},
		{
			kmns: New(FPoint{2.0}, FPoint{5.0}, FPoint{8.0}),
			pnt:  FPoint{5.0},
			exp:  1,
		},
		{
			kmns: New(FPoint{2.0}, FPoint{5.0}, FPoint{8.0}),
			pnt:  FPoint{5.5},
			exp:  1,
		},
		{
			kmns: New(FPoint{2.0}, FPoint{5.0}, FPoint{8.0}),
			pnt:  FPoint{7.5},
			exp:  2,
		},
		{
			kmns: New(FPoint{2.0}, FPoint{5.0}, FPoint{8.0}),
			pnt:  FPoint{8.0},
			exp:  2,
		},
		{
			kmns: New(FPoint{2.0}, FPoint{5.0}, FPoint{8.0}),
			pnt:  FPoint{8.5},
			exp:  2,
		},
	}

	for _, test := range tests {
		if rec := test.kmns.Class(test.pnt); test.exp != rec {
			t.Errorf("\nexpected %d\nreceived %d\n", test.exp, rec)
		}

		if rec, _ := test.kmns.(kMeans).classDist(test.pnt); test.exp != rec {
			t.Errorf("\nexpected %d\nreceived %d\n", test.exp, rec)
		}

		meanDists := newDistMtx(test.kmns.(kMeans)...)
		if rec, _ := test.kmns.(kMeans).classDistMem(test.pnt, meanDists); test.exp != rec {
			t.Errorf("\nexpected %d\nreceived %d\n", test.exp, rec)
		}
	}
}

func TestOneDimensionalPoints(t *testing.T) {
	seed := int64(time.Now().Nanosecond())
	rand.Seed(seed)
	fmt.Printf("seed: %d\n", seed)

	tests := []struct {
		k    int
		data []Point
	}{
		{
			k: 3,
			data: []Point{
				FPoint{1.9},
				FPoint{2.0},
				FPoint{2.1},
				FPoint{4.9},
				FPoint{5.0},
				FPoint{5.1},
				FPoint{7.9},
				FPoint{8.0},
				FPoint{8.1},
			},
		},
	}

	for _, test := range tests {
		kmns := KMeans(test.k, 3, test.data...)
		fmt.Println(kmns)
	}

	t.Fail()
}

func TestSepals(t *testing.T) {
	seed := int64(time.Now().Nanosecond())
	rand.Seed(seed)

	var n int
	for _, v := range species {
		n += len(v)
	}

	data := make([]Point, 0, n)
	for _, v := range species {
		data = append(data, v...)
	}

	rand.Shuffle(n, func(i, j int) { data[i], data[j] = data[j], data[i] })

	var (
		kmns    = KMeans(len(species), 3, data...)
		expKMns = kMeans{
			KMeans(1, 3, species["setosa"]...).(kMeans)[0],
			KMeans(1, 3, species["versicolor"]...).(kMeans)[0],
			KMeans(1, 3, species["virginica"]...).(kMeans)[0],
		}
		classes         = make([]int, len(data))
		expClasses      = make([]int, len(data))
		clusterFreqs    = make([]int, len(species))
		expClusterFreqs = make([]int, len(species))
	)

	fmt.Printf("\n"+
		"      seed: %d\n"+
		"   k-means: %v\n"+
		"    setosa: %v\n"+
		"versicolor: %v\n"+
		" virginica: %v\n",
		seed,
		kmns,
		expKMns[0],
		expKMns[1],
		expKMns[2],
	)

	for i := 0; i < len(data); i++ {
		classes[i] = expKMns.Class(data[i])
		expClasses[i] = expKMns.Class(data[i])
		clusterFreqs[classes[i]]++
		expClusterFreqs[expClasses[i]]++
	}

	fmt.Printf("\n"+
		"         cluster sizes: %v\n"+
		"expected cluster sizes: %v\n",
		clusterFreqs,
		expClusterFreqs,
	)

	r := Report(kmns, data...)
	if err := r.ToCSV("iris_clusters_" + time.Now().Format(time.RFC3339) + ".csv"); err != nil {
		t.Fatal(err)
	}

	fmt.Println(r.Format(table.Fmt4))
	t.Fail()
}

func TestSepalsFindK(t *testing.T) {
	seed := int64(time.Now().Nanosecond())
	rand.Seed(seed)

	var n int
	for _, v := range species {
		n += len(v)
	}

	data := make([]Point, 0, n)
	for _, v := range species {
		data = append(data, v...)
	}

	tbl := table.New(table.NewHeader("K", "Score"))
	scores := make([]float64, 0, 10)
	for k := 1; k <= 10; k++ {
		tbl.Append(table.NewRow(k, KMeans(k, 3, data...).Score(data...)))
		scores = append(scores, tbl.GetFlt(k-1, 1))
	}

	fmt.Printf(
		"%s\n%s\n",
		tbl.Format(table.Fmt4),
		asciigraph.Plot(scores, asciigraph.Caption("Scores")),
	)

	t.Fail()
}

func BenchmarkSepals(b *testing.B) {
	var n int
	for _, v := range species {
		n += len(v)
	}

	data := make([]Point, 0, n)
	for _, v := range species {
		data = append(data, v...)
	}

	var kmns Model
	for i := 0; i < b.N; i++ {
		// rand.Seed(1) // This probably ruins the benchmark
		kmns = KMeans(len(species), 1, data...)
	}

	_ = kmns
}
