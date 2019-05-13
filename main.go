package main

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
)

// seeded indicates if the random number generator has been seeded.
var seeded bool

func main() {
	test3()
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
	// for k := 1; k <= n; k++ {
	for k := 2; k <= 2; k++ {
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

	// fmt.Println(asciigraph.Plot(variances))
	fmt.Printf("variances: %0.2f\n", variances)
}
