package kmeans2

// A Point is a labeled value.
type Point struct {
	label string
	value Interface
}

// NewPoint returns a new point given a label and a value.
func NewPoint(label string, value Interface) Point {
	return Point{label: label, value: value}
}

// Compare two points.
func (p *Point) Compare(point Point) int {
	return p.value.Compare(point.value)
}

// Copy a point.
func (p *Point) Copy() Point {
	return NewPoint(p.label, p.value)
}

// Dist returns the distance between two points.
func (p *Point) Dist(q Point) float64 {
	return p.value.Dist(q.value)
}

// IsZero returns true if a point has no value.
func (p *Point) IsZero() bool {
	return p.value == nil
}

// Label returns the label.
func (p *Point) Label() string {
	return p.label
}

// Len returns the length of a point.
func (p *Point) Len() int {
	return p.value.Len()
}

// Relabel a point.
func (p *Point) Relabel(label string) {
	p.label = label
}
