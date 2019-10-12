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

// At ...
func (p *Point) At(i int) float64 {
	return p.values.At(i)
}

// Compare ...
func (p *Point) Compare(q Point) int {
	return p.values.Compare(q.values)
}

// Copy ...
func (p *Point) Copy() Point {
	return NewPoint(p.label, p.values)
}

// Dist ...
func (p *Point) Dist(q Point) float64 {
	return p.values.Dist(q.values)
}

// Len ...
func (p *Point) Len() int {
	return p.values.Len()
}

// String ...
func (p *Point) String() string {
	return fmt.Sprintf("{ %s, %s }", p.label, p.values)
}
