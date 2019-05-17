package main

import (
	"fmt"
	"sort"

	"github.com/guptarohit/asciigraph"
)

// seeded indicates if the random number generator has been seeded.
var seeded bool

func main() {
	test5()
}

func test0() {
	pnts := []Point{
		Point{4, 3},
		Point{1, 1},
		Point{3, 3},
		Point{2, 1},
		Point{4, 2},
	}

	n := len(pnts)
	variances := make([]float64, 0, n)
	var v float64
	var clstrs []Cluster
	for k := 1; k <= n; k++ {
		fmt.Printf("k = %d\n", k)
		clstrs = KMeans(k, pnts, false)
		v = 0
		for i := range clstrs {
			fmt.Printf("cluster %d\n", i)
			for j := range clstrs[i] {
				fmt.Printf("point %d: %0.2f\n", j, clstrs[i][j])
			}

			fmt.Printf("mean: %0.2f\n\n", Mean(clstrs[i]))
			v += Variance(clstrs[i])
		}

		variances = append(variances, v/float64(k))
	}

	fmt.Println(asciigraph.Plot(variances))
	fmt.Printf("variances: %0.2f\n", variances)
}

func test1() {
	pnts := []Point{
		Point{4, 3},
		Point{1, 1},
		Point{3, 3},
		Point{2, 1},
		Point{4, 2},
	}

	k, clstrs := OptimalKMeans(pnts, false)
	fmt.Printf("k = %d\n", k)
	for i := range clstrs {
		fmt.Printf("cluster %d: %0.2f\n", i, clstrs[i])
	}
}

func test2() {
	pnts := []Point{
		Point{5},
		Point{5},
	}

	k, _ := OptimalKMeans(pnts, false)
	fmt.Printf("k = %d\n", k)
}

func test3() {
	pnts := []Point{
		Point{2, 1}, // (0.50, 0.33)
		Point{4, 3}, // (1.00, 1.00)
		Point{4, 2}, // (1.00, 0.67)
		Point{3, 3}, // (0.75, 1.00)
		Point{1, 1}, // (0.25, 0.33)
	}

	n := len(pnts)
	variances := make([]float64, 0, n)
	var v float64
	var clstrs []Cluster
	for k := 1; k <= n; k++ {
		fmt.Printf("k = %d\n", k)
		clstrs = KMeans(k, pnts, true)
		v = 0
		for i := range clstrs {
			fmt.Printf("cluster %d\n", i)
			for j := range clstrs[i] {
				fmt.Printf("point %d: %0.2f\n", j, clstrs[i][j])
			}

			fmt.Printf("mean: %0.2f\n\n", Mean(clstrs[i]))
			v += Variance(clstrs[i])
		}

		variances = append(variances, v/float64(k))
	}

	fmt.Println(asciigraph.Plot(variances))
	fmt.Printf("variances: %0.2f\n", variances)
}

func test4() {
	pnts := []Point{
		Point{5, 10},
		Point{5, 10},
		Point{5, 10},
		Point{5, 10},
		Point{5, 10},
	}

	n := len(pnts)
	variances := make([]float64, 0, n)
	var v float64
	var clstrs []Cluster
	for k := 1; k <= n; k++ {
		fmt.Printf("k = %d\n", k)
		clstrs = KMeans(k, pnts, false)
		v = 0
		for i := range clstrs {
			fmt.Printf("cluster %d\n", i)
			for j := range clstrs[i] {
				fmt.Printf("point %d: %0.2f\n", j, clstrs[i][j])
			}

			fmt.Printf("mean: %0.2f\n\n", Mean(clstrs[i]))
			v += Variance(clstrs[i])
		}

		variances = append(variances, v/float64(k))
	}

	fmt.Println(asciigraph.Plot(variances))
	fmt.Printf("variances: %0.2f\n", variances)
}

