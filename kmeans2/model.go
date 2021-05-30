package kmeans2

import (
	"fmt"

	"github.com/nathangreene3/table"
)

// Model defines the k-means model.
type Model interface {
	Class(Point) int
	Clusters(...Point) [][]Point
	Dist(int, Point) float64
	Err(int, ...Point) (float64, int)
	Errs(...Point) ([]float64, []int)
	K() int
	Mean(int) Point
	Score(...Point) float64
	Train(...Point) Model
}

// Report ...
func Report(mdl Model, data ...Point) *table.Table {
	tbl := table.New(table.NewHeader("", "Point", "Class"))
	for i := 0; i < len(data); i++ {
		tbl.Append(table.NewRow(i, fmt.Sprintf("%v", data[i]), mdl.Class(data[i])))
	}

	return tbl
}
