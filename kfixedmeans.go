package main

import "math"

func distRatio(pnt Point, mns Points) float64 {
	return SquaredDistance(pnt, mns[minIndexOf(pnt, mns)]) / SquaredDistance(pnt, mns[minIndexOf(pnt, mns)])
}

func minIndexOf(pnt Point, mns Points) int {
	var (
		index int
		sd    float64
		minSD = math.MaxFloat64
	)
	for i := range mns {
		sd = SquaredDistance(pnt, mns[i])
		if sd < minSD {
			index = i
			minSD = sd
		}
	}

	return index
}
