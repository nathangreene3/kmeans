package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	test1()
}

func test0() {
	species := map[string]Points{
		"setosa": Points{
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
		},
		"versicolor": Points{
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
		},
		"virginica": Points{
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
		},
	}

	// Sort so species can be searched.
	for s := range species {
		species[s].Sort()
	}

	var (
		numSetosa     = len(species["setosa"])
		numVersicolor = len(species["versicolor"])
		numVirginica  = len(species["virginica"])
		numPnts       = numSetosa + numVersicolor + numVirginica
		pnts          = make(Points, 0, numPnts)
	)

	for _, s := range species {
		for _, p := range s {
			pnts = append(pnts, p)
		}
	}

	// Get the assignment of the species means and verify each flower in each species is assigned to the correct cluster
	var (
		mdl         = New(len(species), pnts, false)
		assignments = map[string]int{
			"setosa":     mdl.Assignment(species["setosa"].ToCluster().Mean()),
			"versicolor": mdl.Assignment(species["versicolor"].ToCluster().Mean()),
			"virginica":  mdl.Assignment(species["virginica"].ToCluster().Mean()),
		}
	)

	// Check that assignments are distinct.
	for species0, assignment0 := range assignments {
		for species1, assignment1 := range assignments {
			if species0 != species1 && assignment0 == assignment1 {
				log.Fatal("failed to categorize test species means")
			}
		}
	}

	var correct float64
	for _, p := range pnts {
		switch mdl.Assignment(p) {
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

	fmt.Printf(
		"Training resulted in %0.2f%% correct\n\n%s\n",
		100*correct/float64(numPnts),
		PlotMeanWeightedVars(1, 5, pnts, false),
	)
}

func test1() {
	species := map[string]Points{
		"setosa": Points{
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
		},
		"versicolor": Points{
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
		},
		"virginica": Points{
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
		},
	}

	// Sort so species can be searched.
	for s := range species {
		species[s].Sort()
	}

	var (
		numSetosa     = len(species["setosa"])
		numVersicolor = len(species["versicolor"])
		numVirginica  = len(species["virginica"])
		numPnts       = numSetosa + numVersicolor + numVirginica
		pnts          = make(Points, 0, numPnts)
	)

	for _, s := range species {
		for _, p := range s {
			pnts = append(pnts, p)
		}
	}

	fmt.Printf("Optimal K: %d\n", optimalK(pnts, false, 1, 10))
}