func test5() {
	// species := []string{"setosa", "versicolor", "virginica"}
	species := map[string][]Point{
		"setosa": []Point(SortCluster(Cluster([]Point{
			Point{5.1, 3.5, 1.4, 0.2},
			Point{4.9, 3.0, 1.4, 0.2},
			Point{4.7, 3.2, 1.3, 0.2},
			Point{4.6, 3.1, 1.5, 0.2},
			Point{5.0, 3.6, 1.4, 0.2},
			Point{5.4, 3.9, 1.7, 0.4},
			Point{4.6, 3.4, 1.4, 0.3},
			Point{5.0, 3.4, 1.5, 0.2},
			Point{4.4, 2.9, 1.4, 0.2},
			Point{4.9, 3.1, 1.5, 0.1},
			Point{5.4, 3.7, 1.5, 0.2},
			Point{4.8, 3.4, 1.6, 0.2},
			Point{4.8, 3.0, 1.4, 0.1},
			Point{4.3, 3.0, 1.1, 0.1},
			Point{5.8, 4.0, 1.2, 0.2},
			Point{5.7, 4.4, 1.5, 0.4},
			Point{5.4, 3.9, 1.3, 0.4},
			Point{5.1, 3.5, 1.4, 0.3},
			Point{5.7, 3.8, 1.7, 0.3},
			Point{5.1, 3.8, 1.5, 0.3},
			Point{5.4, 3.4, 1.7, 0.2},
			Point{5.1, 3.7, 1.5, 0.4},
			Point{4.6, 3.6, 1.0, 0.2},
			Point{5.1, 3.3, 1.7, 0.5},
			Point{4.8, 3.4, 1.9, 0.2},
			Point{5.0, 3.0, 1.6, 0.2},
			Point{5.0, 3.4, 1.6, 0.4},
			Point{5.2, 3.5, 1.5, 0.2},
			Point{5.2, 3.4, 1.4, 0.2},
			Point{4.7, 3.2, 1.6, 0.2},
			Point{4.8, 3.1, 1.6, 0.2},
			Point{5.4, 3.4, 1.5, 0.4},
			Point{5.2, 4.1, 1.5, 0.1},
			Point{5.5, 4.2, 1.4, 0.2},
			Point{4.9, 3.1, 1.5, 0.1},
			Point{5.0, 3.2, 1.2, 0.2},
			Point{5.5, 3.5, 1.3, 0.2},
			Point{4.9, 3.1, 1.5, 0.1},
			Point{4.4, 3.0, 1.3, 0.2},
			Point{5.1, 3.4, 1.5, 0.2},
			Point{5.0, 3.5, 1.3, 0.3},
			Point{4.5, 2.3, 1.3, 0.3},
			Point{4.4, 3.2, 1.3, 0.2},
			Point{5.0, 3.5, 1.6, 0.6},
			Point{5.1, 3.8, 1.9, 0.4},
			Point{4.8, 3.0, 1.4, 0.3},
			Point{5.1, 3.8, 1.6, 0.2},
			Point{4.6, 3.2, 1.4, 0.2},
			Point{5.3, 3.7, 1.5, 0.2},
			Point{5.0, 3.3, 1.4, 0.2},
		}), LexiSort)),
		"versicolor": []Point(SortCluster(Cluster([]Point{
			Point{7.0, 3.2, 4.7, 1.4},
			Point{6.4, 3.2, 4.5, 1.5},
			Point{6.9, 3.1, 4.9, 1.5},
			Point{5.5, 2.3, 4.0, 1.3},
			Point{6.5, 2.8, 4.6, 1.5},
			Point{5.7, 2.8, 4.5, 1.3},
			Point{6.3, 3.3, 4.7, 1.6},
			Point{4.9, 2.4, 3.3, 1.0},
			Point{6.6, 2.9, 4.6, 1.3},
			Point{5.2, 2.7, 3.9, 1.4},
			Point{5.0, 2.0, 3.5, 1.0},
			Point{5.9, 3.0, 4.2, 1.5},
			Point{6.0, 2.2, 4.0, 1.0},
			Point{6.1, 2.9, 4.7, 1.4},
			Point{5.6, 2.9, 3.6, 1.3},
			Point{6.7, 3.1, 4.4, 1.4},
			Point{5.6, 3.0, 4.5, 1.5},
			Point{5.8, 2.7, 4.1, 1.0},
			Point{6.2, 2.2, 4.5, 1.5},
			Point{5.6, 2.5, 3.9, 1.1},
			Point{5.9, 3.2, 4.8, 1.8},
			Point{6.1, 2.8, 4.0, 1.3},
			Point{6.3, 2.5, 4.9, 1.5},
			Point{6.1, 2.8, 4.7, 1.2},
			Point{6.4, 2.9, 4.3, 1.3},
			Point{6.6, 3.0, 4.4, 1.4},
			Point{6.8, 2.8, 4.8, 1.4},
			Point{6.7, 3.0, 5.0, 1.7},
			Point{6.0, 2.9, 4.5, 1.5},
			Point{5.7, 2.6, 3.5, 1.0},
			Point{5.5, 2.4, 3.8, 1.1},
			Point{5.5, 2.4, 3.7, 1.0},
			Point{5.8, 2.7, 3.9, 1.2},
			Point{6.0, 2.7, 5.1, 1.6},
			Point{5.4, 3.0, 4.5, 1.5},
			Point{6.0, 3.4, 4.5, 1.6},
			Point{6.7, 3.1, 4.7, 1.5},
			Point{6.3, 2.3, 4.4, 1.3},
			Point{5.6, 3.0, 4.1, 1.3},
			Point{5.5, 2.5, 4.0, 1.3},
			Point{5.5, 2.6, 4.4, 1.2},
			Point{6.1, 3.0, 4.6, 1.4},
			Point{5.8, 2.6, 4.0, 1.2},
			Point{5.0, 2.3, 3.3, 1.0},
			Point{5.6, 2.7, 4.2, 1.3},
			Point{5.7, 3.0, 4.2, 1.2},
			Point{5.7, 2.9, 4.2, 1.3},
			Point{6.2, 2.9, 4.3, 1.3},
			Point{5.1, 2.5, 3.0, 1.1},
			Point{5.7, 2.8, 4.1, 1.3},
		}), LexiSort)),
		"virginica": []Point(SortCluster(Cluster([]Point{
			Point{6.3, 3.3, 6.0, 2.5},
			Point{5.8, 2.7, 5.1, 1.9},
			Point{7.1, 3.0, 5.9, 2.1},
			Point{6.3, 2.9, 5.6, 1.8},
			Point{6.5, 3.0, 5.8, 2.2},
			Point{7.6, 3.0, 6.6, 2.1},
			Point{4.9, 2.5, 4.5, 1.7},
			Point{7.3, 2.9, 6.3, 1.8},
			Point{6.7, 2.5, 5.8, 1.8},
			Point{7.2, 3.6, 6.1, 2.5},
			Point{6.5, 3.2, 5.1, 2.0},
			Point{6.4, 2.7, 5.3, 1.9},
			Point{6.8, 3.0, 5.5, 2.1},
			Point{5.7, 2.5, 5.0, 2.0},
			Point{5.8, 2.8, 5.1, 2.4},
			Point{6.4, 3.2, 5.3, 2.3},
			Point{6.5, 3.0, 5.5, 1.8},
			Point{7.7, 3.8, 6.7, 2.2},
			Point{7.7, 2.6, 6.9, 2.3},
			Point{6.0, 2.2, 5.0, 1.5},
			Point{6.9, 3.2, 5.7, 2.3},
			Point{5.6, 2.8, 4.9, 2.0},
			Point{7.7, 2.8, 6.7, 2.0},
			Point{6.3, 2.7, 4.9, 1.8},
			Point{6.7, 3.3, 5.7, 2.1},
			Point{7.2, 3.2, 6.0, 1.8},
			Point{6.2, 2.8, 4.8, 1.8},
			Point{6.1, 3.0, 4.9, 1.8},
			Point{6.4, 2.8, 5.6, 2.1},
			Point{7.2, 3.0, 5.8, 1.6},
			Point{7.4, 2.8, 6.1, 1.9},
			Point{7.9, 3.8, 6.4, 2.0},
			Point{6.4, 2.8, 5.6, 2.2},
			Point{6.3, 2.8, 5.1, 1.5},
			Point{6.1, 2.6, 5.6, 1.4},
			Point{7.7, 3.0, 6.1, 2.3},
			Point{6.3, 3.4, 5.6, 2.4},
			Point{6.4, 3.1, 5.5, 1.8},
			Point{6.0, 3.0, 4.8, 1.8},
			Point{6.9, 3.1, 5.4, 2.1},
			Point{6.7, 3.1, 5.6, 2.4},
			Point{6.9, 3.1, 5.1, 2.3},
			Point{5.8, 2.7, 5.1, 1.9},
			Point{6.8, 3.2, 5.9, 2.3},
			Point{6.7, 3.3, 5.7, 2.5},
			Point{6.7, 3.0, 5.2, 2.3},
			Point{6.3, 2.5, 5.0, 1.9},
			Point{6.5, 3.0, 5.2, 2.0},
			Point{6.2, 3.4, 5.4, 2.3},
			Point{5.9, 3.0, 5.1, 1.8},
		}), LexiSort)),
	}

	numSetosa := len(species["setosa"])
	numVersicolor := len(species["versicolor"])
	numVirginica := len(species["virginica"])
	n := numSetosa + numVersicolor + numVirginica
	pnts := make([]Point, 0, n)
	for i := range species["setosa"] {
		pnts = append(pnts, species["setosa"][i])
	}

	for i := range species["versicolor"] {
		pnts = append(pnts, species["versicolor"][i])
	}

	for i := range species["virginica"] {
		pnts = append(pnts, species["virginica"][i])
	}

	clstrs := SortAllClusters(KMeans(len(species), pnts, false), LexiSort)
	mns := Means(clstrs)
	var correct float64
	for i := range pnts {
		switch assignPoint(pnts[i], mns) {
		case 0:
			if sort.Search(numSetosa, func(j int) bool { return ComparePoints(pnts[i], species["setosa"][j]) <= 0 }) < numSetosa {
				correct++
			}
		case 1:
			if sort.Search(numVersicolor, func(j int) bool { return ComparePoints(pnts[i], species["versicolor"][j]) <= 0 }) < numVersicolor {
				correct++
			}
		case 2:
			if sort.Search(numVirginica, func(j int) bool { return ComparePoints(pnts[i], species["virginica"][j]) <= 0 }) < numVirginica {
				correct++
			}
		}
	}

	fmt.Printf("%0.2f%%\n", 100.0*correct/float64(n))
}
