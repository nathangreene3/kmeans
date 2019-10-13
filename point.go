package kmeans

import "fmt"

// Point is an n-dimensional point in n-space.
type Point struct {
	label  interface{} // TODO: Make this field a string
	values Interface
}

// Interface ...
type Interface interface {
	At(i int) float64         // Return the ith dimension
	Compare(p Interface) int  // Return a value from {-1,0,1}
	Copy() Interface          // Return a copy of a point
	Dist(p Interface) float64 // Return the distance between points
	Len() int                 // Return the number of dimensions of a point
}

// NewPoint ...
func NewPoint(label interface{}, x Interface) Point {
	return Point{label: label, values: x.Copy()}
}

// Zero ...
func Zero() Point {
	return Point{}
}

// At returns the ith value.
func (p *Point) At(i int) float64 {
	return p.values.At(i)
}

// Compare two points.
func (p *Point) Compare(q Point) int {
	return p.values.Compare(q.values)
}

// Copy a point.
func (p *Point) Copy() Point {
	return NewPoint(p.label, p.values)
}

// Dist ...
func (p *Point) Dist(q Point) float64 {
	return p.values.Dist(q.values)
}

// Relabel a point.
func (p *Point) Relabel(label interface{}) {
	p.label = label
}

// Len ...
func (p *Point) Len() int {
	return p.values.Len()
}

// String ...
func (p *Point) String() string {
	return fmt.Sprintf("{ %s, %s }", p.label, p.values)
}
